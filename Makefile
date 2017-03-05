
.PHONY: all vendor install fmt test

all: vendor test install

vendor:
	go get -u github.com/kardianos/govendor
	cd src/github.com/vendor && govendor sync && cd ../../..

install:
	go install github.com/azohrabyan/oftask

fmt:
	go fmt github.com/azohrabyan/oftask/...

test:
	go test github.com/azohrabyan/oftask/...