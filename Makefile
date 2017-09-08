CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-catalog
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-catalog/flags
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-catalog/index
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-catalog/probe
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-catalog/record
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-catalog/utils
	cp *.go src/github.com/whosonfirst/go-whosonfirst-catalog
	cp flags/*.go src/github.com/whosonfirst/go-whosonfirst-catalog/flags
	cp index/*.go src/github.com/whosonfirst/go-whosonfirst-catalog/index
	cp probe/*.go src/github.com/whosonfirst/go-whosonfirst-catalog/probe
	cp record/*.go src/github.com/whosonfirst/go-whosonfirst-catalog/record
	cp utils/*.go src/github.com/whosonfirst/go-whosonfirst-catalog/utils
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-log"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-uri"
	# @GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-tile38"
	# @GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-pgis"

vendor-deps: rmdeps deps
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go
	go fmt flags/*.go
	go fmt index/*.go
	go fmt probe/*.go
	go fmt record/*.go
	go fmt *.go

bin: 	self
	@GOPATH=$(GOPATH) go build -o bin/wof-catalog-report cmd/wof-catalog-report.go
