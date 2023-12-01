## **iptables-viz Helm Chart**

This Helm chart simplifies the deployment of the iptables-viz application in a Kubernetes environment.

#Prerequisites
Kubernetes cluster (e.g., Minikube)

Helm installed and initialized
#Download Helm:

Visit the Helm GitHub releases page: ``` https://github.com/helm/helm/releases ```

Download the appropriate version of Helm for your operating system. For example, for Linux:
```bash
wget https://get.helm.sh/helm-v3.x.x-linux-amd64.tar.gz
tar -zxvf helm-v3.x.x-linux-amd64.tar.gz
```
#Move Helm binary to a directory in your PATH:
After extracting the Helm archive, move the helm binary to a directory that is included in your system's PATH. For example:
```bash
sudo mv linux-amd64/helm /usr/local/bin/helm
```
#Verify Installation:
Check if Helm is installed correctly by running:
```bash
helm version
```
#Installation
1. Clone the Repository
```bash
git clone <repository_url>
```
```bash
cd iptables-viz/iptables-viz
```
In this thier are two folder frontend backend

2. Start Minikube

#Ensure Minikube is started with the Docker driver:
```bash
minikube start --force --driver=docker
```

3. Install iptables-viz Backend

go to
```bash
cd iptables-viz/iptables-viz/backend
```
Run
```bash
helm install backend ./ -n iptables-viz --create-namespace
```
4. Install iptables-viz Frontend
go to
```bash
cd iptables-viz/iptables-viz/frontend
```
Run
```bash
helm install frontend ./ -n iptables-viz
```
#Upgrade Existing Release

1. Upgrade Backend Release
go to
```bash
cd iptables-viz/iptables-viz/backend
```
Run
```bash
helm upgrade backend ./ -n iptables-viz
```
3. Upgrade Frontend Release
go to
```bash
cd iptables-viz/iptables-viz/frontend
```
Run
```bash
helm upgrade frontend ./ -n iptables-viz
```
#Verify Deployment
```bash
kubectl get pods -n iptables-viz
```
```bash
kubectl get services -n iptables-viz
```
#Accessing iptables-viz
Once the deployment is successful, you can access the iptables-viz frontend using the Minikube IP:
```bash
minikube ip
```
Open a web browser and navigate to ``` http://<minikube_ip>:30025``` or your localhost ip to view the iptables-viz application.
