GOCMD=go
GOTEST=$(GOCMD) test
GOCOVER=$(GOCMD) tool cover
GOFMT=gofmt
GOFILES=$(shell find . -type f -name '*.go')

.DEFAULT_GOAL := all

.PHONY: all
all: check-fmt test coverage

.PHONY: test
test:
	$(GOTEST) -v ./... -coverprofile=./coverage.out

.PHONY: coverage
coverage:
	$(GOCOVER) -func=./coverage.out

.PHONY: check-fmt
check-fmt:
	@$(GOFMT) -d ${GOFILES}

.PHONY: check-fmt-list
check-fmt-list:
	@$(GOFMT) -l ${GOFILES}

.PHONY: fmt
fmt:
	$(GOFMT) -w ${GOFILES}
