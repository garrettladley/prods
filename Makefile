NODE_BIN := ./node_modules/.bin

.PHONY:build
build: gen-css gen-swag gen-templ
	@go build -tags dev -o bin/prods cmd/server/main.go

.PHONY:build-prod
build-prod: gen-css gen-swag gen-templ
	@go build -tags prod -o bin/prods cmd/server/main.go

.PHONY:run
run: build
	@./bin/prods

.PHONY: install
install: install-templ install-swag gen-swag gen-templ
	@go get ./...
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss


.PHONY: gen-css
gen-css:
	@$(NODE_BIN)/tailwindcss build -i internal/views/css/app.css -o cmd/server/public/styles.css --minify

.PHONY: watch-css
watch-css:
	@$(NODE_BIN)/tailwindcss -i internal/views/css/app.css -o cmd/server/public/styles.css --minify --watch

.PHONY: install-templ
install-templ:
	@go install github.com/a-h/templ/cmd/templ@latest

.PHONY: gen-templ
gen-templ:
	@templ generate

.PHONY: watch-templ
watch-templ:
	@templ generate --watch --proxy=http://127.0.0.1:8000

.PHONY: install-swag
install-swag:
	@go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: gen-swag
gen-swag:
	@swag init -g cmd/server/main.go

.PHONY: fmt-swag
fmt-swag:
	@swag fmt -g cmd/server/main.go

.PHONY: ci-scaffold
ci-scaffold:
	@mkdir -p cmd/server/public
	@echo "hello world" > cmd/server/public/hello.txt
