
.PHONY: build
build:
	go build -o bin/app main.go

.PHONY: vet
vet:
	go vet ./...