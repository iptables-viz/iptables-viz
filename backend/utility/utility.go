package utility

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func ClientSetup() *kubernetes.Clientset {
	// Create a Kubernetes REST config
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Error while getting in-cluster config, %s", err.Error())
	}

	// Create a Kubernetes client from the REST config
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error while creating the kubernetes client, %s", err.Error())
	}

	return client
}

func RunPodShellCommand(podName, tableName string) (string, error) {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("kubectl exec -n kube-system %s -- sh -c \"iptables -w -L -t %s\" | jc --iptables --quiet", podName, tableName))
	out, err := cmd.CombinedOutput()
	output := strings.TrimSpace(string(out))
	if err != nil {
		if output != "" {
			return "", errors.Errorf("failed to list iptables rules, %s, %s", err.Error(), output)
		} else {
			return "", errors.Errorf("failed to list iptables rules, %s", err.Error())
		}
	}
	return string(output), nil
}

func GetPodList(clientSet *kubernetes.Clientset) ([]string, error) {
	var podList []string
	pods, err := clientSet.CoreV1().Pods("kube-system").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, errors.Errorf("failed to get kube-proxy pods, %s", err.Error())
	}
	if len(pods.Items) == 0 {
		return nil, fmt.Errorf("kube-proxy pods not found")
	}
	for _, p := range pods.Items {
		if strings.Contains(p.Name, "kube-proxy") {
			podList = append(podList, p.Name)
		}
	}
	return podList, nil
}
