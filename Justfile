set positional-arguments

alias u := update

default:
  @just --list

run:
    go run main.go

test:
	go test ./...

build:
	go build .

update:
	go get -u
	go mod tidy

fmt:
	go fmt ./...

tidy:
	go mod tidy