.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/upload upload/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/download download/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
