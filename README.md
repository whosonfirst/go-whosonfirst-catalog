# go-whosonfirst-inspector

Who's On First datastore inspector.

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.7 so let's just assume you need [Go 1.9](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Important

Too soon. Move along.

## Usage

_Please write me_

## Tools

### wof-inspector

```
./bin/wof-inspector -h
Usage of ./bin/wof-inspector:
  -es value
    	The endpoint of an Elasticsearch database to inspect. Endpoints are defined as HOST + ':' + PORT + '#' + INDEX
  -fs value
    	The path of a Who's On First data directory to inspect. Paths are defined as ROOT + '#' + COMMA-SEPARATED LIST OF REPOSITORIES. If the value of list of repositories is '*' then all the repos in the 'whosonfirst-data' origanization will be used.
  -gh value
    	The name of a GitHub repos to inspect. If '*' then all the repos in the 'whosonfirst-data' organization will be used.
  -pg value
    	The DSN of a PostgreSQL endpoint to inspect.
  -s3 value
    	The name of an AWS S3 buckets to inspect.
  -t38 value
    	The endpoint of a Tile38 endpoints to inspect. Endpoints are defined as HOST + ':' + PORT + '#' + COMMA-SEPARATED LIST OF REPOSITORIES. If the value of list of repositories is '*' then all the repos in the 'whosonfirst-data' origanization will be used.
  -wof
    	Inspect records hosted on whosonfirst.mapzen.com/data.
```

### wof-inspectord

```
./bin/wof-inspectord -h
Usage of ./bin/wof-inspectord:
  -api-key string
    	A valid Mapzen API key (necessary for displaying maps). (default "mapzen-xxxxxxx")
  -es value
    	The endpoint of an Elasticsearch database to inspect. Endpoints are defined as HOST + ':' + PORT + '#' + INDEX
  -fs value
    	The path of a Who's On First data directory to inspect. Paths are defined as ROOT + '#' + COMMA-SEPARATED LIST OF REPOSITORIES. If the value of list of repositories is '*' then all the repos in the 'whosonfirst-data' origanization will be used.
  -gh value
    	The name of a GitHub repos to inspect. If '*' then all the repos in the 'whosonfirst-data' organization will be used.
  -host string
    	The hostname to listen for requests on. (default "localhost")
  -local string
    	The path to a local directory containing HTML (and CSS/Javascript) assets that the wof-inspectord daemon should serve.
  -pg value
    	The DSN of a PostgreSQL endpoint to inspect.
  -port int
    	The port number to listen for requests on. (default 8080)
  -root string
    	The root path to host the wof-inspectord daemon on. (default "/")
  -s3 value
    	The name of an AWS S3 buckets to inspect.
  -t38 value
    	The endpoint of a Tile38 endpoints to inspect. Endpoints are defined as HOST + ':' + PORT + '#' + COMMA-SEPARATED LIST OF REPOSITORIES. If the value of list of repositories is '*' then all the repos in the 'whosonfirst-data' origanization will be used.
  -wof
    	Inspect records hosted on whosonfirst.mapzen.com/data.
```

## See also

