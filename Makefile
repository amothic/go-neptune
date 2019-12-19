test:
	@go test -v ./...

graph:
	@docker run --rm -d -p 8182:8182 amothic/neptune
