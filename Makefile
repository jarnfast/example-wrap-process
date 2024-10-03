

.PHONY: build
build:
	@mkdir -p bin
	go build -race -o bin/example-wrap-process .
	
.PHONY: docker-buildx
docker-buildx:
	docker buildx build -t example-wrap-process-init:latest -f Dockerfile .
