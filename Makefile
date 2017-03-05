
.PHONY: all vendor

all: vendor
vendor:
	go get -u github.com/kardianos/govendor
