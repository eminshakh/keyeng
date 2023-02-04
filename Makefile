run-mac: build
	./keyeng

run-linux: build
	./keyeng

build:
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o keyeng cmd/main.go

build_linux:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o keyeng cmd/main.go