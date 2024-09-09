.DEFAULT_GOAL := help

BUF_INPUT_MYIP ?= ../../api-specs/myipapis

# standard target

.PHONY: help

help:		## show this message
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-8s\033[0m %s\n", $$1, $$2}'

build: 		## ko build
build: build-server

test: 		## go test
	go test ./...

lint: ## staticcheck
	staticcheck ./...

# start extra target

generate:	## generate code
	rm -rf "pkg/apis/!('staticcheck.conf')"
	buf generate $(BUF_INPUT_MYIP)

tidy:		## go mod tidy
	go mod tidy

# internal terget

build-server:
	env SOURCE_DATE_EPOCH=$$(git log -1 --format="%ct") ko build --local -B ./cmd/myipapis

