BUILD_IMAGE=compose-spec-builder:latest
DOCKER_CMD=docker run -it --rm -v $(PWD):/compose-spec \
	-w /compose-spec \
	--privileged \
	$(BUILD_IMAGE)

.PHONY: builder
builder:
	docker build -t $(BUILD_IMAGE) - <Dockerfile.builder

.PHONY: docs
docs: builder
	$(DOCKER_CMD) jsonschema2md -d schema/config_schema_v3.9.json -o docs -v 04
	rm docs/config_schema_v3.9.json
