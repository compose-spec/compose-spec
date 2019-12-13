.PHONY: docs

docs:
	DOCKER_BUILDKIT=1 docker build --target docs  --output docs .
