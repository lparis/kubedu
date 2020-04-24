

// Load library
FROM ubuntu:16.04 

// Run at runtime
RUN apt-get update && apt-get install -y apache2 
RUN apt-get clean && rm -rf /var/lib/apt/lists/* 

// Run at build time
ENV APACHE_RUN_USER www 
ENV APACHE_RUN_GROUP www 
ENV APACHE_LOG_DIR /var/log/apache2 

// Default port when container starts
EXPOSE 80 

// How app starts
CMD ["/usr/sbin/apache2", "-D", "FOREGROUND"]


----

curl http://localhost:8001/api/v1/pods{   "kind": "PodList",   "apiVersion": "v1",   "metadata": {     "selfLink": "/api/v1/pods",     "resourceVersion": "3606680"   },   "items": [



----

apiVersion: v1   
kind: Pod   
metadata:     
name: mypod1     
labels:       
app: mynginx       
tier: web   
spec:     containers:     - name: nginx       image: nginx:1.13.1     - name: www-syncer       image: syncer:1.2



---
Services are Layer 4 load balancers. Provide a gateway to a pod application. Services are frontends for pods.

apiVersion: v1   
kind: Service   
metadata:     
  name: my-blog     
  labels:       
    app: blog   spec:     selector:       app: blog       server: nginx     ports:     - protocol: TCP       port: 80       targetPort: 80