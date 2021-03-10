# Makefile to build push-notifications-go-sdk

build:
	go build ./...

runUnitTests:
	make build
	cd pushservicev1 && go test
	cd common && go test

tidy:
	go mod tidy