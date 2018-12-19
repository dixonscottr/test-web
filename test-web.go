package main

import (
  "fmt"
  "net/http"
  "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
  "github.com/gorilla/mux"
  "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
   muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

func getHeaderValue(r *http.Request, key string) string {
    value := r.Header.Get(key)
    return value
}

func tagSpan(r *http.Request, key string, value string) {
  if span, ok := tracer.SpanFromContext(r.Context()); ok {
    fmt.Print(span)
      span.SetTag(key, value)
    }
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
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

func CreateWord(w http.ResponseWriter, r *http.Request) {
      vars := mux.Vars(r)
      w.WriteHeader(http.StatusOK)
      w.Write([]byte(fmt.Sprintf("ID: %v\n", vars["id"])))
    }

func main() {
      tracer.Start(tracer.WithDebugMode(true), tracer.WithPrioritySampling())
	// r := mux.NewRouter()
	   // tracer.Start(tracer.WithServiceName("my-service"), tracer.WithDebugMode(true))
      mux := muxtrace.NewRouter()
            //r.HandleFunc("/", YourHandler)
      mux.HandleFunc("/", YourHandler)
      mux.HandleFunc("/words/{word}", GetWord)
      mux.HandleFunc("/ids/{id}", CreateWord)
      http.ListenAndServe(":8000", mux)
      fmt.Printf("Hello, World!\n")
	   //  defer tracer.Stop()
    }