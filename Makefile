BIN = cf-search
CLIBIN = cf-search

all:
	@templ generate
	@go build -o bin/$(BIN) cmd/web/*

cli-build:
	@go build -o bin/$(CLIBIN) cmd/cli/*

run: all
	@./bin/$(BIN)

run-cli: cli-build
	@./bin/$(CLIBIN)
