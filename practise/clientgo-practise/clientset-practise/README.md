## ClientSet 使用记录
### 示例
通过 client-go 提供的 Clientset 对象获取资源数据主要有以下三个步骤：
- 使用 kubeconfig 文件或者 ServiceAccount 提供的 secret 文件创建访问 Kubernetes API 的 Restful 参数，也就是代码中的 `rest.Config` 对象（ kubeconfig 是 pod 外访问，ServiceAccount 是 pod 内访问）
- 使用 `rest.config` 参数创建 Clientset 对象，这一步直接调用 `NewForConfig` 函数就行
- Clientset 对象的方法可以获取各个 Group 下面的对应资源对象进行 CRUD 操作

### 代码分析
对示例中使用到的部分 clientset 代码进行分析
```
// 对于代码名称就能发现此处代码是通过集群内部文件获取 config 信息
func InClusterConfig() (*Config, error) {
    // 常量部分值就是 ServiceAccount 产生的 token 与 ca，在集群内部无法获取 kubeconfig 文件
    // 所以就通过 ServiceAccount 生成的权限对集群进行操作
    // 通过常量得出，serviceaccount 生成的文件目录是固定的
	const (
		tokenFile  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
		rootCAFile = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
	)
	...
    // 返回值是一个 config 结构体，这个结构体包含了调用 Clientset 的一些基础信息
	return &Config{
		// TODO: switch to using cluster DNS.
		Host:            "https://" + net.JoinHostPort(host, port),
		TLSClientConfig: tlsClientConfig,
		BearerToken:     string(token),
		BearerTokenFile: tokenFile,
	}, nil
}

// 对于代码名称就能发现此处代码是通过集群外部文件获取 config 信息
func BuildConfigFromFlags(masterUrl, kubeconfigPath string) (*restclient.Config, error) {
	if kubeconfigPath == "" && masterUrl == "" {
		klog.Warning("Neither --kubeconfig nor --master was specified.  Using the inClusterConfig.  This might not work.")
		...
	}
	// 最终是通过集群外部 kubeconfig 文件内容构造一个 config 结构体
	return NewNonInteractiveDeferredLoadingClientConfig(
		&ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: masterUrl}}).ClientConfig()
}

// 可以通过 config 结构体访问 NewForConfig 函数，获得 Clientset 对象
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	// share the transport between all clients
	// 构造一个 httpclient 对象
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}
    // config 对象与 httpclient 对象传给 NewForConfigAndClient 函数，获得 clientset 对象以及其方法
	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// c 代表着发送 http 服务的一些认证信息文件，httpClient 代表着 http 服务
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
    configShallowCopy := *c
    ...
	var cs Clientset
	var err error
	// 封装 Clientset 结构体
	cs.admissionregistrationV1, err = admissionregistrationv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.admissionregistrationV1beta1, err = admissionregistrationv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.internalV1alpha1, err = internalv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	// 用 appsV1 举例，Clientset的 appsV1 对象使用的是 appsv1 包的 NewForConfigAndClient 函数
	// 此处还把 config、httpClient 传给函数
	cs.appsV1, err = appsv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	...
}

// 通过 config、httpClient 完善 AppsV1Client 
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*AppsV1Client, error) {
	config := *c
	// 用的指针，返回的 err，说明函数内部完善 config 
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	// 通过函数生成 RESTClient
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	// 结构体有多个对象，说明 AppsV1Client 才完善了里面的 client 部分
	// 此处 AppsV1Client 对象已经构造出来
	// 具体为完善了 clientset.AppsV1
	return &AppsV1Client{client}, nil
}

// 此处调用了 clientset.AppsV1().Deployments().List()
// 具体看看为什么 clientset 有 AppsV1().Deployments().List()
deployment, err := clientset.AppsV1().Deployments(types.DefaultNamespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	
// clientset 这个结构体有 AppsV1() 方法，返回的值 clientset.AppsV1 结构体本身的数据
func (c *Clientset) AppsV1() appsv1.AppsV1Interface {
	return c.appsV1
}
// appsV1 是一个结构体，包含了 Deployment 的方法
func (c *AppsV1Client) Deployments(namespace string) DeploymentInterface {
	return newDeployments(c, namespace)
}	
// 此处通过 appsV1 结构体里的数据完善 deployments 结构体
func newDeployments(c *AppsV1Client, namespace string) *deployments {
	return &deployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}
// deployments 结构体拥有 List 方法
// 之前生成 RESTClient 时已经把网络请求的一些格式认证定义好了
func (c *deployments) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	// 定义一个 DeploymentList 结构体
	result = &v1.DeploymentList{}
	// 通过 deployments 结构体的 client 就可以发送 Get、Post、Delete 请求，这个是 Get 请求
	// 此处多个 . 目的是完善 client 之后通过 Do 函数执行， Into 是把获取的数据放到 result 结构体里
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}
```