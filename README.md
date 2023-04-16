# iptables-viz
![iptables-viz](https://user-images.githubusercontent.com/43271557/232320410-2f83ee5b-7765-4a15-9b98-90a637bd26d6.png)

A simple and scalable iptables visualization tool which can be used across Kubernetes and Linux.

## **Supported platforms**

- Kubernetes
- Linux

## **Pre-requisites**

### Kubernetes

- `kubectl` (https://kubernetes.io/docs/reference/kubectl/)

### Linux

- `jc` (https://github.com/kellyjonbrazil/jc)
- `serve` (https://www.npmjs.com/package/serve)

## **Installation**

### Kubernetes

The [kubernetes-deploy.yaml](manifests/kubernetes-deploy.yaml) manifest creates the following Kubernetes resources as part of the installation:

- Namespace `iptables-viz`.
- Deployments `iptables-viz-backend` and `iptables-viz-frontend` for backend and frontend respectively.
- RBAC for both backend and frontend Deployments for providing appropriate permissions to access `kube-proxy` pods in the `kube-system` namespace and accessing the backend service respectively.
- Services for both backend and frontend of type `ClusterIP` and `NodePort` respectively.

Execute the following command to deploy the application:

```bash
kubectl apply -f https://raw.githubusercontent.com/iptables-viz/iptables-viz/main/manifests/kubernetes-deploy.yaml
```

### Linux

The [install.sh](scripts/install.sh) script performs the following steps as part of the installation:

1. Downloads the appropriate backend server binary and the frontend web app.
2. Copies the downloaded artifacts to their executable paths in the users' system.
3. Creates Systemd unit files for both frontend and backend.
4. Executes the unit files as Linux service.

```bash
curl https://raw.githubusercontent.com/iptables-viz/iptables-viz/main/scripts/install.sh | sudo bash
```

## **Uninstallation**

### Kubernetes

```bash
kubectl delete -f https://raw.githubusercontent.com/iptables-viz/iptables-viz/main/manifests/kubernetes-deploy.yaml
```

### Linux

```bash
curl https://raw.githubusercontent.com/iptables-viz/iptables-viz/main/scripts/uninstall.sh | sudo bash
```
