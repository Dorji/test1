.PHONY: run
run:
	go run main.go
.PHONY: build
build:
	go build main.go
.PHONY: exec
exec:
	./main
.PHONY: mytest
mytest:
	go test -v
.PHONY: docker_build
docker_build:
	sudo docker build -t my-go-app .
.PHONY: docker_run
docker_run:
	sudo docker run -p 8080:8082 -it my-go-app
.PHONY: images
images:
	sudo docker images

DEFAULT_GOAL := run

