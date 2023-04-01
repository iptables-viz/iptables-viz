package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
)


func GetDockerIptablesOutput(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	containerName := vars["container"]
	tableName := vars["table"]
	var resp Response
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
	var resp Response
	output, err := RunPodShellCommand(kubeProxyPodName, tableName)
	if err != nil {
		fmt.Println("error in running the shell command: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	// cmd := exec.Command("bash", "-c", fmt.Sprintf("kubectl exec -n kube-system %s -- sh -c \"iptables -L -t %s\" | jc --iptables", kubeProxyPodName, tableName))
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Printf("error in listing Iptables chains: %v\n", err)
	// 	return
	// }

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
    var resp KubernetesDefaultResponse
    clientSet := ClientSetup()
    podList, err := GetPodList(clientSet)
	if err != nil {
		fmt.Println("error in getting pod list: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// kubeProxyPodName := pods.Items[0].Name
	kubeProxyPodName := podList[0]
    tableName := "nat"
	output, err := RunPodShellCommand(kubeProxyPodName, tableName)
	if err != nil {
		fmt.Println("error in running the shell command: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
    // cmd := exec.Command("bash", "-c", fmt.Sprintf("kubectl exec -n kube-system %s -- sh -c \"iptables -L -t %s\" | jc --iptables", kubeProxyPodName, tableName))
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Printf("error in listing Iptables chains: %v\n", err)
	// 	return
	// }
    resp.PodName = kubeProxyPodName
    resp.IptableOutput = output
    w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		fmt.Println("error in encoding output: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetAllKubeProxyPods(w http.ResponseWriter, r *http.Request) {
	var resp KubeProxy
	// var podList []string
	clientSet := ClientSetup()
    // pods, err := clientset.CoreV1().Pods("kube-system").List(context.Background(), metav1.ListOptions{LabelSelector: "k8s-app=kube-proxy"})
	// if err != nil {
    //     fmt.Printf("Error getting kube-proxy pod: %v\n", err)
	//  	return
	// }
	// if len(pods.Items) == 0 {
	//  	fmt.Println("kube-proxy replica not found.")
	//  	return
	// }
	// for _, p := range pods.Items {
	// 	podList = append(podList, p.Name)
	// }
	podList, err := GetPodList(clientSet)
	if err != nil {
		fmt.Println("error in getting pod list: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp.PodNames = podList
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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/iptables/health", healthCheck)
	r.HandleFunc("/iptables/kubernetes", GetKubernetesDefault).Methods("GET")
	r.HandleFunc("/iptables/kubernetes/kubeProxyPods", GetAllKubeProxyPods).Methods("GET")
	r.HandleFunc("/iptables/kubernetes/{pod}/{table}", GetKubernetesPodIptablesOutput).Methods("GET")
	// r.HandleFunc("/iptables/docker/{container}/{table}", GetDockerIptablesOutput).Methods("GET")

	err := http.ListenAndServe(":5000", r)
	if err != nil {
		fmt.Println("error in starting the server: ", err)
		return
	}
}