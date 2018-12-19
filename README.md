## Requirements:

* [dd-trace-go](https://github.com/DataDog/dd-trace-go) (gopkg.in/DataDog/dd-trace-go.v1/ddtrace)
* [Gorilla mux](https://github.com/gorilla/mux) (github.com/gorilla/mux)
* datadog-agent with APM enabled

## How to use:

* clone repo to `$GOPATH`
* `go run /path/to/test-web.go`
* `curl localhost:8888`
