package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"flag"
)

func clientset() *kubernetes.Clientset {
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
	return clientset

}

func homeDir() string {
	return os.Getenv("HOME")
}

func main() {


	clientss := clientset()
	svc, err := clientss.CoreV1().Services("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _,v := range svc.Items {
		fmt.Println(v.Name)
	}

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "jixingxing",
			Labels:
		},
	}
}