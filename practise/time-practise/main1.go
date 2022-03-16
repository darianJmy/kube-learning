package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path"
)

func main() {
	containerConfig, err := clientcmd.BuildConfigFromFlags("", path.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		panic(err)
	}
	containerClient, err := kubernetes.NewForConfig(containerConfig)
	if err != nil {
		panic(err)
	}

	stopper := make(chan struct{})
	defer close(stopper)

	factory := informers.NewSharedInformerFactory(containerClient, 10)
	containerInformer := factory.Core().V1().Pods()
	informer := containerInformer.Informer()
	defer runtime.HandleCrash()

	go factory.Start(stopper)

	if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Time out waiting for caches to sync"))
		return
	}

	containerLister := containerInformer.Lister()

	containerlist, err := containerLister.Pods("default").List(labels.Everything())
	if err != nil {
		panic(err)
	}
	fmt.Println(containerlist)
	for _, containers := range containerlist {
		fmt.Printf("%s\n", containers.Status)

	}

}
