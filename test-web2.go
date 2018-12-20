package main

import (
  "fmt"
  "net/http"
  "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
   muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

func getHeaderValue(r *http.Request, key string) string {
    value := r.Header.Get(key)
    return value
}

func printSpan(r *http.Request) {
    if span, ok := tracer.SpanFromContext(r.Context()); ok {
      fmt.Print(span)
      fmt.Print("\n")
    }
}

func tagSpan(r *http.Request, key string, value string) {
    if span, ok := tracer.SpanFromContext(r.Context()); ok {
      fmt.Print(span)
      span.SetTag(key, value)
    }
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
      user_agent := getHeaderValue(r, "User-Agent")
      tagSpan(r, "user-agent", user_agent)
      printSpan(r)
      fmt.Print("Headers: %n", r.Header)
      resp, err := http.Get("http://localhost:8888/")
      if err != nil {
        fmt.Print(err)
      }
      fmt.Print(resp)
      w.Write([]byte("You hit me!\n"))
    }

func main() {
      tracer.Start(tracer.WithDebugMode(true), tracer.WithPrioritySampling())
      mux := muxtrace.NewRouter(muxtrace.WithServiceName("top-level"))
      mux.HandleFunc("/", IndexHandler)
      http.ListenAndServe(":9999", mux)
    }