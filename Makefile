all:
	@go build

format:
	@go fmt

clean:
	@git clean -dffxq
