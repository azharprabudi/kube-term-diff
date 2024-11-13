OS 								:= $(shell uname -s | tr '[:upper:]' '[:lower:]')
OS_ARCH						:= $(shell uname -m)
BIN_FOLDER				:= "bin"
APP_NAME					:= kube-term-diff

build-darwin-arm64: export GOOS=darwin
build-darwin-arm64: export GOARCH=arm64
build-darwin-arm64: export CGO_ENABLED=0
build-darwin-arm64:
	@go build \
		-o ${BIN_FOLDER}/${APP_NAME}_$${GOOS}_$${GOARCH} \
		-ldflags "-s -w" \
		*.go

build-darwin-amd64: export GOOS=darwin
build-darwin-amd64: export GOARCH=amd64
build-darwin-amd64: export CGO_ENABLED=0
build-darwin-amd64:
	@go build \
		-o ${BIN_FOLDER}/${APP_NAME}_$${GOOS}_$${GOARCH} \
		-ldflags "-s -w" \
		*.go

build-linux-arm64: export GOOS=linux
build-linux-arm64: export GOARCH=arm64
build-linux-arm64: export CGO_ENABLED=0
build-linux-arm64:
	@go build \
		-o ${BIN_FOLDER}/${APP_NAME}_$${GOOS}_$${GOARCH} \
		-ldflags "-s -w" \
		*.go

build-linux-amd64: export GOOS=linux
build-linux-amd64: export GOARCH=amd64
build-linux-amd64: export CGO_ENABLED=0
build-linux-amd64:
	@go build \
		-o ${BIN_FOLDER}/${APP_NAME}_$${GOOS}_$${GOARCH} \
		-ldflags "-s -w" \
		*.go

install:
	@go mod download

.PHONY: install build
build:
	@${MAKE} build-darwin-arm64
	@${MAKE} build-darwin-amd64
	@${MAKE} build-linux-arm64
	@${MAKE} build-linux-amd64
	@cp ${BIN_FOLDER}/${APP_NAME}_${OS}_${OS_ARCH} ${HOME}/go/bin/${APP_NAME}