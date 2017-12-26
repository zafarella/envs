GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=envs
BINARY_UNIX=$(BINARY_NAME)_unix
COMMIT_HASH=-X main.GitCommit=$(shell git rev-parse --short HEAD)
BUILT_USER=-X main.BuiltUser=$(shell echo ${USER})
LDFLAGS=-ldflags "$(BUILT_USER) $(COMMIT_HASH)"

all: deps test build
build:
	$(GOBUILD) -o $(BINARY_NAME) $(LDFLAGS) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) -d ./...

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

darwin:
	cd ${BUILD_DIR};
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-darwin-${GOARCH} . ;
	cd - >/dev/null

windows:
	cd ${BUILD_DIR};
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-windows-${GOARCH}.exe . ;
	cd - >/dev/null
