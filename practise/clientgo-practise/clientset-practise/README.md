## ClientSet 使用记录
### 示例
通过 client-go 提供的 Clientset 对象获取资源数据主要有以下三个步骤：
- 使用 kubeconfig 文件或者 ServiceAccount 提供的 secret 文件创建访问 Kubernetes API 的 Restful 参数，也就是代码中的 `rest.Config` 对象（ kubeconfig 是 pod 外访问，ServiceAccount 是 pod 内访问）
- 使用 `rest.config` 参数创建 Clientset 对象，这一步直接调用 `NewForConfig` 函数就行
- Clientset 对象的方法可以获取各个 Group 下面的对应资源对象进行 CRUD 操作

### 代码分析
对示例中使用到的部分 clientset 代码进行分析
```
// 对于代码名称就能发现此处代码是用于在集群里面
func InClusterConfig() (*Config, error) {
    // 常量部分值就是 ServiceAccount 产生的 token 与 ca，在集群内部无法获取 kubeconfig 文件
    // 所以就通过 ServiceAccount 生成的权限对集群进行操作
    // 通过常量得出，serviceaccount 生成的文件目录是固定的
	const (
		tokenFile  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
		rootCAFile = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
	)
	...
    // 主要看返回值，此处就是返回了一个可以访问 NewForConfig 函数的 config 结构体
	return &Config{
		// TODO: switch to using cluster DNS.
		Host:            "https://" + net.JoinHostPort(host, port),
		TLSClientConfig: tlsClientConfig,
		BearerToken:     string(token),
		BearerTokenFile: tokenFile,
	}, nil
}

// $HOME/.kube/config 文件是默认的目录文件，但是并不是固定的，所以集群外访问函数对于 config 文件目录做了一个参数，可以人为改变
func BuildConfigFromFlags(masterUrl, kubeconfigPath string) (*restclient.Config, error) {
	if kubeconfigPath == "" && masterUrl == "" {
		klog.Warning("Neither --kubeconfig nor --master was specified.  Using the inClusterConfig.  This might not work.")
		...
	}
	// 最终是通过 kubeconfig 文件内容构造一个 config 结构体
	return NewNonInteractiveDeferredLoadingClientConfig(
		&ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: masterUrl}}).ClientConfig()
}

// 可以通过 config 结构体访问 NewForConfig函数，最终返回了一个 clinetset 对象
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	// share the transport between all clients
	// 构造一个 httpclient 对象
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}
    // 通过把 config 对象与 httpclient 对象传给 NewForConfigAndClient，完善 config 对象，并且获得 clientset 对象以及其方法
	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	// 完善 config 对象
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	将其他 Group 和版本的资源的 RESTClient 封装到全局的 Clientset 对象中
	cs.admissionregistrationV1, err = admissionregistrationv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	...
	}
	cs.appsV1beta1, err = appsv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	...
}
```