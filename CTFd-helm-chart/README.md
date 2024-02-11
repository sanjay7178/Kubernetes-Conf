![image](https://github.com/sanjay7178/Kubernetes-Conf/assets/97831658/e767938f-8f99-42aa-ba4f-c27b3bf0f6cd)### High Available CTFd based on premise/cloud kubernetes cluster Infrastructer

### Final Agenda
- Need to make helm chart or custom kubernetes object configuration (.yaml) files for custom ctf cluster

### Tips
- For fast development use minikube or kind
- For Productions try out microk8s or any kubernetes from cloud providers like aks , eks
  
### TODO
- [ ] Try out deploying the ctfd helm chart given in this on minikube or  kind or microk8s
- [ ] try to increase the default values `replicaCount` to 2-4 , and check kubernetes with the existing loadbalancer working or not
- [ ] If the helm chart contains istio proxy go with it or Make a istio proxy  in same service of the ctfd helm chart
- [ ] Add certmanager if possible to kubernetes cluster
- [ ] add the loadbalancer IP (public/private) from istio
- [ ] Try if possible to add the ArgoCD 

### open source ctfd helm chart to try out
https://artifacthub.io/packages/helm/ctfd/ctfd

### Proposed Pipeline

[visit updated pipeline](https://app.eraser.io/workspace/sNFhpzqv1y84jPh0JfIC?origin=share)

![image](https://github.com/sanjay7178/Kubernetes-Conf/assets/97831658/4efa5068-abfc-4916-86ed-8439b6758d53)

### Reference Screenshots
![image](https://github.com/sanjay7178/Kubernetes-Conf/assets/97831658/2fa18821-6cac-42a7-9a90-031bdc28ef56)

