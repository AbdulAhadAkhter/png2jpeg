build:
	@go build -o ./dist/png2jpeg	./main.go
clean: 
	rm -rf test/jpeg/ && rm -rf dist/