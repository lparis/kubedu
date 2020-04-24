<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> Custom Scheduler Lab

> In the spirit of Halloween, write a party scheduler that will only schedule
> a pod if it has a costume.

1. Define a custom party scheduler aka partysched, that check if a given pod
   has a costume label set to either ghoul or goblin or your choice of costumes...
1. Make sure your pod gets scheduled if it has a correct costume label
1. For all other cases ie no label or no correct attire, make sure your scheduler
   selected no nodes and spews out a log.
1. Setup a partysched deployment with a service account and RBAC rules
1. Using the given Makefile publish your party scheduler as a Docker container
1. Deploy your party scheduler in your cluster
1. Rinse and repeat your checks in a deployed configuration
1. Delete your application and scheduler!

## Commands

### Run your party scheduler manually

    ```shell
    go run main.go
    ```

### Provision deployment

    ```shell
    kubectl apply -f k8s/nginx.yml
    ```

### Build and deploy a Docker image

  > NOTE! Update the Makefile and setup for your DockerHub registry

    ```shell
    make img
    ```

### Deploy your scheduler and nginx deployment

    ```shell
    kubectl apply -f k8s
    ```

# Check RBAC Rules for serviceaccount

    ```shell
    ku auth can-i get pods -n defaut --as system:serviceaccount:default:partysched
    ```

<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
