BUILD := $(shell git rev-parse --short HEAD 2>/dev/null || echo "latest")
VERSION := $(shell git describe --tags 2>/dev/null || echo "latest")
IMAGE := $(shell basename "$(PWD)"):$(BUILD)

.PHONY: docker
docker:
	-docker build --label "version=$(VERSION)" --label "build=$(BUILD)" -t $(IMAGE) .

.PHONY: example
example: docker
	-docker run -it --rm --env-file .env \
		-v $(PWD)/config/attributes-example.yaml:/attributes.yaml \
		-v $(PWD)/config/metrics-example.yaml:/metrics.yaml \
		$(IMAGE) /metrics.yaml /attributes.yaml
