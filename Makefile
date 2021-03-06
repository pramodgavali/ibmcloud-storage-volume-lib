#/*******************************************************************************
# * Licensed Materials - Property of IBM
# * IBM Cloud Container Service, 5737-D43
# * (C) Copyright IBM Corp. 2017, 2018 All Rights Reserved.
# * US Government Users Restricted Rights - Use, duplication or
# * disclosure restricted by GSA ADP Schedule Contract with IBM Corp.
# ******************************************************************************/

GOPACKAGES=$(shell go list ./... | grep -v /vendor/) # With glide: GOPACKAGES=$(shell glide novendor)
GOFILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all
all: deps fmt vet test

.PHONY: deps
deps:
	glide install
	go get github.com/pierrre/gotestcover

.PHONY: fmt
fmt:
	gofmt -l ${GOFILES}
	@if [ -n "$$(gofmt -l ${GOFILES})" ]; then echo 'Above Files needs gofmt fixes. Please run gofmt -l -w on your code.' && exit 1; fi

.PHONY: test
test:
	$(GOPATH)/bin/gotestcover -v -race -coverprofile=cover.out ${GOPACKAGES}

.PHONY: coverage
coverage:
	go tool cover -html=cover.out -o=cover.html

.PHONY: vet
vet:
	go vet ${GOPACKAGES}

.PHONY: build
build:
	go build -gcflags '-N -l' -o libSample . #go build main.go

.PHONY: clean
clean:
	rm -rf libSample
