.PHONY: generate-docs
generate-docs:
	swag fmt
	swag init -g server.go --generatedTime --requiredByDefault=false


.PHONY: test-all
test-all:
	ginkgo ./...

.PHONY: test-api
test-api:
	ginkgo api/...



.PHONY: test-pkg
test-pkg:
	ginkgo pkg/...

