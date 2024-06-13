package main

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const appName = "novus-whoami"

func kubernetesClient() error {
	namespace := "default"
	byteNamespace, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		fmt.Printf("Error reading namespace, uses default namespace. %v\n", err)
	} else {
		namespace = string(byteNamespace)
	}

	config, err := rest.InClusterConfig()
	if err != nil { 
		return fmt.Errorf("Error getting cluster config: %w", err)
	} 
	config.UserAgent = fmt.Sprintf("%s %s", config.UserAgent, appName)
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("Error creating Kubernetes client: %w", err)
	}
	
	ctx := context.TODO()
	opts := metav1.ListOptions{}

	pods, err := client.CoreV1().Pods(namespace).List(ctx, opts)
	if err != nil {
		fmt.Printf("Error getting pods: %v\n", err)
	} else {
		for _, pod := range pods.Items {
			fmt.Printf("Pod name: %s\n", pod.Name)
			fmt.Printf("Pod status: %s\n", pod.Status.Phase)
		}
		fmt.Println()
	}

	services, err := client.CoreV1().Services(namespace).List(ctx, opts)
	if err != nil {
		fmt.Printf("Error getting services: %v\n", err)
	} else {
		for _, service := range services.Items {
			fmt.Printf("Service name: %s\n", service.Name)
		}
		fmt.Println()
	}
	
	deployments, err := client.AppsV1().Deployments(namespace).List(ctx, opts)
	if err != nil {
		fmt.Printf("Error getting deployments: %v\n", err)
	} else {
		for _, deployment := range deployments.Items {
			fmt.Printf("Deployment name: %s\n", deployment.Name)
		}
		fmt.Println()
	}
	
	replicasets, err := client.AppsV1().ReplicaSets(namespace).List(ctx, opts)
	if err != nil {
		fmt.Printf("Error getting replicasets: %v\n", err)
	} else {
		for _, replicaset := range replicasets.Items {
			fmt.Printf("ReplicaSet name: %s\n", replicaset.Name)
		}
		fmt.Println()
	}
	
	events, err := client.CoreV1().Events(namespace).List(ctx, opts)
	if err != nil {
		fmt.Printf("Error getting events: %v\n", err)
	} else {
		for _, event := range events.Items {
			fmt.Printf("Event name: %s\n", event.Name)
		}
		fmt.Println()
	}
	
	namespaces, err := client.CoreV1().Namespaces().List(ctx, opts)
	if err != nil {
		fmt.Printf("Error getting namespaces: %v\n", err)
	} else {
		for _, namespace := range namespaces.Items {
			fmt.Printf("Namespace name: %s\n", namespace.Name)
		}
		fmt.Println()
	}
	
	nodes, err := client.CoreV1().Nodes().List(ctx, opts)
	if err != nil {
		fmt.Printf("Error getting nodes: %v\n", err)
	} else {
		for _, node := range nodes.Items {
			fmt.Printf("Node name: %s\n", node.Name)
			for _, addr := range node.Status.Addresses {
				if addr.Type == "InternalIP" {
					fmt.Printf("Node: %s, InternalIP: %s\n", node.Name, addr.Address)
				}
				if addr.Type == "ExternalIP" {
					fmt.Printf("Node: %s, ExternalIP: %s\n", node.Name, addr.Address)
				}
			}
		}
	}
	return nil
}
