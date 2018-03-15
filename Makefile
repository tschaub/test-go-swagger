MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash
.SHELLFLAGS := -o pipefail -euc
.DEFAULT_GOAL := generate

SWAGGER := docker run --rm -it -e GOPATH=$(HOME)/go:/go -v $(HOME):$(HOME) -w $(shell pwd) quay.io/goswagger/swagger:0.13.0
GEN := ./gen
SPEC := ./api/spec.yml

.PHONY: validate
validate:
	@$(SWAGGER) validate $(SPEC)

.PHONY: generate
generate: validate
	@mkdir -p $(GEN)
	@$(SWAGGER) generate server --spec=$(SPEC) --target=$(GEN) --exclude-main

.PHONY: clean
clean:
	@rm -rf $(GEN)