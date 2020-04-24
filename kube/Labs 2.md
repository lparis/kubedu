# Docker and Kubernetes Lab

## Lab 01: Containerize Applications

### Part I: Setup

#### Step 1: Download the gowebapp code under your home directory and untar the file

wget https://s3.eu-central-1.amazonaws.com/heptio-edu-static/how/gowebapp.tar.gz

tar -zxvf gowebapp.tar.gz

#### Step 2: Create Dockerfile for your frontend application

cd $HOME/gowebapp/gowebapp

Create a file named Dockerfile in this directory for the frontend Go application. Use vi or any preferred text editor. Populate it with the following code snippet.

```
LABEL maintainer "education@heptio.com"
LABEL version "1.0"

COPY ./code /opt/gowebapp
COPY ./config /opt/gowebapp/config

EXPOSE 80

WORKDIR /opt/gowebapp/
ENTRYPOINT ["/opt/gowebapp/gowebapp"]
```

#### Step 3: Build gowebapp Docker image locally

`cd $HOME/gowebapp/gowebapp`

Build the gowebapp image locally. Make sure to include “.“ at the end. Make sure the build runs to completion without errors. You should get a success message.

`docker build -t gowebapp:v1 .`


### Part II: Build Docker image for your backend application

#### Step 1: Create Dockerfile for your backend application

`cd $HOME/gowebapp/gowebapp-mysql`

Create a file named Dockerfile in this directory for the backend MySQL database. Use vi or any preferred text editor. Populate it with the following code snippet.

```
FROM mysql:5.6

LABEL maintainer "education@heptio.com"
LABEL version="1.0"

COPY gowebapp.sql /docker-entrypoint-initdb.d/
```

#### Step 2: Build gowebapp-mysql Docker image locally

`cd $HOME/gowebapp/gowebapp-mysql`

Build the gowebapp-mysql image locally. Make sure to include “.“ at the end. Make sure the build runs to completion without errors. You should get a success message.

`docker build -t gowebapp-mysql:v1 .`

### Part III. Run and test Docker images locally

Before deploying to Kubernetes, let’s run and test the Docker images locally, to ensure that the frontend and backend containers run and integrate properly.

#### Step 1: Create Docker user-defined network

To facilitate cross-container communication, let’s first define a user-defined network in which to run the frontend and backend containers:

`docker network create gowebapp`

#### Step 2: Launch frontend and backend containers

Next, let’s launch a frontend and backend container using the Docker CLI.

```
docker run --net gowebapp --name gowebapp-mysql --hostname gowebapp-mysql -d -e MYSQL_ROOT_PASSWORD=heptio gowebapp-mysql:v1

sleep 20

docker run -p 9000:80 --net gowebapp -d --name gowebapp --hostname gowebapp gowebapp:v1
```

First, we launched the database container, as it will take a bit longer to startup, and the frontend container depends on it. Notice how we are injecting the database password into the MySQL configuration as an environment variable:

Next we paused for 20 seconds to give the database a chance to start and initialize.

Finally we launched a frontend container, mapping the container port 80 - where the web application is exposed - to port 9000 on the host machine:

#### Step 3: Test the application locally

Now that we’ve launched the application containers, let’s try to test the web application locally. You should be able to access the application at http://<EXTERNAL-IP>:9000.

To get the EXTERNAL-IP of your host, run the command lab-info in your lab terminal

Create an account and login. Write something on your Notepad and save it. This will verify that the application is working and properly integrates with the backend database container.

#### Step 4: Inspect the MySQL database

Let’s connect to the backend MySQL database container and run some queries to ensure that application persistence is working properly:

docker exec -it gowebapp-mysql mysql -u root -pheptio gowebapp
Once connected, run some simple SQL commands to inspect the database tables and persistence:

```
#Simple SQL to navigate
SHOW DATABASES;
USE gowebapp;
SHOW TABLES;
SELECT * FROM <table_name>;
exit;
```

#### Step 5: Cleanup application containers

When we’re finished testing, we can terminate and remove the currently running frontend and backend containers from our local machine:

`docker rm -f gowebapp gowebapp-mysql`

### Part IV. Create and push Docker images to Docker registry

#### Step 1: Tag images to target another registry

We are finished testing our images. We now need to push our images to an image registry so our Kubernetes cluster can access them. First, we need to tag our Docker images to use the registry in your lab environment:

`docker tag gowebapp:v1 localhost:5000/gowebapp:v1`

`docker tag gowebapp-mysql:v1 localhost:5000/gowebapp-mysql:v1`

#### Step 2: publish images to the registry

```
docker push localhost:5000/gowebapp:v1
docker push localhost:5000/gowebapp-mysql:v1
```

#### Step 3: Verify repository images

To verify push of images to the registry:

```
curl localhost:5000/v2/_catalog
```

Lab 01 Conclusion

Congratulations! By containerizing your application components, you have taken the first important step toward deploying your application to Kubernetes.


-----

## Lab 02: Deploy Applications Using Kubernetes

### Getting Started with kubectl

kubectl is the command line interface for interacting with a Kubernetes cluster. Let’s explore some of it’s features.

#### Step 1: introduction to kubectl

By executing kubectl you will get a list of options you can utilize kubectl for. kubectl allows you to control Kubernetes cluster manager.

kubectl

