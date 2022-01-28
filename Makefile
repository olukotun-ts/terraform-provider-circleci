HOSTNAME=github.com
NAMESPACE=olukotun-ts
NAME=circleci
BINARY=terraform-provider-${NAME}
VERSION=0.0.1
OS_ARCH=darwin_amd64

default: 
	make install

build:
	go get -d ./...
	go build -o ${BINARY}

fmt:
	gofmt -s -w .

lint:
	golangci-lint run ./...

install: 
	make fmt
	make build
	mkdir -p ${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	go test ./...
