CWD=$(shell pwd)

assets:
	go build -o bin/go-bindata cmd/go-bindata/main.go
	go build -o bin/go-bindata-assetfs cmd/go-bindata-assetfs/main.go
	bin/go-bindata -o templates/templates.go -pkg templates -prefix static/templates/ static/templates/*.html
	rm -f static/*~ static/css/*~ static/javascript/*~
	@PATH=$(PATH):$(CWD)/bin bin/go-bindata-assetfs -o http/assets.go -pkg http static static/javascript static/css

tools:
	go build -mod vendor -o bin/wof-inspector cmd/wof-inspector/main.go
	go build -mod vendor -o bin/wof-inspectord cmd/wof-inspectord/main.go

maps: wwwdirs nextzenjs tangram refill

wwwdirs:
	if test ! -d www/javascript; then mkdir www/javascript; fi
	if test ! -d www/css; then mkdir www/css; fi
	if test ! -d www/tangram; then mkdir www/tangram; fi

tangram:
	curl -s -o www/javascript/tangram.js https://nextzen.org/tangram/tangram.debug.js
	curl -s -o www/javascript/tangram.min.js https://nextzen.org/tangram/tangram.min.js

refill:
	curl -s -o www/tangram/refill-style.zip https://nextzen.org/carto/refill-style/refill-style.zip

nextzenjs:
	curl -s -o www/css/nextzen.js.css https://nextzen.org/js/nextzen.css
	curl -s -o www/javascript/nextzen.js https://nextzen.org/js/nextzen.js
	curl -s -o www/javascript/nextzen.min.js https://nextzen.org/js/nextzen.min.js

docker:
	docker build -t whosonfirst-inspector .
