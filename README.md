# go-whosonfirst-catalog

Go package for working with names and RFC 5646 language tags in Who's On First documents.

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.6 so let's just assume you need [Go 1.8](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Important

Too soon. Move along.

## Example

### Simple

```
./bin/wof-catalog-report -s3 whosonfirst.mapzen.com -id 420546521 -gh whosonfirst-data | jq '.recordset.records[].body.properties["wof:parent_id"]'
-1
-1
```