# iptables-viz
A simple and scalable iptables visualisation tool which can integrate across Kubernetes and Linux.

## **Platforms Support**
- Kubernetes
- Linux

## **Pre-requisites**
- `kubectl` (for Kubernetes platform)
- `iptables` (for Linux platform)
- `jc` (for Linux platform)

## **Installation**

#### Kubernetes

The `kubernetes-deploy.yaml` creates the following:
- namespace `iptables-viz`
- role bindings for both backend and frontend for providing the namespace with appropriate functions to access the `kube-system` `kube-proxy` pods
- backend service of type `ClusterIP` and a new frontend service of type `NodePort`
- deployments for both backend and frontend

```bash
kubectl apply -f https://raw.githubusercontent.com/iptables-viz/iptables-viz/main/manifests/kubernetes-deploy.yaml
```

#### Linux

The `install.sh` creates the following:
- downloads the respective binaries for both backend and frontend 
- copies the binaries to their executable paths in the users' system
- creates the systemd unit files for both frontend and backend
- executes the unit files as linux service


```bash
bash https://raw.githubusercontent.com/iptables-viz/iptables-viz/main/scripts/install.sh
```

## **Uninstallation**

#### Kubernetes

```bash
kubectl delete -f https://raw.githubusercontent.com/iptables-viz/iptables-viz/main/manifests/kubernetes-deploy.yaml
```

#### Linux

```bash
bash https://raw.githubusercontent.com/iptables-viz/iptables-viz/main/scripts/uninstall.sh
````
