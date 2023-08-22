# Kubernetes

## Knowledge

Determine the DNS name format:

The DNS name of a Kubernetes service is usually in the format: `<service-name>.<namespace>.svc.cluster.local`. Here:

- `<service-name>` is the name of your service.
- `<namespace>` is the Kubernetes namespace where the service is deployed. If your service is in the default namespace, you can omit the namespace from the DNS name.



## Commands


- List pods in a project:

```
kubectl -n <project> get pods
```
  
  
- Show logs for a pod:
```
kubectl -n <project> logs -f <pod> 
# If the pod has more than one container add "-c <container>"
```

- Restart pods:
```
kubectl -n <project> rollout restart deploy/<service>
```


- Scale Deploy:
```
kubectl -n <project> scale deploy <service> --replicas=<total replicas>
```

- Delete pods:
```
kubectl -n <project> delete pod <pod>
```


- Port Forward:
```
kubectl -n <project> port-forward svc/<service> <external port>:<internal port> 
```


- Get Pod events:
```
kubectl -n <project> get events --output json
```


- Get Log from the previous Pod:
```
kubectl -n <project> logs <pod> --previous 
```

- Describe resource:
```
kubectl -n <project> describe <resource-type> <resource-name>
```