#### Step 2: use kubectl to understand an object

Use explain to get documentation of various resources. For instance pods, nodes, services, etc.

kubectl explain pods

#### Step 3: get more information on an object

Shortcut to object names:

kubectl describe

#### Step 4: autocomplete

Use TAB to autocomplete:

kubectl describe <TAB> <TAB>

### Create Service Object for MySQL

#### Step 1: define a Kubernetes Service object for the backend MySQL database

cd $HOME/gowebapp/gowebapp-mysql

Create a file named gowebapp-mysql-service.yaml in this directory. Use vi or any preferred text editor. Populate it with the following code snippet.

gowebapp-mysql-service.yaml

```
apiVersion: v1
kind: Service
metadata: 
  name: gowebapp-mysql
  labels:
    run: gowebapp-mysql
    tier: backend
spec:
  type: ClusterIP
  ports:
  - port: 3306
    targetPort: 3306
  selector:
    run: gowebapp-mysql
    tier: backend
```

#### Step 2: create a Service defined above

Use kubectl to create the service defined above

kubectl apply -f gowebapp-mysql-service.yaml

#### Step 3: test to make sure the Service was created

kubectl get service -l "run=gowebapp-mysql"

### Create Deployment Object for MySQL

#### Step 1: define a Kubernetes Deployment object for the backend MySQL database

`cd $HOME/gowebapp/gowebapp-mysql`

Create a file named gowebapp-mysql-deployment.yaml in this directory. Use vi or any preferred text editor. Populate it with the following code snippet.

gowebapp-mysql-deployment.yaml

```
apiVersion: apps/v1
kind: Deployment
metadata: 
  name: gowebapp-mysql
  labels:
    run: gowebapp-mysql
    tier: backend
spec:
  replicas: 1
  strategy: 
    type: Recreate
  selector:
    matchLabels:
      run: gowebapp-mysql
      tier: backend
  template:
    metadata:
      labels:
        run: gowebapp-mysql
        tier: backend
    spec:
      containers:
      - name: gowebapp-mysql
        env:
        - name: MYSQL_ROOT_PASSWORD
          value: heptio
        image: localhost:5000/gowebapp-mysql:v1
        ports:
        - containerPort: 3306
```

#### Step 2: create the Deployment defined above

Use kubectl to create the service defined above

kubectl apply -f gowebapp-mysql-deployment.yaml

#### Step 3: test to make sure the Deployment was created

kubectl get deployment -l "run=gowebapp-mysql"

### Create Service Object for frontend application: gowebapp

#### Step 1: define a Kubernetes Service object for the frontend gowebapp

`cd $HOME/gowebapp/gowebapp`

Create a file named gowebapp-service.yaml in this directory. Use vi or any preferred text editor. Populate it with the following code snippet.

gowebapp-service.yaml

```
apiVersion: v1
kind: Service
metadata: 
  name: gowebapp
  labels:
    run: gowebapp
    tier: frontend
spec:
  type: NodePort
  ports:
  - port: 80
  selector:
    run: gowebapp
    tier: frontend
```

#### Step 2: create a Service defined above

Use kubectl to create the service defined above

`kubectl apply -f gowebapp-service.yaml`

#### Step 3: test to make sure the Service was created

`kubectl get service -l "run=gowebapp"`

### Create Deployment Object for gowebapp

#### Step 1: define a Kubernetes Deployment object for the frontend gowebapp

`cd $HOME/gowebapp/gowebapp`

Create a file named gowebapp-deployment.yaml in this directory. Use vi or any preferred text editor. Populate it with the following code snippet.

gowebapp-deployment.yaml

```
apiVersion: apps/v1
kind: Deployment
metadata: 
  name: gowebapp
  labels:
    run: gowebapp
    tier: frontend
spec:
  replicas: 2
  selector:
    matchLabels:
      run: gowebapp
      tier: frontend
  template:
    metadata:
      labels:
        run: gowebapp
        tier: frontend
    spec:
      containers:
      - name: gowebapp
        env:
        - name: MYSQL_ROOT_PASSWORD
          value: heptio
        image: localhost:5000/gowebapp:v1       
        ports:
        - containerPort: 80
```

#### Step 2: create the Deployment defined above

Use kubectl to create the service defined above

`kubectl apply -f gowebapp-deployment.yaml`

#### Step 3: test to make sure the Deployment was created

`kubectl get deployment -l "run=gowebapp"`


### Test Your Application

#### Step 1: Access gowebapp through the NodePort service

Access your application at the NodePort Service endpoint: http://<EXTERNAL-IP>:<NodePort>.

Note

To get the EXTERNAL-IP of your host, run the command lab-info in your lab terminal
Note

To get the NodePort for your service, run the command kubectl get svc gowebapp in your lab terminal

*** EXAMPLE OUTPUT ****

$ kubectl get svc gowebapp

NAME       TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
gowebapp   NodePort   10.107.169.105   <none>        80:30500/TCP   2m

In this example, the NodePort is 30500

You can now sign up for an account, login and use the Notepad.

Note: Your browser may cache an old instance of the gowebapp application from previous labs. When the webpage loads, look at the top right. If you see ‘Logout’, click it. You can then proceed with creating a new test account.
Lab 02 Conclusion

Congratulations! You have successfully deployed your applications using Kubernetes!

