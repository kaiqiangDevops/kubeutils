/*
 * @Time : 2024/7/29 14:38
 * @Author : diehao.yuan
 * @Email : diehao.yuan@outlook.com
 * @File : namespace.go
 */
package kubeutils

import (
	"context"
	"github.com/YuanDieHao/kubeutils/utils/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// 定义结构体
type Namespace struct {
	InstanceInterface typedv1.CoreV1Interface
	Item              *corev1.Namespace
}

// New函数可以用于设置一些默认值
func NewNamespace(kubeconfig string, item *corev1.Namespace) *Namespace {
	// 首先调用instance的init函数，生成一个ResourceInstance的实例，并配置默认值和生成clientset
	instance := ResourceInstance{}
	instance.Init(kubeconfig)

	// 定义一个Namespace实例
	resource := Namespace{}
	resource.InstanceInterface = instance.Clientset.CoreV1()
	resource.Item = item
	return &resource
}

// 创建资源
func (c *Namespace) Create(namespace string) error {
	log.Infof("Name: ", c.Item.Name, "Create Namespace!")
	_, err := c.InstanceInterface.Namespaces().Create(context.TODO(), c.Item, metav1.CreateOptions{})
	return err
}

// 删除资源
func (c *Namespace) Delete(namespace, name string, gracePeriodSeconds *int64) error {
	log.Warnf("Name: ", name, "Delete Namespace!")
	deleteOptions := metav1.DeleteOptions{}

	// gracePeriodSeconds可配置，如果为0代表是强制删除
	if gracePeriodSeconds != nil {
		// 说明传递了gracePeriodSeconds
		deleteOptions.GracePeriodSeconds = gracePeriodSeconds
	}
	err := c.InstanceInterface.Namespaces().Delete(context.TODO(), name, deleteOptions)
	return err
}

// 删除多个资源
func (c *Namespace) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	// 删除多个时，结构体会接收一个nameList的切片，循环该切片，然后调用Delete函数即可
	for _, name := range nameList {
		// 调用删除函数
		c.Delete("", name, gracePeriodSeconds)
	}
	return nil
}

// 更新资源
func (c *Namespace) Update(namespace string) error {
	log.Warnf("Name: ", c.Item.Name, "Update Namespace!")
	_, err := c.InstanceInterface.Namespaces().Update(context.TODO(), c.Item, metav1.UpdateOptions{})
	return err
}

// 获取资源列表
func (c *Namespace) List(namespace, labelSelector, fieldSelector string) (items interface{}, err error) {
	log.Infof("Get Namespace List!")
	// 有可能是根据条件进行查询
	listOptions := metav1.ListOptions{
		FieldSelector: fieldSelector,
		LabelSelector: labelSelector,
	}
	list, err := c.InstanceInterface.Namespaces().List(context.TODO(), listOptions)
	items = list.Items
	return items, err
}

// 获取资源详情
func (c *Namespace) Get(namespace, name string) (item interface{}, err error) {
	log.Infof("Name: ", name, "Get Namespace Info!")
	i, err := c.InstanceInterface.Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	i.APIVersion = "v1"
	i.Kind = "Namespace"
	item = i
	return item, err
}
