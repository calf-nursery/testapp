
.PHONY: build
build:
	go build -o bin/app main.go

.PHONY: vet
vet:
	go vet ./...

.PHONY: generate-modules
generate-modules: 
	go mod tidy

.PHONY: verify
verify: generate-modules
	@if !(git diff --quiet HEAD -- go.sum go.mod); then \
		git diff; \
		echo "go module files are out of date"; exit 1; \
	fi