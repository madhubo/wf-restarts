package main

import (
	"bufio"
	"flag"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	//"sort"
	//"time"

	//"k8s.io/client-go/util/retry"
	"os"
	"path/filepath"

	//appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//"github.com/satori/go.uuid"
	//"k8s.io/client-go/util/retry"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}


	//deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	//podClient := clientset.RESTClient().Get().Namespace(apiv1.NamespaceDefault)


	//deployment := &appsv1.Deployment{
	//	ObjectMeta: metav1.ObjectMeta{
	//		Name: "demo-deployment-" + uuid.NewV4().String(),
	//	},
	//	Spec: appsv1.DeploymentSpec{
	//		Replicas: int32Ptr(2),
	//		Selector: &metav1.LabelSelector{
	//			MatchLabels: map[string]string{
	//				"app": "demo",
	//			},
	//		},
	//
	//		Template: apiv1.PodTemplateSpec{
	//			ObjectMeta: metav1.ObjectMeta{
	//				Labels: map[string]string{
	//					"app": "demo",
	//				},
	//			},
	//			Spec: apiv1.PodSpec{
	//				Containers: []apiv1.Container{
	//					{
	//						Name:  "web",
	//						Image: "nginx:1.12",
	//						Ports: []apiv1.ContainerPort{
	//							{
	//								Name:          "http",
	//								Protocol:      apiv1.ProtocolTCP,
	//								ContainerPort: 80,
	//							},
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}

	//// Create Deployment
	//fmt.Println("Creating deployment...")
	//result, err := deploymentsClient.Create(deployment)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	//
	//// Update Deployment
	//prompt()
	//fmt.Println("Updating deployment...")
	////    You have two options to Update() this Deployment:
	////
	////    1. Modify the "deployment" variable and call: Update(deployment).
	////       This works like the "kubectl replace" command and it overwrites/loses changes
	////       made by other clients between you Create() and Update() the object.
	////    2. Modify the "result" returned by Get() and retry Update(result) until
	////       you no longer get a conflict error. This way, you can preserve changes made
	////       by other clients between Create() and Update(). This is implemented below
	////			 using the retry utility package included with client-go. (RECOMMENDED)
	////
	//// More Info:
	//// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
	//
	//retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
	//	// Retrieve the latest version of Deployment before attempting update
	//	// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
	//	result, getErr := deploymentsClient.Get("demo-deployment", metav1.GetOptions{})
	//	if getErr != nil {
	//		panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
	//	}
	//
	//	result.Spec.Replicas = int32Ptr(1)                           // reduce replica count
	//	result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13" // change nginx version
	//	_, updateErr := deploymentsClient.Update(result)
	//	return updateErr
	//})
	//if retryErr != nil {
	//	panic(fmt.Errorf("Update failed: %v", retryErr))
	//}
	//fmt.Println("Updated deployment...")

	// List Deployments
	prompt()

	listAllDeployments(clientset)

	//for _, d1 := range list.Items {
	//
	//	replicaClient :=
	//	repList, err1 := replicaClient.List(metav1.ListOptions{})
	//	if err1 != nil {
	//		panic(err1)
	//	}
	//	for _, d2 := range repList.Items {
	//		fmt.Printf(" * %s (%d replicas)\n", d2.Name, d2.Status)
	//	}
	//



/*
	// Delete Deployment
	prompt()


 */
}

// Check if nodes have pods

func deleteDeployment(clientset *kubernetes.Clientset) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	fmt.Println("Deleting deployment...")
	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentsClient.Delete("demo-deployment", &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	fmt.Println("Deleted deployment.")
}

func getAllPods(clientset *kubernetes.Clientset) (*apiv1.PodList, error) {

	// get all pods
	fmt.Println("Printing all pods...")
	pods, err := clientset.CoreV1().Pods("default").List(metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

//	sort.Slice(pods, func(i, j int) bool { return pods[i].GetCreationTimestamp() < pods[i].GetCreationTimestamp()})

	fmt.Println("Done printing all pods...")
	return pods, err
}

func getOldestPod(clientset *kubernetes.Clientset) {
	//pods , err := getAllPods(clientset)
	//if err != nil {
	//	panic(err)
	//}
	//for _, pod := range pods.Items {
	//	//fmt.Printf(" pod:  * %s \n", pod.Status);
	//	age := time.Since(pod.GetCreationTimestamp().Time).Round(time.Second)
	//	sort.Slice(pods, func(i, j int) bool { return pods[i].GetCreationTimestamp() < pods[i].GetCreationTimestamp()})
	//	fmt.Printf(" Pod Name: {%s}, pod Age:  * %s \n", pod.GetName(), age);
	//}
	////
	//
}

func listAllDeployments(clientset *kubernetes.Clientset){
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	list, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Printing all deployments...")
	for _, d := range list.Items {

		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	}
}

func getDeployment(clientset *kubernetes.Clientset) (*v1.Deployment, error) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)

	return deploymentsClient.Get("wf-query-service", metav1.GetOptions{})

}

func doNodesHavePods(clientset *kubernetes.Clientset) error {
	nodeLabelSelector := "nodelabel=interesting_nodes"
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{LabelSelector: nodeLabelSelector})

	if err != nil {
		return err
	}

	nodeNames := []string{}
	for _, node := range nodes.Items {
		nodeNames = append(nodeNames, node.Name)
	}
	// --all-namespaces -> listing and looping on namespaces
	namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})

	if err != nil {
		return err
	}
	for _, namespace := range namespaces.Items {
		for _, name := range nodeNames {
			// pods need a namespace to be listed.
			pods, err := clientset.CoreV1().Pods(namespace.Name).List(metav1.ListOptions{FieldSelector: "spec.nodeName=" + name})
			if err != nil {
				println("%v", err)
			}
			for _, pod := range pods.Items {
				fmt.Println(pod.Namespace, pod.Name)
			}
		}
	}
	return nil
}



func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}

func int32Ptr(i int32) *int32 { return &i }