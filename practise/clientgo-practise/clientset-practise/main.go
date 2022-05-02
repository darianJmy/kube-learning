package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/storageos/go-api/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"path/filepath"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := restclient.InClusterConfig()
	if err != nil {
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
			panic(err)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deployment, err := clientset.AppsV1().Deployments(types.DefaultNamespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for k, v := range deployment.Items {
		fmt.Printf("%d -> %s \n", k, v.Name)
	}
}

func homeDir() string {
	return os.Getenv("HOME")
}
