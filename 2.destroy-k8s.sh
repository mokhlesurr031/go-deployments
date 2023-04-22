#!/bin/zsh

echo "Deleting the cluster..."
echo "-----------------------"

kubectl delete -f deployment/k8s/client/3-ingress-client.yaml &&
kubectl delete -f deployment/k8s/client/2-service-client.yaml &&
kubectl delete -f deployment/k8s/client/1-deployment-client.yaml &&
kubectl delete -f deployment/k8s/client/0-namespace-client.yaml &&

kubectl delete -f deployment/k8s/psql/6-service-psql.yaml &&
kubectl delete -f deployment/k8s/psql/5-statefulset-psql.yaml &&
kubectl delete -f deployment/k8s/psql/4-persistent-volume-claim-psql.yaml &&
kubectl delete -f deployment/k8s/psql/3-persistent-volume-psql.yaml &&
kubectl delete -f deployment/k8s/psql/2-secret-psql.yaml &&
kubectl delete -f deployment/k8s/psql/1-configmap-psql.yaml &&
kubectl delete -f deployment/k8s/psql/0-namespace-psql.yaml &&

echo "Cluster resources have been deleted."
