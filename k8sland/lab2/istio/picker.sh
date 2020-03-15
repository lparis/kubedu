#!/bin/bash
while true; do
  clear
  echo "Hangman"
  http --print=b --pretty=colors $(minikube ip):31380/new_game
  # OR if no httpie...
  # curl $(minikube ip):31380/new_game
  sleep ${1:-2}
done
