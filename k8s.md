# Kubernetes

List pods in a project:

```
kubectl -n <project> get pods
```
  
  
Show logs for a pod:
```
kubectl -n <project> logs -f <pod> 
# if pod has more than one container add "-c <container>"
```

Restart pods:

```
kubectl -n <project> rollout restart deploy/<service>
```


Scale Deploy:

```
kubectl -n <project> scale deploy <service> --replicas=<total replicas>
```

Delete pods:

```
kubectl -n <project> delete pod <pod>
```


Port Forward:

```
kubectl -n <project> port-forward svc/<service> <external port>:<internal port> 
```
