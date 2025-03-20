package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubectl/pkg/scheme"
)

func main() {
	// 1. 加载 kubeconfig 文件
	var err error
	var config *rest.Config
	var kubeconfig *string

	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// 使用 ServiceAccount 创建集群配置（InCluster模式）
	if config, err = rest.InClusterConfig(); err != nil {
		// 使用 KubeConfig 文件创建集群配置
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
			panic(err.Error())
		}
	}

	// 创建 clientset
	clientset, err := kubernetes.NewForConfig(config)

	// 3. 定义要 apply 的 Deployment 对象
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "example-deployment",
			Namespace: "default",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "example",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "example",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx:1.14.2",
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// 4. 将对象序列化为 JSON
	serializer := serializer.NewCodecFactory(scheme.Scheme).LegacyCodec(appsv1.SchemeGroupVersion)
	deploymentJSON, err := runtime.Encode(serializer, deployment)
	if err != nil {
		log.Fatalf("Failed to encode deployment: %v", err)
	}

	// 5. 使用 Patch 方法 apply 对象
	_, err = clientset.AppsV1().Deployments("default").Patch(
		context.TODO(),
		deployment.Name,
		types.ApplyPatchType,
		deploymentJSON,
		metav1.PatchOptions{
			FieldManager: "example-client", // 指定 field manager
			Force:        boolPtr(true),    // 强制覆盖冲突字段
		},
	)
	if err != nil {
		log.Fatalf("Failed to apply deployment: %v", err)
	}

	fmt.Println("Deployment applied successfully!")
}

// 辅助函数：返回 int32 指针
func int32Ptr(i int32) *int32 {
	return &i
}

// 辅助函数：返回 bool 指针
func boolPtr(b bool) *bool {
	return &b
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
