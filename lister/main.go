package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	// "go/format"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	// "k8s.io/apimachinery/pkg/runtime"
	// "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	// "github.com/k0kubun/pp/v3"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/sanjay7178/.kube/config", "location of the kubeconfig file")
	flag.Parse()
	config , err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
		config , err = rest.InClusterConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

    config.Timeout = 120*time.Second
	clientset , err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	informerfactory := informers.NewSharedInformerFactory(clientset, 30*time.Second)


	podInformer := informerfactory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				fmt.Println("AddFunc: ", obj)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				fmt.Println("UpdateFunc: ", oldObj, newObj)
			},
			DeleteFunc: func(obj interface{}) {
				fmt.Println("DeleteFunc: ", obj)
			},
		},
	)

	informerfactory.Start(wait.NeverStop)
	informerfactory.WaitForCacheSync(wait.NeverStop)
	pod  , err := podInformer.Lister().Pods("default").Get("default")
	fmt.Println(pod)
	
	
	// ctx := context.Background()
	// pods, err := clientset.CoreV1().Pods("default").List( ctx , metav1.ListOptions{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// // pp.Print(pods)
	// // fmt.Println(pods)
	// for _, pod := range pods.Items {
	// 	fmt.Println(pod.Name)
	// }	
	// fmt.Println("Deployments:")
	// deployments, err := clientset.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// for _, deployment := range deployments.Items {
	// 	fmt.Println(deployment.Name)
	// }
}
