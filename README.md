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

## **Usage**

### Kubernetes

Upon the execution of the deployment manifest, check if all the pods in the namespace `iptables-viz` are in a `Running` status:

```bash
❯ kubectl get pods -n iptables-viz

NAME                                     READY   STATUS    RESTARTS        AGE
iptables-viz-backend-98bd5fcfb-mxwvh     1/1     Running   2 (2m33s ago)   22h
iptables-viz-frontend-7fff54cb4d-nj9zl   1/1     Running   7 (2m ago)      22h
```

After this, list the services in the `iptables-viz` namespace:

```bash
❯ kubectl get services -n iptables-viz

NAME                        TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
iptables-viz-backend-svc    ClusterIP   10.102.34.226   <none>        8080/TCP       22h
iptables-viz-frontend-svc   NodePort    10.109.72.229   <none>        80:30025/TCP   22h
```

To access the frontend web app in your browser, you can use the `iptables-viz-frontend-svc` services' NodePort (For example, `30025` in the above output) at the following URL:

```bash
http://<node-external-ip>:<node-port>
```

### Linux

Upon the execution of the installation script, check the status of the `iptables-viz` service:

```bash
❯ systemctl status iptables-viz.service

● iptables-viz.service - Oneshot service for iptables-viz
     Loaded: loaded (/etc/systemd/system/iptables-viz.service; enabled; vendor preset: enabled)
     Active: active (exited) since Sun 2023-04-16 20:47:51 IST; 23s ago
    Process: 34848 ExecStart=/bin/true (code=exited, status=0/SUCCESS)
   Main PID: 34848 (code=exited, status=0/SUCCESS)
        CPU: 822us

Apr 16 20:47:51 ubuntu systemd[1]: Starting Oneshot service for iptables-viz...
Apr 16 20:47:51 ubuntu systemd[1]: Finished Oneshot service for iptables-viz.
```

Upon ensuring its status is `active`, you can access the frontend web app in your browser at the following URL:

```bash
http://localhost:3000
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
