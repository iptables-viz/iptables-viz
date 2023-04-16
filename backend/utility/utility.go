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
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/klog"
)

var KubeConfigFilePath string

// checks if kubeConfigPath is given or not, if not uses the inCluster config
func buildConfigFromFlags(masterURL, kubeConfigPath string) (*rest.Config, error) {
	if kubeConfigPath == "" && masterURL == "" {
		kubeconfig, err := rest.InClusterConfig()
		if err == nil {
			return kubeconfig, nil
		}
		klog.Warningf("Neither --kubeconfig nor --master was specified. Using the inClusterConfig. Error creating inClusterConfig: %v", err)
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath},
		&clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: masterURL}}).ClientConfig()
}


// fetches the kubeconfig
func getKubeConfig() (*rest.Config, error) {
	config, err := buildConfigFromFlags("", KubeConfigFilePath)
	return config, err
}

// create a new k8s REST client
func ClientSetup() *kubernetes.Clientset {
	// Create a Kubernetes REST config
	config, err := getKubeConfig()
	if err != nil {
		log.Fatalf("Error while getting config, %s", err.Error())
	}

	// Create a Kubernetes client from the REST config
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error while creating the kubernetes client, %s", err.Error())
	}

	return client
}

// runs the siptables command inside the kube-proxy pod
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

// fetches the list of available kube-proxy pods
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
