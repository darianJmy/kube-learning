package main

import "fmt"

//
//import "fmt"
//
//type factory struct {
//	namespace string
//}
//
//type SharedInformerOption func(*factory) *factory
//
//func main() {
//	a := Options(WithNamespace("namespace"), WithTweakListOptions("tweakListOptions"))
//	fmt.Printf(a.namespace)
//
//	b := WithNamespace("default")
//
//}
//
//func Options(options ...SharedInformerOption) *factory {
//	factory := &factory{
//		namespace: "default",
//	}
//
//	// Apply all options
//	for _, opt := range options {
//		factory = opt(factory)
//	}
//
//	return factory
//
//}
//
//func WithNamespace(namespace string) SharedInformerOption {
//	return func(factory *factory) *factory {
//		factory.namespace = namespace
//		return factory
//	}
//}
//
//func WithTweakListOptions(namespace string) SharedInformerOption {
//	return func(factory *factory) *factory {
//		factory.namespace = namespace
//		return factory
//	}
//}

func main() {
	options(WithNamespace("default"), WithTweak("default1"))
}

func options(options ...string) {
	for k, v := range options {
		fmt.Printf("%d, %v\n", k, v)
	}

}
func WithNamespace(namespace string) string {
	return namespace
}

func WithTweak(namespace string) string {
	return namespace
}
