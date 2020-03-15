# <img src="../assets/lab.png" width="32" height="auto"/> KubeBuilder Lab

> Deploy a sample CRD using KubeBuilder.

In this lab you are going to install and setup kubebuilder in your local dev
environment on minikube. You are going to define and provision a sample CRD
called containerset. A containerset allows you to define a deployment by merely
defining a CRD with a Docker image name and a replica count.

> NOTE! The current version of kubebuilder cans the image and replicaCount to be
> nginx and 1 which is ok for this initial lab.

> NOTE! Sadly, as of this writing, kubebuilder does not leverage GO modules.
> So if you are running GO >= 1.11, you will need to turn modules off as
> kubebuilder currently leverages dep for package management.

1. Install GO

    If you don't already have go running on your machine, install GO per
    [these instructions](https://golang.org/doc/install).

1. Install KubeBuilder

    ```shell
    export KB_ARCH=amd64
    export KB_REV=1.0.4

    # Download the release for your architecture
    cd Downloads
    curl -LO https://github.com/kubernetes-sigs/kubebuilder/releases/download/v${KB_REV}/kubebuilder_${KB_REV}_darwin_${KB_ARCH}.tar.gz
    # Extract the archive and setup directory
    tar -zxvf ~/Downloads/kubebuilder_${KB_REV}_darwin_${KB_ARCH}.tar.gz && sudo mv kubebuilder_${KB_REV}_darwin_${KB_ARCH} /usr/local/kubebuilder

    # update your PATH to include /usr/local/kubebuilder/bin
    # NOTE: We recommend setting this up in your terminal profile for your shell
    export PATH=$PATH:/usr/local/kubebuilder/bin

    # Verify!
    kubebuilder --help
    ```

1. Install Kustomize

    Kubebuilder requires [Kustomize](https://github.com/kubernetes-sigs/kustomize)
    to manipulate your deployment artifacts. Kustomize version must be > 1.0.4

    ```shell
    # For OSX install use homebrew
    brew install kustomize
    # For other platforms please see install instruction on the link above!
    ```

1. Setup a GO workspace

    ```shell
    mkdir -p $HOME/k8sland/src
    export GOPATH=$HOME/k8sland
    export GOBIN=$GOPATH/bin
    export GO111MODULE=off
    export PATH=$PATH:$GOBIN
    ```

1. Install Dep

    KubeBuilder currently leverages dep for dependency management. Please install
    Dep per these [installation instructions](https://github.com/golang/dep)

1. Initialize KubeBuilder and install dependencies

    > NOTE! This will take a while...

    ```shell
    kubebuilder init --domain k8sland.io --license apache2 --owner "K8sland Training"
    ```

1. Define a sample resource

    ```shell
    cd $GOPATH/src/github.com/k8sland.io/crds
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
   kubectl apply -f config/
   # In your watch window you should see a new containerset-sample-deployment pod
   # and associated container-set deployment
   ```

1. How cool was that?

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)