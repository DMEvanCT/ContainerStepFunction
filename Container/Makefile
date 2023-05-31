.PHONY: build clean deploy

build:
	#dep ensure -v
	env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/state  main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock