package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	httpClient, err := rest.HTTPClientFor(config)
	httpClient.Transport = http.DefaultTransport
	
	clientset, err := kubernetes.NewForConfigAndClient(config, httpClient)
	if err != nil {
		panic(err)
	}

	w, err := clientset.AppsV1().Deployments("default").Watch(context.Background(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for event := range w.ResultChan() {
		fmt.Println(event.Type, event.Object)
	}
}
