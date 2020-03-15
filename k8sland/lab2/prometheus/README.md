<img src="../assets/k8sland.png" align="right" width="auto" height="128"/>

<br/>
<br/>


# <img src="../assets/lab.png" width="32" height="auto"/> Prometheus Hangman Lab

> We're going to play a hangman game. The game consist of a couple of
> services hangman and dictionary and a CLI to submit guesses. The hangman
> service queries the dictionary service to get a list of words for the guess
> word. The code is already implemented and deployment manifests are in the k8s
> directory. To play the game, we are going to leverage Prometheus metrics to
> track good/bad guess counts as well as tracking win rates. Next, we are going
> to display the tally metrics in a Grafana dashboard. Sounds cool?

1. Instrument the hangman code base and add 2 prometheus counters to track your
   good and bad guesses (game.go).
2. Next define a prometheus gauge to track your game results:
   ie +1 for wins and -1 for loss (tally.go)
3. Before you get to play the game, your will need to tell Prometheus to
   track your hangman service by setting the ServiceMonitor CRD (k8s/prom/crd.yml)
4. Using the provided deployment templates, deploy Prometheus using the awesome
   CoreOS operator, Grafana and the hangman services namely dictionary and hangman.
5. Launch the Grafana UI.
6. You can now enjoy the fruits of your labor by firing off the hangman CLI and
   try out your guessing skills while watching your game performance in Grafana...
7. Delete all resources when done!

<br/>

---
## Commands

### Deploy Prometheus

```shell
kubectl apply -f k8s/prom
```

### Deploy Hangman

```shell
kubectl apply -f k8s/hangman
```

### Deploy Grafana

```shell
# Create ConfigMaps for datasource and dashboard
kubectl create cm prom-ds -n monitoring --from-file grafana/datasource.yml
kubectl create cm prom-dash -n monitoring --from-file grafana/dashboard.yml
kubectl create cm hm-dash -n monitoring --from-file grafana/dashboard.json
# NOTE! Grafana default credentials: admin/admin
kubectl apply -f k8s/grafana.yml
```

### Open Grafana

  ```shell
  minikube service -n monitoring grafana
  ```

### Play!

```shell
kubectl run -i --tty --rm hm --image k8sland/go-hangman-cli:0.0.1 \
  --command -- /app/hangman_cli --hm hangman:5000
```

<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
