# deploy command
- add webassemblyhub registry secret
```
kubectl create secret docker-registry webassemblyhub-secret --docker-server=webassemblyhub-secret --docker-username=<user name> --docker-password=<password> -n istio-system
```
- deploy WasmPlugin
```
kubectl apply -f istio/wasmplugin.yaml
```
- deploy ServiceEntry
```
kubectl apply -f istio/serviceentry.yaml
```
