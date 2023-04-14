package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/iptables-viz/iptables-viz/backend/handler"
)

func main() {
	port := flag.Uint("port", 8080, "port number")
	platform := flag.String("platform", "linux", "platform for iptables visualization")
	flag.Parse()
	r := mux.NewRouter()
	if *platform == "kubernetes" {
		r.HandleFunc("/iptables/health", handler.HealthCheck)
		r.HandleFunc("/iptables/kubernetes", handler.GetKubernetesDefault).Methods("GET")
		r.HandleFunc("/iptables/kubernetes/{pod}/{table}", handler.GetKubernetesPodIptablesOutput).Methods("GET")
	} else {
		r.HandleFunc("/iptables/health", handler.HealthCheck)
		r.HandleFunc("/iptables/linux/{table}", handler.GetLinuxIptableOutput).Methods("GET")
	}
	r.NotFoundHandler = http.HandlerFunc(handler.DefaultHandler)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), handlers.CORS(originsOk, headersOk, methodsOk)(r)); err != nil {
		log.Fatalf("Error in starting the server, %s", err.Error())
	}
}
