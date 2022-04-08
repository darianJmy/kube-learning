package main

import (
	"flag"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path"
	"time"
)

func main() {
	var kubeconfig *string
	var config *restclient.Config
	var clientset *kubernetes.Clientset
	var err error
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", path.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	if config, err = restclient.InClusterConfig(); err != nil {
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
			panic(err.Error())
		}
	}

	if clientset, err = kubernetes.NewForConfig(config); err != nil {
		panic(err.Error())
	}

	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)

	deployInformer := informerFactory.Apps().V1().Deployments()

	informer := deployInformer.Informer()

	lister := deployInformer.Lister()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    OnAdd,
		UpdateFunc: nil,
		DeleteFunc: nil,
	})

	stopCh := make(chan struct{})
	defer close(stopCh)

	informerFactory.Start(stopCh)

	informerFactory.WaitForCacheSync(stopCh)

	deployments, err := lister.Deployments("default").List(labels.Everything())
	if err != nil {
		panic(err.Error())
	}
	for idx, deploy := range deployments {
		fmt.Printf("%d -> %s\n", idx+1, deploy.Name)
	}

	<-stopCh

}

func OnAdd(obj interface{}) {
	deploy := obj.(*v1.Deployment)
	fmt.Println("add a deployment:", deploy.Name)
}
