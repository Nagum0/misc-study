#!/bin/bash

# Start minikube without a CNI
minikube start \
  --network-plugin=cni \
  --cni=false \
  --memory=8192 \
  --cpus=4

# Install Calico
kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml

# Check Calico
kubectl get pods -n kube-system

# Install Istio into the minikube cluster
istioctl install --set profile=demo -y

# Enable Istio sidecar injection
kubectl label namespace demo4-app istio-injection=enabled

# Check Istio
kubectl get pods -n istio-system
