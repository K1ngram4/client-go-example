package main

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

var namespace = "kube-system"

func main() {

	config, err := clientcmd.BuildConfigFromFlags("", "../config")
	if err != nil {
		panic(err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	// 定义组版本资源
	gvr := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
	unStructObj, err := dynamicClient.Resource(gvr).Namespace(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	podList := &apiv1.PodList{}

	if err = runtime.DefaultUnstructuredConverter.FromUnstructured(unStructObj.UnstructuredContent(), podList); err != nil {
		panic(err)
	}

	for _, v := range podList.Items {
		fmt.Printf("namespaces:%v  name:%v status:%v \n", v.Namespace, v.Name, v.Status.Phase)
	}
}
