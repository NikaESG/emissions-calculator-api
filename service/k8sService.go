package service

import (
	"bytes"
	"context"
	_ "encoding/json"
	"errors"
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"path/filepath"
	"strings"
	"time"
	"trino.com/trino-connectors/data"
	"trino.com/trino-connectors/util"

	"github.com/gin-gonic/gin"
	// "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/util/homedir"
)

var k8sclient *kubernetes.Clientset

func InitK8sClient() {

	var k8sconfig *restclient.Config

	if gin.Mode() == "release" {
		// creates the in-cluster config, and we use release mode in Docker
		var err error
		k8sconfig, err = restclient.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	} else {
		// creates the ouf-of-cluster config
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()
		// use the current context in kubeconfig
		var err error
		k8sconfig, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	}

	// create the clientset
	var err error
	k8sclient, err = kubernetes.NewForConfig(k8sconfig)
	if err != nil {
		panic(err.Error())
	}

	// Uncomment to test locally if kubeconfig is set properly
	for i := 1; i <= 1; i++ {
		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := k8sclient.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		// _, err = k8sclient.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
		// if errors.IsNotFound(err) {
		// 	fmt.Printf("Pod example-xxxxx not found in default namespace\n")
		// } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		// 	fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		// } else if err != nil {
		// 	panic(err.Error())
		// } else {
		// 	fmt.Printf("Found example-xxxxx pod in default namespace\n")
		// }

		time.Sleep(3 * time.Second)
	}
}

func ListTrinoCatalogData() (configMapData map[string]string, err error) {
	configMap, err := k8sclient.CoreV1().ConfigMaps(K8sNamespace).Get(context.TODO(), TrinoConfigMapName, metav1.GetOptions{})
	if err != nil {
		return configMapData, err
	}

	return configMap.Data, nil
}

//func CreateTrinoCatalogData(catalogId string, config map[string]string) (configMapData map[string]data.ConfigMapData, err error) {
//
//	jsonString, err := json.Marshal(config)
//	if err != nil {
//		return configMapData, err
//	}
//
//	configmap := v1.ConfigMap{
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      catalogId + ".properties",
//			Namespace: K8sNamespace,
//		},
//		Data: map[string]string{
//			catalogId + ".properties": string(jsonString),
//		},
//		BinaryData: nil,
//	}
//
//	configMap, err := k8sclient.CoreV1().ConfigMaps(K8sNamespace).Create(context.TODO(), &configmap, metav1.CreateOptions{})
//	if err != nil {
//		return configMapData, err
//	}
//	configMapData = make(map[string]data.ConfigMapData)
//
//	configMapData[configMap.ObjectMeta.Name] = configMap.Data
//
//	return configMapData, nil
//}

func DeleteTrinoCatalogData(catalogId string, config map[string]string) (configMapData map[string]data.ConfigMapData, err error) {

	delete(config, catalogId)

	return UpdateTrinoCatalog(config)
}

func UpdateTrinoCatalogData(catalogId string, config map[string]string, origin map[string]string) (configMapData map[string]data.ConfigMapData, err error) {

	configData := make(map[string]string)
	configDetail := ""
	for i, v := range config {
		configDetail = configDetail + fmt.Sprintf("%v=%v\n", i, v)
	}
	configData[catalogId] = configDetail[:len(configDetail)-1]
	config = util.Merge(configData, origin)

	err = RestartDeployments([]string{
		TrinoCoordinator, TrinoWorker,
	})

	if err != nil {
		return
	}

	return UpdateTrinoCatalog(config)
}

func UpdateTrinoCatalog(detail map[string]string) (configMapData map[string]data.ConfigMapData, err error) {

	configmap := v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      TrinoConfigMapName,
			Namespace: K8sNamespace,
		},
		Data:       detail,
		BinaryData: nil,
	}

	configMap, err := k8sclient.CoreV1().ConfigMaps(K8sNamespace).Update(context.TODO(), &configmap, metav1.UpdateOptions{})

	configMapData = make(map[string]data.ConfigMapData)

	configMapData[configMap.ObjectMeta.Name] = configMap.Data

	return configMapData, nil
}

func RestartDeployments(deploymentNames []string) (err error) {

	for _, v := range deploymentNames {

		result, getErr := k8sclient.AppsV1().Deployments(K8sNamespace).Get(context.TODO(), v, metav1.GetOptions{})
		if getErr != nil {
			return errors.New(fmt.Sprintf("[RestartDeployments] Failed to get latest configuration of deployment: %v, %v", v, getErr))
		}

		result.Spec.Template.Annotations = map[string]string{"startTime": time.Now().Format(time.RFC3339)}

		_, err = k8sclient.AppsV1().Deployments(K8sNamespace).Update(context.TODO(), result, metav1.UpdateOptions{})

		if err != nil {
			return errors.New(fmt.Sprintf("[RestartDeployments] Failed to update deployments: %v, %v", v, err))
		}
	}
	return nil
}

func ExecPod(podName string, cmd string) (m map[string]string, e error) {

	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		return
	}

	pods, err := k8sclient.CoreV1().Pods(K8sNamespace).List(context.TODO(), metav1.ListOptions{})
	p := v1.Pod{}

	for _, v := range pods.Items {
		if strings.Index(v.Name, podName) != -1 {
			p = v
			break
		}
	}

	return cmdExecuter(restconfig, p.Name, p.Namespace, cmd)
}

// 入参为kubeconfig、pod名字、命名空间、命令字符串
func cmdExecuter(config *restclient.Config, podName, namespace, cmd string) (map[string]string, error) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	// 构造执行命令请求
	req := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec").
		VersionedParams(&v1.PodExecOptions{
			Command: []string{"bin/bash", "-c", cmd},
			Stdin:   true,
			Stdout:  true,
			Stderr:  true,
			TTY:     false,
		}, scheme.ParameterCodec)
	// 执行命令
	executor, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		return nil, err
	}
	// 使用bytes.Buffer变量接收标准输出和标准错误
	var stdout, stderr bytes.Buffer
	err = executor.Stream(remotecommand.StreamOptions{
		Stdin:  strings.NewReader(""),
		Stdout: &stdout,
		Stderr: &stderr,
	})

	if err != nil {
		return nil, err
	}
	// 返回数据
	ret := map[string]string{
		"SCHEMAS": stdout.String(),
	}
	return ret, nil
}
