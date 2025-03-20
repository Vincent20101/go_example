package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1apply "k8s.io/client-go/applyconfigurations/apps/v1"
	corev1apply "k8s.io/client-go/applyconfigurations/core/v1"
	metav1apply "k8s.io/client-go/applyconfigurations/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1. 加载 kubeconfig
	kubeconfig := flag.String("kubeconfig", "", "Path to kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// 2. 创建 Clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 3. 定义要 Apply 的 Deployment 对象（示例）
	deployName := "my-nginx"
	namespace := "default"

	// 构造 Apply 配置（结构化方式）
	applyConfig := appsv1apply.Deployment(deployName, namespace).
		WithLabels(map[string]string{"app": "nginx"}).
		WithSpec(appsv1apply.DeploymentSpec().
			WithReplicas(2).
			WithSelector(metav1apply.LabelSelector().
				WithMatchLabels(map[string]string{"app": "nginx"}),
			).
			WithTemplate(corev1apply.PodTemplateSpec().
				WithLabels(map[string]string{"app": "nginx"}).
				WithSpec(corev1apply.PodSpec().
					WithContainers(corev1apply.Container().
						WithName("nginx").
						WithImage("nginx:1.21").
						WithPorts(corev1apply.ContainerPort().
							WithContainerPort(80),
						),
					),
				),
			),
		)

	// 4. 执行 Server-Side Apply
	result, err := clientset.AppsV1().Deployments(namespace).Apply(
		context.TODO(),
		applyConfig,
		metav1.ApplyOptions{
			FieldManager: "my-client", // 标识操作者，用于冲突管理
			Force:        true,        // 强制覆盖冲突（可选）
		},
	)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Applied Deployment %q\n", result.GetName())
}
