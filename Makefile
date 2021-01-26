BUILD := $(shell git rev-parse --short HEAD 2>/dev/null || echo "latest")
VERSION := $(shell git describe --tags 2>/dev/null || echo "latest")
IMAGE := $(shell basename "$(PWD)"):$(BUILD)

.PHONY: docker
docker:
	-docker build --label "version=$(VERSION)" --label "build=$(BUILD)" -t $(IMAGE) .

.PHONY: example
example: docker
	-docker run -it --rm --env-file .env \
		-v $(PWD)/yaml/example.yaml:/config.yaml \
		$(IMAGE) /config.yaml
