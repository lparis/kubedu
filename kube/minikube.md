- minkube -

kubernetes distribution on one vm (1 master and 1 worker)

- install minikube -

brew cask install minikube

https://github.com/kubernetes/minikube/releases

- run minikube -

minikube start

minibube stop

minikube dashboard

-- notes -- 

H4RX
http://tinyurl.com/vmheptio

docker images
docker ps



vmware.kubescale.com/intro
vmware.kubesclae.com/intro2

heptio/H3pt10

lab-info

ssh ubuntu@ec2-18-195-88-204.eu-central-1.compute.amazonaws.com

H3pt10

https://kubernetes.io/docs/reference/kubectl/cheatsheet/

kubeclt get deployment

kubectl describe <deployment>

kubectl describe pod <pod-name>

kubectl describe events

---

https://app.strigo.io/event/6LwW7R9ke5RXGY5Eb

https://kubernetes.io/docs/reference/kubectl/cheatsheet/
https://kubernetes.io/docs/reference/#api-reference

http://vmware.kubescale.com/intro/instructions/lab01/lab01.html
http://vmware.kubescale.com/intro/instructions/lab02/lab02.html


http://vmware.kubescale.com/intro/index.html

---

Kubenetes Service types

ClusterIP (default) -- Only accessible within the cluster (internal only)
NodePort -- Allows external access from the outside to worker nodes (port opened on each worker node)
LoadBalancer -- Alloses external access (requires LB)
External -- USed if service is fronting an app that sits outside of the cluster

