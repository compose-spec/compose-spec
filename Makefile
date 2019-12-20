.PHONY: docs,golang

docs:
	DOCKER_BUILDKIT=1 docker build --target docs  --output docs .

golang:
	DOCKER_BUILDKIT=1 docker build --target golang  --output . .
