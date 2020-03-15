package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	app       = "PartyScheduler"
	namespace = "default"
	schedName = "partysched"
)

var (
	// Version set via build tags
	Version = ""
	srv     kubernetes.Interface

	attires = []string{"ghoul", "goblin"}
	rootCmd = &cobra.Command{
		Use:   strings.ToLower(app),
		Short: "Schedules pods based on costumes",
		Long:  "Schedules pods based on costumes",
		Run:   listen,
	}
)

// Execute runs the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Version = Version
}

func setupSignal(cancel context.CancelFunc) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-signalChan:
			log.Printf("Shutdown signal received, exiting...")
			cancel()
			os.Exit(0)
		}
	}
}

func listen(cmd *cobra.Command, args []string) {
	opts := metav1.ListOptions{
		FieldSelector: "spec.nodeName=",
	}
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	go setupSignal(cancel)

	c := make(chan *corev1.Pod)
	go scheduler(ctx, c)

	// Watch pods in the given namespace
	!!YOUR_CODE!!

	log.Printf("%s is listening for pods...", app)
	for e := range w.ResultChan() {
		p := e.Object.(*corev1.Pod)
		if p.Spec.SchedulerName != schedName {
			continue
		}
		log.Println("Unscheduled pod detected", p.ObjectMeta.Name)
		switch e.Type {
		case watch.Added:
			c <- p
		}
	}
}

func scheduler(ctx context.Context, c <-chan *corev1.Pod) {
	for {
		select {
		case p := <-c:
			var (
				nodes corev1.NodeList
				err   error
			)
			if nodes, err = checkFit(p); err != nil {
				log.Println(err)
				break
			}
			if len(nodes.Items) == 0 {
				log.Printf("Unable to fit pod %s on any nodes", p.ObjectMeta.Name)
				break
			}
			for _, n := range nodes.Items {
				if err := schedule(n, p); err != nil {
					log.Fatal(err)
				}
				break
			}
		case <-ctx.Done():
			return
		}
	}
}

func schedule(n corev1.Node, p *corev1.Pod) error {
	body := map[string]interface{}{
		"target": map[string]string{
			"kind":       "Node",
			"apiVersion": "v1",
			"name":       n.ObjectMeta.Name,
			"namespace":  namespace,
		},
		"metadata": map[string]string{
			"name":      p.ObjectMeta.Name,
			"namespace": namespace,
		},
	}
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	bind := apiServer().
		CoreV1().
		RESTClient().
		Post().
		Namespace(namespace).
		Resource("bindings").
		Body(data)
	var b corev1.Binding
	if err := bind.Do().Into(&b); err != nil {
		return err
	}
	log.Printf("Ye! scheduled pod %s onto node %s", p.ObjectMeta.Name, n.ObjectMeta.Name)
	return nil
}

func checkFit(p *corev1.Pod) (corev1.NodeList, error) {
	candidates := corev1.NodeList{}

  var (
		costume string
		ok bool
	)
	// Check if given pod has a costume on it? If Not log and bail out
	!!YOUR_CODE!!

	nn, err := getNodes()
	if err != nil {
		return candidates, err
	}

	for _, n := range nn.Items {
		// Check if pod has the correct costume, if so pod can be scheduled on that
		// current node otherwise log message
		!!YOUR_CODE!!
	}
	return candidates, nil
}

func trickOrTreat(costume string) bool {
	for _, c := range attires {
		if c == costume {
			return true
		}
	}
	return false
}

func getNodes() (*corev1.NodeList, error) {
	return apiServer().CoreV1().Nodes().List(metav1.ListOptions{})
}

func apiServer() kubernetes.Interface {
	if srv != nil {
		return srv
	}

	var (
		cfg *restclient.Config
		err error
	)
	path := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	if _, err = os.Stat(path); os.IsNotExist(err) {
		cfg, err = restclient.InClusterConfig()
	} else {
		cfg, err = clientcmd.BuildConfigFromFlags("", path)
	}
	if err != nil {
		log.Panic(err)
	}

	srv, err = kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Panic(err)
	}

	return srv
}
