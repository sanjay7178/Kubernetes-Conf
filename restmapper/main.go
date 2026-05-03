package main

import (
	// "k8s.io/apimachinery/pkg/api/meta"
	"flag"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func main() {
	res := "pods"
	flag.StringVar(&res , "res", "", "resource that you want to interact with")
	flag.Parse()
	configFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	matchVersionFlags  := cmdutil.NewMatchVersionFlags(configFlags)
	m , err := matchVersionFlags.ToRESTMapper()
	if err != nil {
		fmt.Printf("creating rest mapper: %v\n", err.Error())
		return 
	}
	gvr , err := m.ResourceFor(schema.GroupVersionResource{
		Resource : res,
	})
	if err != nil {
		fmt.Printf("getting gvr for resource: %v\n", err.Error())
		return 
	}
	fmt.Printf("Complete gvr is group : %s , version : %s  , resoruce : %s ", gvr.Group , gvr.Version, gvr.Resource)
	
	// restConfig, err := configFlags.ToRESTConfig()
	if err != nil {
		
	}
	
}
