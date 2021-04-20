NAME?=line-notification

CONTAINER_IMAGE?=kcskbcnd93.kcs:5000/utility/$(NAME)
VERSION?=$(shell git tag --points-at HEAD)
CACHE_IMAGE?=$$(docker images --filter "dangling=true" -q --no-trunc)

run:
	go run main.go

clean:
	rm -f goapp

test: clean
	go test -v -cover ./...

create: test
	docker build . --no-cache -t $(CONTAINER_IMAGE):$(VERSION) -f build/Dockerfile

push: create
	docker push $(CONTAINER_IMAGE):$(VERSION)
	docker rmi $(CACHE_IMAGE) $(CONTAINER_IMAGE):$(VERSION)