<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>


<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> Init Container Lab

> Leverage Init Containers to provision dictionaries for a Dictionary service.

The dictionary service will load dictionary data from a given asset directory and
dictionary name mounted on a volume. Use an init-container to provision the
volume with a set of dictionaries by cloning a dictionary assets repo.

1. Define a pod using the following Docker image:
   k8sland/dictionary-svc-go:0.0.2
1. The dictionary service takes in an asset directory containing word dictionaries
   and a dictionary filename ie trip_or_treat.txt.
   The service uses the following command:

   ```shell
   /app/dictionary -a dictionary_dir -d dictionary_name
   ```

1. This service runs on port 4000 and exposes /words endpoint to list out the words
   contained in the dictionary loaded via *-a/d* options above.
1. Define an init container to provision a volume to be used by the dictionary
   service. The init container needs to clone this repo
   [Dictionaries](https://github.com/k8sland/dictionaries.git) in order to
   provision the volume
1. Next change the init container command to cause the pod to fail
1. What's happening with your dictionary pod?
1. Delete you pod!

<br/>

---
## Commands

1. Launch your pod

    ```shell
    kubectl apply -f dictionary.yml
    ```

2. Verify the init container is successful and pod is launched

    ```shell
    kubectl get po
    ```

3. Verify the volume was provisioned correctly

    ```shell
    kubectl exec -it dictionary -- wget -q -O - http://localhost:4000/words
    ```

4. Change git url so that it does not resolve

    ```shell
    kubectl delete -f dictionary.yml --force --grace-period=0
    kubectl apply -f dictionary.yml
    ```

<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)