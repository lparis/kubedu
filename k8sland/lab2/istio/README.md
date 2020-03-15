<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>


# <img src="../assets/lab.png" width="32" height="auto"/> Istio Lab

> Let's play Hangman! Deploy a Hangman service to an Istio Cluster

> The hangman game is composed of 2 K8s services: dictionary and hangman. The hangman
> service fetches a list of words from the dictionary service and pick a random word to
> initialize the guessing game.


1. Install Istio

   1. Download
      As of this writing the latest version of istio is 1.0.2.

      ```shell
      mkdir ~/istio && cd ~/istio
      curl -L https://git.io/getLatestIstio | sh -
      # NOTE! Change this based on downloaded version
      cd istio-1.0.X
      # NOTE istio comes bundle with it's own cli aka istioctl.
      export PATH=$PWD/bin:$PATH
      ```

  1. Install **Istio** on your Minikube cluster

      ```shell
      cd ~/istio/istio-1.0.X/install/kubernetes
      # NOTE! This will take a while. Monitor all pods are running
      # before proceeding!
      kubectl apply -f istio-demo.yaml
      ```

1. Ensure all the Istio components are up and running in the *istio-system* namespace
1. Enable Sidecar injection in your default namespace:

    ```shell
    kubectl label ns default istio-injection=enabled
    ```

1. Deploy k8s/dictionary_v1 (trump words)
   1. Ensure your dictionary v1 is up, running and side-cared! ie Ready=[2/2]
1. Edit and deploy k8s/dictionary_svc
   1. The service must watch for pods with label app=dictionary
   2. Make sure the dictionary service is accessible on your node using port 30400
   3. Deploy your dictionary service
   4. Ensure this service found your dictionary v1 pod!
1. Deploy the Hangman service in k8s/hangman.yml
   1. Verify the hangman pod and service are correctly configured!
1. Play the game!

    ```shell
    # NOTE! Press enter once the pod is initialized!
    kubectl run -i --tty --rm hm --image k8sland/hangman-cli-go:0.0.1 --command -- /app/hangman_cli --url hangman:5000
    ```

1. Deploy k8s/dictionary_v2 (halloween words)
   1. Ensure the dictionary is up and running correctly
1. Edit istio/routes and complete the routes policy
1. Edit istio/subsets and complete the destination rule
1. Deploy your Istio gateway, routes and subsets manifests
1. Using the picker.sh script check the current hangman behavior

    ```shell
    # This should show 50% of traffic going to either v1 or v2
    ./picker.sh
    ```

1. Edit istio/dictionary-80-20 to route traffic 80% to v2 and 20% to v1
    1. Make sure to apply weighted routing when traffic is origination from hangman.
    2. Provision your new policy
    3. Ensure the VirtualService was created correctly
    4. Check the picker and make sure it produces more v2 words than v1's
1. Delete your weighted traffic policy!
    1. Ensure the picker is working correctly ie shows 50% words from v1 and v2
1. Next edit istio/mirror to mirror all traffic coming from v1 to v2
    1. Deploy your new policy
    2. Tail both v1 and v2 logs and make sure all traffic destined to v1 also
        hits the v2 version.
1. Delete the entire application including Istio!
1. Well done!

<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
