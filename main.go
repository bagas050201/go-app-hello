package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    go_say_hello_web "github.com/bagas050201/go-say-hello-web/v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!\n"))
    w.Write([]byte(go_say_hello_web.SayHello("Bagas","jl. test 10")))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", Handler)
    log.Fatal(http.ListenAndServe(":8000", r))
}