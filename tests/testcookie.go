package test

import (
	"io"
	"net/http"
)

func Cookie(w http.ResponseWriter, req *http.Request) {
	ck := &http.Cookie{
		Name:   "Cookie",
		Value:  "test",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}
	http.SetCookie(w, ck)
	ck2, err := req.Cookie("Cookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, ck2.Value)
}

func main() {
	http.HandleFunc("/", Cookie2)
	http.ListenAndServe(":8080", nil)
}

func Cookie2(w http.ResponseWriter, req *http.Request) {
	ck := &http.Cookie{
		Name:   "Cookie",
		Value:  "hello",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}
	w.Header().Set("Set-Cookie", ck.String())
	ck2, err := req.Cookie("Cookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, ck2.Value)
}
