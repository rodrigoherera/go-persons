GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=person

test:
		$(GOTEST) ./... -v

build:
		$(GOBUILD) -o $(BINARY_NAME) -v

vet:
		$(GOCMD) vet
cover: 
		$(GOTEST) ./... -race -coverprofile=coverage.txt -covermode=atomic