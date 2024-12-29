NODE_BIN := ./node_modules/.bin

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^//'

.PHONY: confirm
confirm:
	@echo 'Are you sure? [y/N]' && read ans && [ $${ans:-N} = y ]

## build: build the application
.PHONY:build
build: gen/css gen/templ gen/swag
	@go build -tags dev -o bin/gml ./cmd/server

## build/prod: build the application for production
.PHONY:build/prod
build/prod: gen/css gen/templ gen/swag
	@go build -tags prod -o bin/gml ./cmd/server

## run: run the application
.PHONY:run
run: build
	@./bin/gml

## install: install dependencies
.PHONY: install
install:
	@make install/templ
	@make gen/templ
	@make install/go
	@make install/swag
	@make install/css

## install/go: install go dependencies
.PHONY: install/go
install/go:
	@go get ./...
	@go mod tidy
	@go mod download

## install/templ: install templ
.PHONY: install/templ
install/templ:
	@go install github.com/a-h/templ/cmd/templ@latest

## install/swag: install swag
.PHONY: install/swag
install/swag:
	@go install github.com/swaggo/swag/cmd/swag@latest

## install/css: install css
.PHONY: install/css
install/css:
	@npm install -D tailwindcss

## gen/css: generate css
.PHONY: gen/css
gen/css:
	@$(NODE_BIN)/tailwindcss build -i internal/views/css/app.css -o cmd/server/public/styles.css --minify

## gen/templ: generate templ
.PHONY: gen/templ
gen/templ:
	@templ generate

## gen/swag: generate swag
.PHONY: gen/swag
gen/swag:
	@swag init -g cmd/server/main.go

## watch/css: watch css
.PHONY: watch/css
watch/css:
	@$(NODE_BIN)/tailwindcss -i internal/views/css/app.css -o cmd/server/public/styles.css --minify --watch

## watch/templ: watch templ
.PHONY: watch/templ
watch/templ:
	@templ generate --watch --proxy=http://127.0.0.1:8000

## fmt/swag: format swag
.PHONY: fmt/swag
fmt/swag:
	@swag fmt -g cmd/server/main.go

## ci/scaffold: scaffold the project
.PHONY: ci/scaffold
ci/scaffold:
	@mkdir -p cmd/server/public
	@echo "hello world" > cmd/server/public/hello.txt
