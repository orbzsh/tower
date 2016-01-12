package test

import (
	"io"
	"log"
	"net/http"
	"os"
)

func test() {
	mux := http.NewServeMux()
	mux.Handle("/", &indexHandler{})
	mux.HandleFunc("/hi", hi)
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(wd))))
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type indexHandler struct{}

func (*indexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "URL: "+req.URL.String())
}
func hi(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world")
}
