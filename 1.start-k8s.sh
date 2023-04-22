#!/bin/zsh

echo "Starting the cluster..."
echo "-----------------------"

kubectl apply -f k8s/psql/0-namespace-psql.yaml && 
kubectl apply -f k8s/psql/1-configmap-psql.yaml &&
kubectl apply -f k8s/psql/2-secret-psql.yaml &&
kubectl apply -f k8s/psql/3-persistent-volume-psql.yaml &&
kubectl apply -f k8s/psql/4-persistent-volume-claim-psql.yaml &&
kubectl apply -f k8s/psql/5-statefulset-psql.yaml &&
kubectl apply -f k8s/psql/6-service-psql.yaml &&

echo "Postgresql statefulset is up and running"
echo "-----------------------"

kubectl apply -f k8s/client/0-namespace-client.yaml &&
kubectl apply -f k8s/client/1-deployment-client.yaml &&
kubectl apply -f k8s/client/2-service-client.yaml &&
kubectl apply -f k8s/client/3-ingress-client.yaml &&

echo "Client deployment is up and running"

echo "Starting port forwarding..."
echo "--------------------------"

# Start port-forwarding in the background
kubectl port-forward services/go-svc 8081:8081 &

echo "Port forwarding is running in the background. To stop, run: kill $!"

# Wait for user to terminate the port-forwarding
echo "Press enter to terminate port-forwarding."
read

# Terminate the port-forwarding
kill %1

echo "Port forwarding terminated."
