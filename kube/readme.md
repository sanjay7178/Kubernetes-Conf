kind-cluster-config.yml

```
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4

nodes:
- role: control-plane
  image: kindest/node:v1.32.2
- role: worker
  image: kindest/node:v1.32.2
- role: worker
  image: kindest/node:v1.32.2
  extraPortMappings:
  - containerPort : 80 
    hostPort : 80
    protocol : TCP 
  - containerPort : 443
#- role: worker
#  image: kindest/node:v1.31.2
networking:
  apiServerAddress: ""   # Bind to all interfaces (including the public IP)
  apiServerPort: 45803           # External port for API server
```

create a kind cluster

```
kind create cluster --config  kind-cluster-config.yml --name kubebrowse-cluster 

```
get kubeconfig from kind 

```
kind get kubeconfig --name=kind-kubebrowse-cluster
```
