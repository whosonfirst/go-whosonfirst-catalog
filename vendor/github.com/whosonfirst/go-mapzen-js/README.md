# go-mapzen-js



## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.6 so let's just assume you need [Go 1.8](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Important

Too soon. Move along.

## Example

### MapzenJSHandler

```
import (
	mz "github.com/whosonfirst/go-mapzen-js/http"
)

func main(){

	mapzenjs_handler, err := mz.MapzenJSHandler()

	if err != nil {
		logger.Fatal("failed to create mapzen.js handler because %s", err)
	}

	mux := gohttp.NewServeMux()

	mux.Handle("/javascript/mapzen.js", mapzenjs_handler)
	mux.Handle("/javascript/mapzen.min.js", mapzenjs_handler)
	mux.Handle("/javascript/tangram.js", mapzenjs_handler)	
	mux.Handle("/javascript/tangram.min.js", mapzenjs_handler)
	mux.Handle("/css/mapzen.js.css", mapzenjs_handler)
	mux.Handle("/tangram/refill-style.zip", mapzenjs_handler)

}
```
