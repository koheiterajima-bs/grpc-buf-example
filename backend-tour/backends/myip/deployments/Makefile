.DEFAULT_GOAL := help

.PHONY: help publish generate .ko_name

KO_DOCKER_REPO ?= crmhv25ua3ddectf7dg.azurecr.io

GO_CMD = ../cmd/myipapis

# gitのtagが付いていたら、それベースでimageにtagをつける
# push前に必ずtagをつける
TAG_PREFIX  ?= backend-tour/v
TAG_VERSION ?= $(shell git tag --sort=-v:refname|grep $(TAG_PREFIX)|head -1|sed -e 's|^$(TAG_PREFIX)||')
TAG_COUNT   ?= $(shell git log --pretty=oneline $(TAG_PREFIX)$(TAG_VERSION)..HEAD|wc -l|bc)

DEPLOY_NAME	= $(shell date +'deploy-%d%H%M%S')

help: ## show this message
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-8s\033[0m %s\n", $$1, $$2}'

build: ## ko build --push=false
	$(MAKE) publish KO_DOCKER_REPO=ko.local

publish: ## ko build --push=true ...
	env SOURCE_DATE_EPOCH=$$(git log -1 --format="%ct") KO_DOCKER_REPO=$(KO_DOCKER_REPO) ko build -B -t $(TAG_VERSION)-$(TAG_COUNT) -t latest $(GO_CMD)

show-acr: ## show acr
	az acr repository show-tags -n $(KO_DOCKER_REPO) --repository $(shell basename $(GO_CMD)) --detail

show-aca: ## show aca
	az containerapp show  -n $(ACA_NAME) -g $(ACA_BACKEND_RGNAME) -o table

deploy: ## deploy to azure
	az deployment sub create \
		--location japaneast \
		--name patientapis-$(DEPLOY_NAME) \
		--template-file ./main.bicep \
		--parameters \
			name=$(ACA_NAME) \
			acrName=$(ACR_NAME) \
			acaRgName=$(ACA_BACKEND_RGNAME) \
			envName=$(ACA_BACKEND_ENVNAME) \
			managedIdName=$(ACA_BACKEND_MSIDNAME) \
			imageName=$(shell basename $(GO_CMD)):$(TAG_VERSION)-$(TAG_COUNT)
