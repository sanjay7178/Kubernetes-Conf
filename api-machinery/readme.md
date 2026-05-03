API Machinery

-----------------

## kind

Pod 
Deployment

Deployment
apps/v1

GroupVersionKind

Status

HTTP Endpoint

## resource

deployments
apps/v1 


apis/apps/v1/namespaces/default/deployments?limit=500

replicasets
apis/apps/v1/namespaces/default/replicasets?limit=50

GroupVersionResource


GVK --> GVR conversion using RestMapping

## RestMapping 
https://pkg.go.dev/k8s.io/apimachinery/pkg/api/meta#RESTMapper
https://pkg.go.dev/k8s.io/apimachinery/pkg/api/meta#RESTMapping

## scheme 


---

Go Struct --> AddKnownTypes --> (using schemes) --> GVK ---> (using RestMapping) --> GVR --> Get HTTP Path
  |
  |
deepCopyObject setGVK