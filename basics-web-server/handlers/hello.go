package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body) // curl -v -d "Altamash Ali" localhost:9090

	if err != nil {
		http.Error(w, "oops..", http.StatusBadRequest)
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("OOOPS!.."))
		return
	}
	fmt.Fprintf(w, "Hello %s", d) // using response write to send response back to client
}
