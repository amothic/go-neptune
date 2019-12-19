.DEFAULT_GOAL := help

## Run test
test:
	@go test -v ./...

## Launch a graph database
graph:
	@docker run --rm -d -p 8182:8182 amothic/neptune

## Prints the synopsis and a list of commands
help:
	@docker run --rm -v `pwd`:/make2help amothic/make2help $(MAKEFILE_LIST)
