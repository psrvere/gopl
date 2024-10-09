package chapter1

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func WebServer1() {
	http.HandleFunc("/", handler1)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

var mu sync.Mutex
var count int

func WebServer2() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func WebServer3() {
	http.HandleFunc("/", handler3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func WebServer4() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
