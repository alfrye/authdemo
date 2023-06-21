# authdemo
Authentication demo using emissary

## Build Demo Application
To build the demo app run make version='docker image version' build/demoapp

## Install Demo Application
The Demo application can be installed using helm.
helm install   --debug   --create-namespace   --wait   --atomic   -n demo  demoapp --set
image.tag="version of image tag to use" helm

## Emissary Ingress AuthService
The following is the configuration for setting up the emissary-ingress AuthService Module

```yaml
apiVersion: getambassador.io/v3alpha1
kind: AuthService
metadata:
  name: authentication
  namespace: api-gateway
spec:
  auth_service: "http://demoapp.demo:8585"
  # status_on_error:csuZ81IjHm
  #   code: 401
  # failure_mode_allow: false
  tls: "true"
  proto: http
  path_prefix: "/bfauth"
 
  allowed_request_headers: 
    - X-myheader
 
  allowed_authorization_headers:
  # - cookie
    - X-myheader2  
```