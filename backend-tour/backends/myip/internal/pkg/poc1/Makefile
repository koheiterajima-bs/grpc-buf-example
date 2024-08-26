.DEFAULT_GOAL := help


# standard target

.PHONY: help

help:		## show this message
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-8s\033[0m %s\n", $$1, $$2}'

format:	## generate code
	buf format -w $(BUF_INPUT_FAXAPIS)
	buf format -w $(BUF_INPUT_EMAILAPIS)

lint:	## generate code
	buf lint $(BUF_INPUT_FAXAPIS)
	buf lint $(BUF_INPUT_EMAILAPIS)

generate:	## generate code
	rm -rf "pkg/apis/!('staticcheck.conf')"
	buf generate $(BUF_INPUT_FAXAPIS)
	buf generate $(BUF_INPUT_EMAILAPIS)
	buf generate 'https://github.com/envoyproxy/protoc-gen-validate.git#subdir=validate'

