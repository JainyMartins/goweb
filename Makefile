//O arquivo makefile serve para criar scripts para executar o programa, por ex, rodar patchs, formatar arquivo, etc.
PACKAGES_PATH = $(shell go list -f '{{ .Dir }}' ./...)

.PHONY: start
start:
	@go run cmd/server/main.go