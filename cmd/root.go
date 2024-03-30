package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"

	"k8s.io/client-go/tools/clientcmd"
)

func Execute() {
	// Load kubeconfig
	kubeConfigPath := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.LoadFromFile(kubeConfigPath)
	if err != nil {
		fmt.Printf("Error loading kubeconfig: %v\n", err)
		os.Exit(1)
	}

	// Get available contexts
	var contextNames []string
	for contextName := range config.Contexts {
		contextNames = append(contextNames, contextName)
	}

	// Prompt user to select a context
	prompt := promptui.Select{
		Label: "Select a Kubernetes context",
		Items: contextNames,
	}

	_, selectedContext, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		os.Exit(1)
	}

	config.CurrentContext = selectedContext

	err = clientcmd.WriteToFile(*config, kubeConfigPath)
	if err != nil {
		fmt.Printf("Error writing kubeconfig: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Switched to cluster/context: %s\n", selectedContext)
}
