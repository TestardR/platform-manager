package k8sclient

import (
	"context"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//go:generate mockgen -package=mock -source=k8sclient.go -destination=$MOCK_FOLDER/k8sclient.go Managerer

// Managerer describres methods associated with a Manager instance.
type Managerer interface {
	// GetDeployments returns services deployed in a namespace
	GetDeployments(ctx context.Context, namespace string) ([]v1.Deployment, error)
	// GetDeployments returns services deployed per application group in a namespace
	GetDeploymentsPerLabel(ctx context.Context, namespace, label, value string) ([]v1.Deployment, error)
}

// Manager holds a client to specifc kubernetes library.
type Manager struct {
	client *kubernetes.Clientset
}

// New creates a new Manager instance.
func New(kubeconfig string) (Managerer, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &Manager{client: clientset}, nil
}

func (m *Manager) GetDeployments(ctx context.Context, namespace string) ([]v1.Deployment, error) {
	dpls, err := m.client.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return dpls.Items, nil
}

func (m *Manager) GetDeploymentsPerLabel(ctx context.Context, namespace, label, value string) ([]v1.Deployment, error) {
	output := make([]v1.Deployment, 0)

	dpls, err := m.client.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: label,
	})
	if err != nil {
		return nil, err
	}

	for _, dpl := range dpls.Items {
		if dpl.GetLabels()[label] == value {
			output = append(output, dpl)
		}
	}

	return output, nil
}
