# overview
- try to implement JWT verification by proxy-wasm-go-sdk
  - verify JWT of GCP service account credential (generated by short lived credentials process)[https://cloud.google.com/iam/docs/create-short-lived-credentials-direct#sa-credentials-jwt]
- **this plugin does not work following reason**
  -  http call is dispatched, but the callback function is never invoked & the request continues to the upstream application.
    - https://github.com/tetratelabs/proxy-wasm-go-sdk/issues/278
  - if add blocing by sync.WaitGroup, time.Sleep or channel, following build error is occured.
    ```
    blocking operation in exported function: proxy_on_request_headers
    ```
  - JWT signature verification is not implemented because tinygo does not support crypto/rsa package.

# build & deploy wasm proxy image
## wasme 
- install command
```
curl -sL https://run.solo.io/wasme/install | sh
export PATH=$HOME/.wasme/bin:$PATH
```
- created by wasme command
```
wasme init custom_jwt_filter
```
- build
```
wasme build tinygo ./ -t webassemblyhub.io/hiroyoshii/sample:dev
```
- tagging
```
wasme tag webassemblyhub.io/hiroyoshii/sample:dev webassemblyhub.io/hiroyoshii/sample:v1.0
```
- push image
```
wasme push webassemblyhub.io/hiroyoshii/sample:v1.0
```