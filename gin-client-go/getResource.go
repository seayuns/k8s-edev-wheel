package main

import (
	"context"
	"flag"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog"
)

func getResource() {
	var kubeConfig *string
	ctx := context.Background()
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "absoute path to the kubeConfig file")
	} else {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", ""), "absoute path to the kubeConfig file")
	}
	flag.Parse()

	confg, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		klog.Fatal("read kube config error", err)
	}

	clientSet, err := kubernetes.NewForConfig(confg)
	if err != nil {
		klog.Fatal("clientset error", err)
	}

	//get all namespace
	namespaceList, err := clientSet.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
	if err != nil {
		klog.Fatal("read cluster ns", err)
	}

	namespaces := namespaceList.Items
	for i := 0; i < len(namespaces); i++ {
		klog.Info("clustes ns  name ==> ", namespaces[i].Name,
			"status ==> ", namespaces[i].Status.Phase,
		)
	}

	//get pod default
	podlist, err := clientSet.CoreV1().Pods("default").List(ctx, v1.ListOptions{})
	if err != nil {
		klog.Fatal("read cluster pods", err)
	}

	pods := podlist.Items
	for i := 0; i < len(pods); i++ {
		klog.Info("clustes pod  name ==> ", pods[i].Name,
			"status ==> ", pods[i].Status.Phase,
		)
	}
}
