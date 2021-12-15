# add-operator
Basic Kubernetes operator that have multiple versions in CRD. This operator can be used to experiment and understand Operator/CRD behaviors.

# Dependencies
- operator-sdk https://sdk.operatorframework.io/docs/installation/

# Build and Deploy
Add different kubebuilder markers in api/<version>/add_types.go depending on the CRD you want to create. Some of the useful markers are
  - //+kubebuilder:storageversion
  - //+kubebuilder:unservedversion
  - //+kubebuilder:skipversion

`make manifests` - To generate the CRD config/crd/bases/math.example.com_adds.yaml
  
`make dockber-build` - To create the operator image
  
 set KUBECONFIG
  
`make deploy` - To deploy the CRD, operator, etc to the cluster
  
  
# Add new version
  - operator-sdk create api --group math --version v3 --kind Add --resource
