package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
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
    }
}

func tagSpan(r *http.Request, key string, value string) {
    if span, ok := tracer.SpanFromContext(r.Context()); ok {
      fmt.Print(span)
      span.SetTag(key, value)
    }
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
      fmt.Printf("hello\n")
      fmt.Print(r.Header)
      user_agent := getHeaderValue(r, "User-Agent")
      tagSpan(r, "user-agent", user_agent)
	    w.Write([]byte("You hit me!\n"))
    }

func GetWord(w http.ResponseWriter, r *http.Request) {
      vars := mux.Vars(r)
      w.WriteHeader(http.StatusOK)
      w.Write([]byte(fmt.Sprintf("Word: %v\n", vars["word"])))
    }

func GetID(w http.ResponseWriter, r *http.Request) {
      vars := mux.Vars(r)
      w.WriteHeader(http.StatusOK)
      printSpan(r)
      w.Write([]byte(fmt.Sprintf("ID: %v\n", vars["id"])))
    }

func Test(w http.ResponseWriter, r *http.Request) {
      printSpan(r)
      resp, err := http.Get("http://localhost:8081/test2")
      fmt.Print(resp.Header)
      w.Write([]byte(fmt.Sprintf("Body: %v\n", body)))
    }

func main() {
      tracer.Start(tracer.WithDebugMode(true), tracer.WithPrioritySampling())
      mux := muxtrace.NewRouter()
      mux.HandleFunc("/", IndexHandler)
      mux.HandleFunc("/words/{word}", GetWord)
      mux.HandleFunc("/ids/{id}", GetID)
      mux.HandleFunc("/test", Test)
      http.ListenAndServe(":8888", mux)
    }