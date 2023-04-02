package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/iptables-viz/iptables-viz/backend/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)



func main() {
	port := os.Args[1]
	convertedPort, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		fmt.Println("error in parsing port number: ", err)
		return
	}
	r := mux.NewRouter()
	r.HandleFunc("/iptables/health", handler.HealthCheck)
	r.HandleFunc("/iptables/kubernetes", handler.GetKubernetesDefault).Methods("GET")
	r.HandleFunc("/iptables/kubernetes/{pod}/{table}", handler.GetKubernetesPodIptablesOutput).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(handler.DefaultHandler)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	err = http.ListenAndServe(fmt.Sprintf(":%d", convertedPort), handlers.CORS(originsOk, headersOk, methodsOk)(r))
	if err != nil {
		fmt.Println("error in starting the server: ", err)
		return
	}
}
