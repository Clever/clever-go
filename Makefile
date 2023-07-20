PKG := "github.com/Clever/clever-go/"
PWD := $(shell pwd)
PKG_NAME := ${PKG}
UNAME := $(shell uname | tr '[:upper:]' '[:lower:]')
SWAGGER_URL := $(shell curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/36578987 | jq -r '.assets[] | select(.name | contains("'"${UNAME}"'_amd64")) | .browser_download_url')
.PHONY: all 

all: deps test ## Make all

deps: ## Go modules download
	@go mod download

generate: ## Generate client and mock server from swagger.yml 
	@swagger generate client -f swagger.yml -A clever
	@swagger generate server -f swagger.yml -A clever

gdoc:  ## View Go Docs
	@echo "==> Running Local Go Docs"
	@echo ""
	@echo "Browse to: http://localhost:8181/pkg/${PKG}"
	@godoc -http=:8181

getswagger: ## Download go-swagger
	@curl -o /usr/local/bin/swagger -L'#' "${SWAGGER_URL}"
	@chmod +x /usr/local/bin/swagger

swaggerdoc:  ## View Swagger Docs
	@echo "==> Running Swagger Docs"
	@echo ""
	@echo "Browse to: http://localhost:8282"
	@docker run -p 8282:8080 -e SWAGGER_JSON=/swagger.yml -v ${PWD}/swagger.yml:/swagger.yml swaggerapi/swagger-ui

compiletest: ## Compiles test
	@go test -v  ./... -run XXxxxXXXxxx  # ensures tests compile before running

test: ## Tests the code
	@go test -v ./... -count=1

cover: ## Generates test coverage report
	@echo "==> Running go test coverage tools: "
	@echo " "
	@go test -v -coverprofile /tmp/${PKG_NAME}.coverage.out ./... || { echo ""; echo "======> Go Tests Failed"; return 1; }
	@echo ""
	@echo "===> Generating go test coverage report: "
	@echo ""
	@go tool cover -html=/tmp/${PKG_NAME}.coverage.out -o /tmp/${PKG_NAME}.coverage.htm || { echo ""; echo "======> Go Coverage Reports Failed"; return 1; }
	@echo ""
	@echo "===> Opening coverage report: "
	@echo ""
	@echo "Browse to: http://localhost:42280/${PKG_NAME}.coverage.htm"
	@cd /tmp && python -m SimpleHTTPServer 42280

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
