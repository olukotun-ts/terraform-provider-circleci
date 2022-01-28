HOSTNAME=olukotun-ts
NAME=circleci-terraform-provider

default: 
	make install

build:
	go get -d ./...
	go build -o ${NAME}

fmt:
	gofmt -s -w .

lint:
	golangci-lint run ./...

install: 
	make fmt
	make build
	mkdir -p ~/terraform.d/plugins/olukotun-ts/circleci-terraform-provider
	mv ${NAME} ~/terraform.d/plugins/${HOSTNAME}/${NAME}

test:
	go test ./...
