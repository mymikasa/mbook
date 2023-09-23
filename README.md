# mbook


![edit.png](pic%2Fedit.png)
![profile.png](pic%2Fprofile.png)

kubectl delete deployment mbook
kubectl apply -f k8s-mbook-deployment.yaml

kubectl apply -f k8s-mbook-service.yaml
kubectl apply -f k8s-mysql-service.yaml
kubectl apply -f k8s-redis-service.yaml
kubectl get deployment

```
rm mbookk || true
GOOS=linux GOARCH=arm go build -tags=k8s -o mbookk .
docker rmi -f mikasa/mbookk:v0.0.1
docker build -t mikasa/mbookk:v0.0.1 .
```

