<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> CRD Controller Lab

> Using KubeBuilder, create a sample ContainerSet CRD

A ContainerSet CRD defines an image and a number of replicas and ensures a
deployment is created in the current namespace with the image/replicas specified
by the CRD.

> NOTE! We will use the generated code here with no modifications, we will do
> the right thing in the next lab. For now the default generated code ensures
> an nginx deployment is created with a single replicas.

1. Define a sample resource

    ```shell
    kubebuilder create api --group workload --version v1alpha1 --kind ContainerSet
    ```

1. Install the CRD schema

    ```shell
    make install
    # Verify!
    kubectl get crd | grep containersets
    ```

1. Run the sample controller

    ```shell
    make run
    ```

1. Watch your local pod and deployment

    ```shell
    kubectl get po,deploy
    ```

1. Install the sample CRD

   ```shell
   kubectl apply -f config/samples
   # In your watch window you should see a new containerset-sample-deployment pod
   # and associated deployment
   ```

1. Delete your containerset and crd

    ```shell
    kubectl delete -f config/samples -f config/crds
    ```


<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)