# author:yaoyao
GOCMD=go
GOBUILD= $(GOCMD) build

all: build clear

build: clear build-apline docker-build docker-push

clean: clear

# linux cmd
clear:
	rm -rf server main

build-apline:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD)  -o server main.go

docker-build:
	docker build -t liuyao/go-pod-server:metrics .

docker-push:
	docker push liuyao/go-pod-server:metrics
