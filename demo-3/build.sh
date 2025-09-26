#!/bin/bash

echo "Building Docker image for demo"
docker build -t oneformanew.azurecr.io/gindemo:latest .
echo "Pushing Docker image to ACR"
docker push oneformanew.azurecr.io/gindemo:latest
echo "Applying Kubernetes manifests"
kubectl.exe rollout restart deployment gin-demo -n demo