# go-whosonfirst-inspector

Who's On First datastore inspector.

![](docs/images/20190724-wof-inspectord.png)

## Tools

To build binary versions of these tools run the `cli` Makefile target. For example:

```
> make cli
go build -mod vendor -o bin/wof-inspector cmd/wof-inspector/main.go
go build -mod vendor -o bin/wof-inspectord cmd/wof-inspectord/main.go
```

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

For example:

```
./bin/wof-inspector -wof 101736545 | python -mjson.tool
{
    "101736545": {
        "recordset": {
            "records": [
                {
                    "body": {
                        "bbox": [
                            -73.9475518701442,
                            45.41459067767231,
                            -73.47619754286481,
                            45.70379826163911
                        ],

                            "wof:population_rank": 12,
                            "wof:repo": "whosonfirst-data",
                            "wof:scale": "1",
                            "wof:superseded_by": [],
                            "wof:supersedes": [],
                            "wof:tags": []
                        },
                        "type": "Feature"
			...more stuff here
                    },
                    "hash": "e779be13d587889322fb2702a1e9f7be",
                    "id": 101736545,
                    "source": "whosonfirst",
                    "timing": 715966842,
                    "type": "geojson",
                    "uri": "https://whosonfirst.mapzen.com/data/101/736/545/101736545.geojson"
                }
            ]
        },
        "timings": 725273569
    }
}
```

### wof-inspectord

```
./bin/wof-inspectord -h
Usage of ./bin/wof-inspectord:
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
  -nextzen-api-key string
    	A valid Nextzen API key (necessary for displaying maps). (default "nextzen-xxxxxxx")
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

For example:

```
./bin/wof-inspectord -s3 data.whosonfirst.org -gh 'whosonfirst-data-admin-ca' -nextzen-api-key {APIKEY}
15:54:01.191617 listening on localhost:8080
```

And then if you inspected ID `101736545` you'd see something like this:

![](docs/images/20190724-wof-inspectord.png)

_Note: The UI/UX of `wof-inspectord` is in its very early days. Some things it doesn't do well and some things it can't do at all, yet._

## Docker

[Yes](Dockerfile). For example:

```
$> docker build -t whosonfirst-inspector .
$> docker run -it -p 8080:8080 whosonfirst-inspector /usr/local/bin/wof-inspectord -host 0.0.0.0 -gh whosonfirst-data-admin-ca -s3 data.whosonfirst.org
```
