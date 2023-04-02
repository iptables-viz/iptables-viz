package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/iptables-viz/iptables-viz/backend/models"
	"github.com/iptables-viz/iptables-viz/backend/utility"
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
		fmt.Fprint(w, fmt.Sprintf("JSON Encode Error: %v", err))
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
		fmt.Fprint(w, fmt.Sprintf("Shell command error: %v", err))
	}
	resp.Entity = "kubernetes"
	resp.IptableOutput = output
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		fmt.Println("error in encoding output: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fmt.Sprintf("JSON Encode Error: %v", err))
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
		fmt.Fprint(w, fmt.Sprintf("Pod List error: %v", err))
		return
	}
	kubeProxyPodName := podList[0]
	tableName := "nat"
	output, err := utility.RunPodShellCommand(kubeProxyPodName, tableName)
	if err != nil {
		fmt.Println("error in running the shell command: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fmt.Sprintf("Shell command error: %v", err))
	}
	resp.CurrentPodName = kubeProxyPodName
	resp.IptableOutput = output
	resp.PodList = podList
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		fmt.Println("error in encoding output: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fmt.Sprintf("JSON Encode Error: %v", err))
		return
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	health := make(map[string]string)
	health["now"] = now.Format(time.ANSIC)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Invalid Access: requested route not found")
}

func GetLinuxIptableOutput(w http.ResponseWriter, r *http.Request) {
	var resp models.LinuxIptableOutput
	vars := mux.Vars(r)
	tableName := vars["table"]
	cmd := exec.Command("bash", "-c", fmt.Sprintf("iptables -L -t %s | jc --iptables", tableName))
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("error in running the shell command: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, fmt.Sprintf("Shell command error: %v", err))
		return
	}
	resp.IptableOutput = string(output)
	resp.TableName = tableName
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		fmt.Println("error in encoding output: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, fmt.Sprintf("JSON Encode Error: %v", err))
		return
	}
}