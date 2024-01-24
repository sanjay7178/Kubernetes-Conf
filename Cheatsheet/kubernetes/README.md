Certainly! Here's a cheat sheet for common Kubernetes commands:

### Kubernetes Basics:

1. **Cluster Info:**
   ```bash
   kubectl cluster-info
   ```
   Displays the cluster endpoint information.

2. **Get Nodes:**
   ```bash
   kubectl get nodes
   ```
   Shows the status of nodes in the cluster.

3. **Get Pods:**
   ```bash
   kubectl get pods
   ```
   Lists all pods in the default namespace.

4. **Get Services:**
   ```bash
   kubectl get services
   ```
   Lists all services in the default namespace.

### Deployments and Replicas:

5. **Create Deployment:**
   ```bash
   kubectl create deployment my-deployment --image=image-name
   ```
   Creates a deployment with the specified image.

6. **Scale Deployment:**
   ```bash
   kubectl scale deployment my-deployment --replicas=3
   ```
   Scales the deployment to the desired number of replicas.

7. **Rollout Status:**
   ```bash
   kubectl rollout status deployment/my-deployment
   ```
   Shows the status of a deployment rollout.

8. **Rollback Deployment:**
   ```bash
   kubectl rollout undo deployment/my-deployment
   ```
   Rolls back a deployment to the previous revision.

### Pods and Containers:

9. **Describe Pod:**
   ```bash
   kubectl describe pod pod-name
   ```
   Provides detailed information about a pod.

10. **Logs from Pod:**
    ```bash
    kubectl logs pod-name
    ```
    Displays the logs of a pod.

11. **Execute Command in Pod:**
    ```bash
    kubectl exec -it pod-name -- /bin/bash
    ```
    Opens a shell in a running pod.

12. **Port Forwarding:**
    ```bash
    kubectl port-forward pod-name 8080:80
    ```
    Forwards local port 8080 to port 80 in the pod.

### Services:

13. **Expose Deployment:**
    ```bash
    kubectl expose deployment my-deployment --type=LoadBalancer --port=80
    ```
    Creates a service to expose the deployment.

14. **Get Service Details:**
    ```bash
    kubectl get svc my-service
    ```
    Displays details about a service.

15. **Delete Service:**
    ```bash
    kubectl delete service my-service
    ```
    Deletes a service.

### ConfigMaps and Secrets:

16. **Create ConfigMap:**
    ```bash
    kubectl create configmap my-config --from-file=config-files/
    ```
    Creates a ConfigMap from files in a directory.

17. **Create Secret:**
    ```bash
    kubectl create secret generic my-secret --from-literal=key1=value1 --from-file=secrets/
    ```
    Creates a secret from literal values and files.

### Namespaces:

18. **Create Namespace:**
    ```bash
    kubectl create namespace my-namespace
    ```
    Creates a new namespace.

19. **Switch Namespace:**
    ```bash
    kubectl config set-context --current --namespace=my-namespace
    ```
    Switches the current context to a different namespace.

### Miscellaneous:

20. **Apply Configuration:**
    ```bash
    kubectl apply -f filename.yaml
    ```
    Applies a configuration from a YAML file.

21. **Delete Resource:**
    ```bash
    kubectl delete deployment my-deployment
    ```
    Deletes a deployment.

22. **Get Detailed Resource Info:**
    ```bash
    kubectl get pod pod-name -o yaml
    ```
    Retrieves detailed information about a resource in YAML format.

23. **Kubernetes Dashboard:**
    ```bash
    kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0/aio/deploy/recommended.yaml
    ```
    Installs the Kubernetes Dashboard. (Note: For production, consider securing the dashboard.)

Remember to replace placeholders like `my-deployment`, `pod-name`, etc., with your actual resource names. This cheat sheet covers basic Kubernetes operations, and there are many more advanced features and concepts to explore in the Kubernetes ecosystem.
