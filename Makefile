BUILD := $(shell git rev-parse --short HEAD 2>/dev/null || echo "latest")
VERSION := $(shell git describe --tags 2>/dev/null || echo "none")
IMAGE := $(shell basename "$(PWD)"):$(BUILD)
CONFIG_FILE := $(PWD)/config-example.yaml

.PHONY:
	run

docker:
	-docker build --label "version=$(VERSION)" --label "build=$(BUILD)" -t $(IMAGE) .

run:
	-docker run -i -t --rm --env-file .env -v "$(CONFIG_FILE):/config.yaml" $(IMAGE)
