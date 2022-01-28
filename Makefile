HOSTNAME=olukotun-ts
NAME=circleci-terraform-provider

default: install

build:
	go build -o ${NAME}

fmt:
	gofmt -s -w $(find . -name '*.go')

lint:
	golangci-lint run ./...

install: build
	mkdir -p ~/terraform.d/plugins/olukotun-ts/circleci-terraform-provider
	~/terraform.d/plugins/${HOSTNAME}/${NAME}

test:
	go test ./...
