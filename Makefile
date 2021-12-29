NAME=circleci
BINARY=terraform-provider-${NAME}
VERSION=0.0.1
OS_ARCH=darwin_amd64

default: build

build:
	go build -o ${BINARY}
