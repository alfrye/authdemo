# authdemo
Authentication demo using emissary

## Build Demo Application
To build the demo app run make version='docker image version' build/demoapp

## Install Demo Application
The Demo application can be installed using helm.
helm install   --debug   --create-namespace   --wait   --atomic   -n demo  demoapp --set
image.tag="version of image tag to use" helm
