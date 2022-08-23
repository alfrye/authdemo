.PHONY: all
all: k3d/up blueframe/start


.PHONY: build/demoapp
build/demoapp:
	cd examples/demoapp && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -gcflags -m -o build/demo cmd/main.go

.PHONY: build/simpleauth
build/simpleauth: 
	cd examples/simpleauth && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -gcflags -m -o build/simpleauth main.go


.PHONY: simpleauthDockerBuild
simpleauthDockerBuild:
	cd examples/simpleauth && docker push $$(docker build -t docker.io/alfrye/simpleauth:0.0.1 . | grep -oE 'Successfully tagged .+' | cut -d ' ' -f3  )



.PHONY: demoAppDockerBuild
demoAppDockerBuild:
	cd examples/demoapp && docker push $$(docker build -t docker.io/alfrye/demoapp:0.0.1 . | grep -oE 'Successfully tagged .+' | cut -d ' ' -f3  )
