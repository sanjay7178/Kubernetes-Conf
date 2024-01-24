### Cheat Sheet

### Introduction to Helm 
Helm is a package manager for Kubernetes that simplifies the deployment and management of applications. Below are some exemplary Helm commands and their explanations:

1. **Initialize Helm:**
   ```bash
   helm init
   ```
   Initializes Helm on your local machine and installs Tiller (Helm's server component) into your Kubernetes cluster.

2. **Add Helm repository:**
   ```bash
   helm repo add stable https://charts.helm.sh/stable
   ```
   Adds the Helm stable repository, which contains a collection of pre-packaged charts.

3. **Search for Helm charts:**
   ```bash
   helm search repo stable
   ```
   Searches for charts in the stable repository.

4. **Install a Helm chart:**
   ```bash
   helm install my-release stable/chart-name
   ```
   Installs a Helm chart with a given release name. Replace `chart-name` with the name of the chart you want to install.

5. **List installed releases:**
   ```bash
   helm list
   ```
   Lists all installed releases along with their status.

6. **Upgrade a Helm release:**
   ```bash
   helm upgrade my-release stable/chart-name
   ```
   Upgrades an existing Helm release to a new version of the chart.

7. **Rollback a Helm release:**
   ```bash
   helm rollback my-release 1
   ```
   Rolls back a release to a specific revision number.

8. **Delete a Helm release:**
   ```bash
   helm delete my-release
   ```
   Deletes a Helm release and associated resources.

9. **Inspect a Helm release:**
   ```bash
   helm get values my-release
   ```
   Shows the values used in a release.

10. **Dry run of Helm install:**
    ```bash
    helm install --dry-run --debug my-release stable/chart-name
    ```
    Performs a dry run of the installation without actually deploying resources, and shows debug information.

11. **Package a Helm chart:**
    ```bash
    helm package my-chart-directory
    ```
    Packages a Helm chart directory into a versioned archive file.

12. **Lint a Helm chart:**
    ```bash
    helm lint my-chart-directory
    ```
    Checks a Helm chart for potential issues.

13. **Create a Helm chart:**
    ```bash
    helm create my-chart
    ```
    Creates a new Helm chart template with default files and directories.

14. **Fetch the latest Helm charts:**
    ```bash
    helm repo update
    ```
    Updates the local cache of Helm charts from their respective repositories.
Certainly! Here are some more Helm commands and use cases:

15. **Show Helm release history:**
    ```bash
    helm history my-release
    ```
    Displays the release history, including revision numbers and deployment times.

16. **Install Helm chart from a local directory:**
    ```bash
    helm install my-release ./my-local-chart
    ```
    Installs a Helm chart from a local directory.

17. **Override Helm chart values during installation:**
    ```bash
    helm install my-release stable/chart-name --set key1=value1,key2=value2
    ```
    Overrides Helm chart values during installation.

18. **Install Helm chart using values from a YAML file:**
    ```bash
    helm install my-release stable/chart-name -f values.yaml
    ```
    Specifies custom values stored in a YAML file during installation.

19. **List available Helm repositories:**
    ```bash
    helm repo list
    ```
    Lists all Helm repositories configured on your local machine.

20. **Remove Helm repository:**
    ```bash
    helm repo remove repository-name
    ```
    Removes a Helm repository from your local machine.

21. **Create a Helm chart dependency file:**
    ```bash
    helm dependency update my-chart
    ```
    Generates or updates the requirements.lock file based on the contents of the requirements.yaml file.

22. **Package and install a Helm chart from a URL:**
    ```bash
    helm install my-release https://example.com/path/to/chart-0.1.0.tgz
    ```
    Installs a Helm chart directly from a URL.

23. **Inspect Helm chart templates:**
    ```bash
    helm template my-release stable/chart-name
    ```
    Renders the Kubernetes YAML manifests for a Helm chart without installing it.

24. **Debug Helm releases:**
    ```bash
    helm debug my-release
    ```
    Enters a debugging shell inside a running container of a Helm release.

25. **Fetch Helm release's status:**
    ```bash
    helm status my-release
    ```
    Shows the status of a Helm release, including deployed resources and their statuses.

26. **Create a Helm chart package and upload to a chart repository:**
    ```bash
    helm package my-chart
    helm push my-chart-version.tgz my-chart-repo
    ```
    Packages a Helm chart and uploads it to a Helm chart repository.

These commands cover a broad range of Helm use cases, from basic operations to more advanced scenarios. Feel free to explore Helm's documentation for in-depth details and additional features: [Helm Documentation](https://helm.sh/docs/).
