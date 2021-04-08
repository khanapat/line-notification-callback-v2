NAME?=line-notification

CONTAINER_IMAGE?=kcskbcnd93.kcs:5000/utility/${NAME}
VERSION?=$(shell git tag --points-at HEAD)

run:
	go run main.go

clean:
	rm -f goapp

test: clean
	go test -v -cover ./...

build: test
	docker build . --no-cache -t $(CONTAINER_IMAGE):$(VERSION) -f build/Dockerfile