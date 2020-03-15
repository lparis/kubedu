<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> Admission Controller Lab

> Trick or Treat? Provision your cluster with a dynamic admission controller
> that rejects all Grim Reaper's deployments!

1. Update the code in main.go to reject all grim reaper deployments. ie
   check for deployment resource with a label app=grim reaper.
1. Generate certs and keys for your webhook and make sure to provide the correct
   service name in your certificate encryption
1. Generate a ca_bundle to validate api-server callback into your webhook.
1. Set the ca_bundle in k8s/adm.yml to the output of the previous command
1. Deploy your admission controller
1. Register your validating admission controller with the api-server
1. Update the provided grim deployment with the magic label and provision your cluster
1. Trace the api-server logs and your admission controller logs to ensure all
   are working nominally.
1. Verify your deployment was denied for the right reason
1. Delete your deployment and change/remove the label
1. Redeploy and make sure your admission controller allows the new deployment
1. Delete your application!

---
## Commands

### Build your admission controller Docker image

    ```shell
    make img
    ```

### Deploy custom admission controller

    ```shell
    kubectl apply -f k8s/dp.yml
    ```

### Generate your certificates

    ```shell
    ./gen.sh
    ```

### Verify your certificate against your webhook

    ```shell
    openssl s_client -connect $(minikube ip):30443/ -CAfile caCert.pem
    ```

### Base64 encode your certificate

    ```shell
    cat caCert.pem | base64 | tr -d '\n'
    ```

# Register with api-server

    ```shell
    kubectl apply -f k8s/adm.yml
    ```

# Create a new `Grim Reaper deployment

    ```shell
    kubectl apply -f k8s/grim.yml
    ```


<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)