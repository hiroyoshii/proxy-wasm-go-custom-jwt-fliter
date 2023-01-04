# overview
## istio demo app is deployed
- install istio
```
curl -L https://istio.io/downloadIstio | sh -
```
- install
```
istioctl install
```
- install demo applications
```
kubectl apply -n bookinfo -f https://raw.githubusercontent.com/istio/istio/release-1.16/samples/bookinfo/platform/kube/bookinfo.yaml
kubectl apply -n bookinfo -f https://raw.githubusercontent.com/istio/istio/release-1.16/samples/bookinfo/networking/bookinfo-gateway.yaml
```
- check apps
```
$ kubectl get pods -n bookinfo
NAME                              READY   STATUS    RESTARTS   AGE
details-v1-7b45484c56-tvmmq       2/2     Running   0          9h
productpage-v1-77f8f885b6-kqb2b   2/2     Running   0          9h
ratings-v1-67b4c86764-75kmw       2/2     Running   0          9h
reviews-v1-c6596c68b-56hgr        2/2     Running   0          9h
reviews-v2-6cf7688ffb-k2q5n       2/2     Running   0          9h
reviews-v3-69f7569f74-bx9hd       2/2     Running   0          9h
```
- istio-system
```
$ kubectl get pods -n istio-system
NAME                                   READY   STATUS    RESTARTS   AGE
istio-egressgateway-558869f64-rmb8q    1/1     Running   0          95m
istio-ingressgateway-774f5575f-t4ccq   1/1     Running   0          95m
istiod-767b964b4d-4xv7t                1/1     Running   0          95m
```

# deploy wasm proxy
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
