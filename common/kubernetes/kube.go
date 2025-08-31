package kubernetes

import (
	"context"

	"cth.release/common/utils"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sClient struct {
	clientset *kubernetes.Clientset
}

func NewK8sClient() (*K8sClient, error) {
	configData := utils.GetConfig()

	config, err := clientcmd.BuildConfigFromFlags("", configData.KubeConfig)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &K8sClient{clientset: clientset}, nil
}

func (c *K8sClient) CreatePod(namespace string, pod *corev1.Pod) (*corev1.Pod, error) {
	return c.clientset.CoreV1().Pods(namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
}

func (c *K8sClient) UpdatePod(namespace string, pod *corev1.Pod) (*corev1.Pod, error) {
	return c.clientset.CoreV1().Pods(namespace).Update(context.TODO(), pod, metav1.UpdateOptions{})
}

func (c *K8sClient) DeletePod(namespace, name string) error {
	return c.clientset.CoreV1().Pods(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (c *K8sClient) GetPod(namespace, name string) (*corev1.Pod, error) {
	return c.clientset.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

func (c *K8sClient) ListPods(namespace string) (*corev1.PodList, error) {
	return c.clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
}

func (c *K8sClient) CreateService(namespace string, service *corev1.Service) (*corev1.Service, error) {
	return c.clientset.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
}

func (c *K8sClient) UpdateService(namespace string, service *corev1.Service) (*corev1.Service, error) {
	return c.clientset.CoreV1().Services(namespace).Update(context.TODO(), service, metav1.UpdateOptions{})
}

func (c *K8sClient) DeleteService(namespace, name string) error {
	return c.clientset.CoreV1().Services(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (c *K8sClient) GetService(namespace, name string) (*corev1.Service, error) {
	return c.clientset.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

func (c *K8sClient) ListServices(namespace string) (*corev1.ServiceList, error) {
	return c.clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
}

func (c *K8sClient) CreateDeployment(namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	return c.clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
}

func (c *K8sClient) UpdateDeployment(namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	return c.clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
}

func (c *K8sClient) DeleteDeployment(namespace, name string) error {
	return c.clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (c *K8sClient) GetDeployment(namespace, name string) (*appsv1.Deployment, error) {
	return c.clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

func (c *K8sClient) ListDeployments(namespace string) (*appsv1.DeploymentList, error) {
	return c.clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
}
