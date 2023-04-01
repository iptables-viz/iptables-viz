package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"text/template"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func ClientSetup() *kubernetes.Clientset {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		fmt.Printf("Error in new client config: %s\n", err)
	}
	clientset := kubernetes.NewForConfigOrDie(config)
	return clientset
}

func RunPodShellCommand(podName, tableName string) (string, error) {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("kubectl exec -n kube-system %s -- sh -c \"iptables -L -t %s\" | jc --iptables", podName, tableName))
	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("error in listing Iptables chains: %v\n", err)
		return "", err
	}

	return string(output), nil
}

func GetPodList(clientSet *kubernetes.Clientset) ([]string, error) {
	var podList []string
	pods, err := clientSet.CoreV1().Pods("kube-system").List(context.Background(), metav1.ListOptions{LabelSelector: "k8s-app=kube-proxy"})
	if err != nil {
		fmt.Printf("Error getting kube-proxy pod: %v\n", err)
		return nil, err
	}
	if len(pods.Items) == 0 {
		fmt.Println("kube-proxy replicas not found")
		return nil, fmt.Errorf("kube-proxy replicas not found")
	}
	for _, p := range pods.Items {
		podList = append(podList, p.Name)
	}
	return podList, nil
}

func htmlTemplate(pd PageData) (string, error) {
	html := `<HTML>
	<head><title>{{.Title}}</title></head>
	<body>
	{{.Body}}
	</body>
	</HTML>`

	// Parse the template
	tmpl, err := template.New("index").Parse(html)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer

	if err := tmpl.Execute(&out, pd); err != nil {
		return "", err
	}

	return out.String(), nil
}
