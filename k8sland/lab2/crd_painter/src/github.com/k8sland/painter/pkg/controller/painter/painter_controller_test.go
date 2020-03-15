// Copyright 2018 Imhotep Software LLC. Apache 2.0 Licence

package painter

import (
	"testing"
	"time"

	workloadv1alpha1 "github.com/k8sland/painter/pkg/apis/workload/v1alpha1"
	"github.com/onsi/gomega"
	"golang.org/x/net/context"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var c client.Client

var expectedRequest = reconcile.Request{NamespacedName: types.NamespacedName{Name: "foo", Namespace: "default"}}
var depKey = types.NamespacedName{Name: "foo-deployment", Namespace: "default"}

const timeout = time.Second * 5

func TestReconcile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	instance := &workloadv1alpha1.Painter{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "painter",
			Namespace: "default",
			Labels: map[string]string{
				"cr": "painter",
			},
		},
		Spec: workloadv1alpha1.PainterSpec{Color: "Green"},
	}

	mgr, err := manager.New(cfg, manager.Options{})
	g.Expect(err).NotTo(gomega.HaveOccurred())
	c = mgr.GetClient()

	recFn, requests := SetupTestReconcile(newReconciler(mgr))
	g.Expect(add(mgr, recFn)).NotTo(gomega.HaveOccurred())
	defer close(StartTestManager(mgr, g))

	// Create the Paint object
	err = c.Create(context.TODO(), instance)
	if apierrors.IsInvalid(err) {
		t.Logf("failed to create object, got an invalid object error: %v", err)
		return
	}
	g.Expect(err).NotTo(gomega.HaveOccurred())
	defer c.Delete(context.TODO(), instance)
	eReq := reconcile.Request{
		NamespacedName: types.NamespacedName{Name: "painter", Namespace: "default"},
	}
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(eReq)))

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod",
			Namespace: "default",
			Labels: map[string]string{
				"app": "nginx",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx",
					Image: "nginx",
				},
			},
		},
	}

	// var ns = types.NamespacedName{Name: "foo", Namespace: "default"}
	opts := client.ListOptions{
		Namespace:     "default",
		LabelSelector: labels.Set{"cr": "painter"}.AsSelector(),
	}
	var bb workloadv1alpha1.PainterList
	g.Eventually(func() error { return c.List(context.TODO(), &opts, &bb) }, timeout).
		Should(gomega.Succeed())

	err = c.Create(context.TODO(), pod)
	if apierrors.IsInvalid(err) {
		t.Logf("failed to create pod %v", err)
	}
	g.Expect(err).NotTo(gomega.HaveOccurred())
	defer c.Delete(context.TODO(), pod)

	depKey := types.NamespacedName{Name: "pod", Namespace: "default"}
	rreq := reconcile.Request{NamespacedName: depKey}
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(rreq)))
	g.Eventually(func() error { return c.Get(context.TODO(), depKey, pod) }, timeout).
		Should(gomega.Succeed())

	var p corev1.Pod
	err = c.Get(context.TODO(), depKey, &p)
	if err != nil {
		t.Logf("crapped out looking up pod")
	}
	g.Expect(p.ObjectMeta.Labels["color"]).Should(gomega.Equal("Green"))
}
