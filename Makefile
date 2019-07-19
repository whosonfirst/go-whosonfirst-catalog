# assets:	self
# 	go build -o bin/go-bindata ./vendor/github.com/jteeuwen/go-bindata/go-bindata/
# 	go build -o bin/go-bindata-assetfs vendor/github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs/main.go
# 	rm -f www/*~ www/css/*~ www/javascript/*~
# 	@PATH=$(PATH):$(CWD)/bin bin/go-bindata-assetfs -pkg http www www/javascript www/css
# 	mv bindata_assetfs.go http/assets.go

tools:
	go build -o bin/wof-inspector cmd/wof-inspector.go
	go build -o bin/wof-inspectord cmd/wof-inspectord.go

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
