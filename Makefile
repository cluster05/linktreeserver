.PHONY: generate-docs
generate-docs:
	swag fmt
	swag init -g server.go --generatedTime --requiredByDefault=false
