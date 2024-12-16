package main

import (
	"fmt"
	"log"
	"net/url"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func main() {

	uri, err := url.Parse("http://minio-0.minio.default.svc.cluster.local:9000")
	fmt.Println(uri.Hostname(), err)
	// 示例的 ResourceRequirements
	resources := corev1.ResourceRequirements{
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("500m"), // 500 millicores
			corev1.ResourceMemory: resource.MustParse("500Mi"),
		},
		Limits: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("1"), // 1 core
			corev1.ResourceMemory: resource.MustParse("1Gi"),
		},
	}

	// 获取 Requests 中的 CPU 量并转换为 millicores
	cpuRequest := resources.Requests[corev1.ResourceCPU]
	cpuRequestMilliCores := cpuRequest.MilliValue()
	fmt.Printf("CPU Request: %d millicores\n", cpuRequestMilliCores)

	// 获取 Limits 中的 CPU 量并转换为 millicores
	cpuLimit := resources.Limits[corev1.ResourceCPU]
	cpuLimitMilliCores := cpuLimit.MilliValue()
	fmt.Printf("CPU Limit: %d millicores\n", cpuLimitMilliCores)

	// 获取 Requests 中的内存量并转换为字节
	memoryRequest := resources.Requests[corev1.ResourceMemory]
	memoryRequestBytes, ok := memoryRequest.AsInt64()
	if !ok {
		log.Fatalf("Failed to convert memory request to int64")
	}
	fmt.Printf("Memory Request: %d bytes\n", memoryRequestBytes)

	// 获取 Limits 中的内存量并转换为字节
	memoryLimit := resources.Limits[corev1.ResourceMemory]
	memoryLimitBytes, ok := memoryLimit.AsInt64()
	if !ok {
		log.Fatalf("Failed to convert memory limit to int64")
	}
	fmt.Printf("Memory Limit: %d bytes\n", memoryLimitBytes)
}
