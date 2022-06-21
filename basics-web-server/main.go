package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Altamashattari/microservices-with-go/basic-web-server/handlers"
)

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // add handler to default serve mux

	// })
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	// http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("Good Bye World")
	// })

	http.ListenAndServe(":9090", sm) // uses default serve mux - is an http handler
}
