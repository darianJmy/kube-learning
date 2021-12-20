package main

import (
	"context"
	"fmt"
	"path"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/runtime"
	cacheddiscovery "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/scale"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/controller-manager/pkg/clientbuilder"
	"k8s.io/kubernetes/pkg/controller/podautoscaler/metrics"
	resourceclient "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1beta1"
	"k8s.io/metrics/pkg/client/custom_metrics"
	"k8s.io/metrics/pkg/client/external_metrics"
)

func main() {
	hpaConfig, err := clientcmd.BuildConfigFromFlags("", path.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		panic(err)
	}
	hpaClient, err := kubernetes.NewForConfig(hpaConfig)
	if err != nil {
		panic(err)
	}

	stopper := make(chan struct{})
	defer close(stopper)

	factory := informers.NewSharedInformerFactory(hpaClient, 10)
	hpaInformer := factory.Autoscaling().V2beta2().HorizontalPodAutoscalers()
	informer := hpaInformer.Informer()
	defer runtime.HandleCrash()

	go factory.Start(stopper)

	if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Time out waiting for caches to sync"))
		return
	}

	hpaLister := hpaInformer.Lister()

	hpaget, err := hpaLister.HorizontalPodAutoscalers(apiv1.NamespaceDefault).Get("podinfo")
	if err != nil {
		panic(err)
	}
	fmt.Println(hpaget)

	hpa := hpaget.DeepCopy()
	targetGV, err := schema.ParseGroupVersion(hpa.Spec.ScaleTargetRef.APIVersion)
	if err != nil {
		panic(err)
	}
	targetGK := schema.GroupKind{
		Group: targetGV.Group,
		Kind:  hpa.Spec.ScaleTargetRef.Kind,
	}
	fmt.Println(targetGK)

	rootClientBuilder := clientbuilder.SimpleControllerClientBuilder{
		ClientConfig: hpaConfig,
	}
	discoveryClient := rootClientBuilder.DiscoveryClientOrDie("horizontal-pod-autoscaler")
	cachedClient := cacheddiscovery.NewMemCacheClient(discoveryClient)
	restMapper := restmapper.NewDeferredDiscoveryRESTMapper(cachedClient)

	scaleKindResolver := scale.NewDiscoveryScaleKindResolver(hpaClient.Discovery())
	scaleClient, err := scale.NewForConfig(hpaConfig, restMapper, dynamic.LegacyAPIPathResolverFunc, scaleKindResolver)
	if err != nil {
		panic(err)
	}

	mappings, err := restMapper.RESTMappings(targetGK)
	if err != nil {
		panic(err)
	}
	targetGR := mappings[0].Resource.GroupResource()
	//获取 scale
	scale, err := scaleClient.Scales("default").Get(context.TODO(),targetGR,"podinfo", v1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(scale.Spec)
	fmt.Println(scale.Status)

	apiVersionsGetter := custom_metrics.NewAvailableAPIsGetter(hpaClient.Discovery())
	metricsClient := metrics.NewRESTMetricsClient(
		resourceclient.NewForConfigOrDie(hpaConfig),
		custom_metrics.NewForConfig(hpaConfig, restMapper, apiVersionsGetter),
		external_metrics.NewForConfigOrDie(hpaConfig),
	)

	selector, err := labels.Parse(scale.Status.Selector)
	if err != nil {
		panic(err)
	}
	metricSelector, err := metav1.LabelSelectorAsSelector(hpaget.Spec.Metrics[0].Pods.Metric.Selector)
	if err != nil {
		panic(err)
	}
	fmt.Println(selector)
	fmt.Println(metricSelector)

	metric, timestamp, err := metricsClient.GetRawMetric("http_requests","default",selector,metricSelector)
	if err != nil {
		panic(err)
	}
	fmt.Println(metric)
	fmt.Println(timestamp)
}
