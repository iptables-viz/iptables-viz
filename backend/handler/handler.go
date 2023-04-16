package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"log"

	"github.com/gorilla/mux"
	"github.com/iptables-viz/iptables-viz/backend/models"
	"github.com/iptables-viz/iptables-viz/backend/utility"
)

// handler to fetch the iptables output from the kube-proxy pod
func GetKubernetesPodIptablesOutput(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	kubeProxyPodName := vars["pod"]
	tableName := vars["table"]
	var resp models.Response
	output, err := utility.RunPodShellCommand(kubeProxyPodName, tableName)
	if err != nil {
		log.Printf("Failed to run the shell command, %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Shell command error: %v", err)
	}
	resp.Entity = "kubernetes"
	resp.IptableOutput = output
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Printf("Failed to JSON encode the response, %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "JSON Encode Error: %v", err)
		return
	}
}

// handler to fetch the list of kube-proxy pods and get default iptables nat table rules
func GetKubernetesDefault(w http.ResponseWriter, r *http.Request) {
	var resp models.KubernetesDefaultResponse
	clientSet := utility.ClientSetup()
	podList, err := utility.GetPodList(clientSet)
	if err != nil {
		log.Printf("Error in getting pod list, %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Pod List error: %v", err)
		return
	}
	if len(podList) == 0 {
		log.Printf("No kube-proxy pods found")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "No kube-proxy pods found")
		return
	}
	kubeProxyPodName := podList[0]
	tableName := "nat"
	output, err := utility.RunPodShellCommand(kubeProxyPodName, tableName)
	if err != nil {
		log.Printf("Error in running the shell command, %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Shell command error: %v", err)
		return
	}
	resp.CurrentPodName = kubeProxyPodName
	resp.IptableOutput = output
	resp.PodList = podList
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Printf("Failed to JSON encode the response, %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "JSON Encode Error: %v", err)
		return
	}
}

// handler for validating whether API is running or not
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	health := make(map[string]string)
	health["now"] = now.Format(time.ANSIC)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(health); err != nil {
		log.Printf("Failed to JSON encode the response, %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "JSON Encode Error: %v", err)
		return
	}
}

// default handler for invalid routes
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Invalid Access: requested route not found")
}

// handler to fetch the iptables output for a given table for linux
func GetLinuxIptableOutput(w http.ResponseWriter, r *http.Request) {
	var resp models.LinuxIptableOutput
	vars := mux.Vars(r)
	tableName := vars["table"]
	cmd := exec.Command("bash", "-c", fmt.Sprintf("iptables -w -L -t %s | jc --iptables --quiet", tableName))
	out, err := cmd.CombinedOutput()
	output := strings.TrimSpace(string(out))
	if err != nil {
		if output != "" {
			log.Printf("Error in running the shell command, %s, %s", err.Error(), output)
			fmt.Fprintf(w, "Shell command error, %s, %s", err.Error(), output)
		} else {
			log.Printf("Error in running the shell command, %s", err.Error())
			fmt.Fprintf(w, "Shell command error, %s", err.Error())
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp.IptableOutput = string(output)
	resp.TableName = tableName
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Printf("Failed to JSON encode the response, %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "JSON Encode Error: %v", err)
		return
	}
}
