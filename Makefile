DRIVE_PREFIX?=
ifeq ($(OS), Windows_NT)
	DRIVE_PREFIX=C:
endif
.PHONY: spec
spec: ## generate spec.md file
	$(eval $@_TMP_OUT := $(shell mktemp -d -t composespec-output.XXXXXXXXXX))
	docker buildx build . \
	--output type=local,dest=$($@_TMP_OUT) \
	-f ./Dockerfile \
	--target spec-update
	rm -f spec.md
	cp -R "$(DRIVE_PREFIX)$($@_TMP_OUT)"/out/spec.md ./spec.md
	rm -rf "$(DRIVE_PREFIX)$($@_TMP_OUT)"/*

.PHONY: validate-spec
validate-spec: ## validate the spec.md does not change
	@docker buildx build . \
	-f ./Dockerfile \
	--target spec-validate

.PHONY: test-spec
test-spec: ## Run tests against compose-spec.json schema
	@docker buildx build . \
	-f ./Dockerfile \
	--target spec-test
