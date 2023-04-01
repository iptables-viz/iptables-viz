package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"backend/models"
	"backend/utility"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func GetDockerIptablesOutput(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	containerName := vars["container"]
	tableName := vars["table"]
	var resp models.Response
	cmd := exec.Command("bash", "-c", fmt.Sprintf("docker exec %s iptables -L -t %s | jc --iptables", containerName, tableName))
	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("error in listing Iptables chains: %v\n", err)
		return
	}

	resp.Entity = "docker"
	resp.IptableOutput = string(output)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		fmt.Println("error in encoding output: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetKubernetesPodIptablesOutput(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	kubeProxyPodName := vars["pod"]
	tableName := vars["table"]
	var resp models.Response
	output, err := utility.RunPodShellCommand(kubeProxyPodName, tableName)
	if err != nil {
		fmt.Println("error in running the shell command: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	resp.Entity = "kubernetes"
	resp.IptableOutput = output
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		fmt.Println("error in encoding output: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetKubernetesDefault(w http.ResponseWriter, r *http.Request) {
	var resp models.KubernetesDefaultResponse
	clientSet := utility.ClientSetup()
	podList, err := utility.GetPodList(clientSet)
	if err != nil {
		fmt.Println("error in getting pod list: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	kubeProxyPodName := podList[0]
	tableName := "nat"
	output, err := utility.RunPodShellCommand(kubeProxyPodName, tableName)
	if err != nil {
		fmt.Println("error in running the shell command: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	resp.PodName = kubeProxyPodName
	resp.IptableOutput = output
	resp.PodList = podList
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		fmt.Println("error in encoding output: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	health := make(map[string]string)
	health["now"] = now.Format(time.ANSIC)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("haha")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Invalid Access: requested route not found")
}

func main() {
	port := os.Args[1]
	convertedPort, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		fmt.Println("error in parsing port number: ", err)
		return
	}
	r := mux.NewRouter()
	r.HandleFunc("/iptables/health", healthCheck)
	r.HandleFunc("/iptables/kubernetes", GetKubernetesDefault).Methods("GET")
	r.HandleFunc("/iptables/kubernetes/{pod}/{table}", GetKubernetesPodIptablesOutput).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(defaultHandler)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	err = http.ListenAndServe(fmt.Sprintf(":%d", convertedPort), handlers.CORS(originsOk, headersOk, methodsOk)(r))
	if err != nil {
		fmt.Println("error in starting the server: ", err)
		return
	}
}
