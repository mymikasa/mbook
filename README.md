# mbook


![edit.png](pic%2Fedit.png)
![profile.png](pic%2Fprofile.png)


kubectl apply -f k8s-mbook-deployment.yaml

kubectl apply -f k8s-mbook-service.yaml
kubectl apply -f k8s-mysql-service.yaml
kubectl apply -f k8s-redis-service.yaml
kubectl apply -f k8s-mysql-pv.yaml
kubectl apply -f k8s-mysql-pvc.yaml
kubectl apply -f k8s-mysql-deployment.yaml
kubectl apply -f k8s-mysql-service.yaml


kubectl delete deployment mbook-mysql
kubectl delete deployment mbook
kubectl delete service mbook-mysql
kubectl delete service mbook
kubectl delete pvc mbook-mysql-claim
kubectl delete pv my-local-pv


kubectl get deployment
kubectl get pv
kubectl get pvc
kubectl get service
kubectl get pods
```
rm mbookk || true
GOOS=linux GOARCH=arm go build -tags=k8s -o mbookk .
docker rmi -f mikasa/mbookk:v0.0.1
docker build -t mikasa/mbookk:v0.0.1 .
```

