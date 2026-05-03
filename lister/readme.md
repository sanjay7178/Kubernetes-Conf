
1. client-go https://github.com/kubernetes/client-go
2. api https://github.com/kubernetes/api
3. apimachinery https://github.com/kubernetes/apimachinery

  go struct 
apimachinery pkg --> runtime.Object interface


pod 
  typemeta


pod (DeepCopyObject)

k8s object
  typeMeta
    kind
      apiVersion 
  objectMeta
  spec (replicas)
  status (controller objeccts)
  
typeMeta mapped to --> 
apiVersion : apps/v1
kind: Deployment

objectMeta mapped to -->




resourceVersion --> how many times a resource get edited (which makes cache resync)