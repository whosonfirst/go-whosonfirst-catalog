CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-catalog
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-catalog/flags
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-catalog/http
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-catalog/index
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-catalog/probe
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-catalog/record
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-catalog/utils
	cp *.go src/github.com/whosonfirst/go-whosonfirst-catalog
	cp flags/*.go src/github.com/whosonfirst/go-whosonfirst-catalog/flags
	cp http/*.go src/github.com/whosonfirst/go-whosonfirst-catalog/http
	cp index/*.go src/github.com/whosonfirst/go-whosonfirst-catalog/index
	cp probe/*.go src/github.com/whosonfirst/go-whosonfirst-catalog/probe
	cp record/*.go src/github.com/whosonfirst/go-whosonfirst-catalog/record
	cp utils/*.go src/github.com/whosonfirst/go-whosonfirst-catalog/utils
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/facebookgo/grace/gracehttp"
	@GOPATH=$(GOPATH) go get -u "github.com/jteeuwen/go-bindata/"
	@GOPATH=$(GOPATH) go get -u "github.com/elazarl/go-bindata-assetfs/"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-mapzen-js"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-hash"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-log"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-uri"
	# @GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-tile38"
	# @GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-pgis"
	rm -rf vendor/github.com/jteeuwen/go-bindata/testdata

vendor-deps: rmdeps deps
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go
	go fmt flags/*.go
	go fmt http/*.go
	go fmt index/*.go
	go fmt probe/*.go
	go fmt record/*.go
	go fmt *.go

assets:	self
	@GOPATH=$(GOPATH) go build -o bin/go-bindata ./vendor/github.com/jteeuwen/go-bindata/go-bindata/
	@GOPATH=$(GOPATH) go build -o bin/go-bindata-assetfs vendor/github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs/main.go
	rm -f www/*~ www/css/*~ www/javascript/*~
	@PATH=$(PATH):$(CWD)/bin bin/go-bindata-assetfs -pkg http www www/javascript www/css
	mv bindata_assetfs.go http/assets.go

bin: 	self
	@GOPATH=$(GOPATH) go build -o bin/wof-catalog-report cmd/wof-catalog-report.go
	@GOPATH=$(GOPATH) go build -o bin/wof-catalog-report-server cmd/wof-catalog-report-server.go

maps: wwwdirs mapzenjs tangram refill

wwwdirs:
	if test ! -d www/javascript; then mkdir www/javascript; fi
	if test ! -d www/css; then mkdir www/css; fi
	if test ! -d www/tangram; then mkdir www/tangram; fi

tangram:
	curl -s -o www/javascript/tangram.js https://mapzen.com/tangram/tangram.debug.js
	curl -s -o www/javascript/tangram.min.js https://mapzen.com/tangram/tangram.min.js

refill:
	curl -s -o www/tangram/refill-style.zip https://mapzen.com/carto/refill-style/refill-style.zip

mapzenjs:
	curl -s -o www/css/mapzen.js.css https://mapzen.com/js/mapzen.css
	curl -s -o www/javascript/mapzen.js https://mapzen.com/js/mapzen.js
	curl -s -o www/javascript/mapzen.min.js https://mapzen.com/js/mapzen.min.js
