package main

import (
	"fmt"

	"github.com/mayadata.io/kubera-backup-restore/pkg/backup"
	"github.com/mayadata.io/kubera-backup-restore/pkg/client"
	kuberadiscovery "github.com/mayadata.io/kubera-backup-restore/pkg/discovery"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func main() {
	fmt.Println("vim-go")
	config, err := client.LoadConfig()
	if err != nil {
		fmt.Printf("Error Loading config: %+v\n", err)
		return
	}
	f := client.NewFactory(config)
	dynamicClient, err := f.DynamicClient()
	if err != nil {
		fmt.Printf("Error dynamic client: %+v\n", err)
		return
	}
	kuberaClient, err := f.Client()
	if err != nil {
		fmt.Printf("Error kubera client: %+v\n", err)
		return
	}
	dh, err := kuberadiscovery.NewHelper(kuberaClient.Discovery())
	if err != nil {
		fmt.Printf("Discovery error %+v\n", err)
		return
	}
	df := client.NewDynamicFactory(dynamicClient)

	server := backup.NewServer(dh, df)
	// objArray, err := server.GeResourcesFromNamespace("openebs")
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }
	// for _, objList := range objArray {
	// 	for _, obj := range objList.Items {
	// 		fmt.Printf("%+v\n\n", obj)
	// 	}
	// }
	// gv := schema.GroupVersion{Version: "v1"}
	// resource := metav1.APIResource{
	// 	Name:         "namespaces",
	// 	SingularName: "ns",
	// 	Namespaced:   false,
	// 	Version:      "v1",
	// 	ShortNames:   []string{"ns"},
	// }
	// dClient, _ := df.ClientForGroupVersionResource(gv, resource, "", context.TODO())
	// objects, err := dClient.List(metav1.ListOptions{})
	// if err != nil {
	// 	fmt.Println("error in namespace fetching", err)
	// 	return
	// }
	// fmt.Printf("%+v\n", objects.Items)

	// To convert from LabelSelector to listOption populate listOptions.LabelSelector
	// Ex: sel, err := metav1.LabelSelectorAsSelector(labelSelector)
	// if err != nil { //Handle error }
	// metav1.ListOptions.LabelSelector = sel
	listOptions := metav1.ListOptions{
		LabelSelector: "test1=key1",
	}
	objList, err := server.GetResourcesFromGVR(schema.GroupVersionResource{Version: "", Resource: "namespaces"}, "", listOptions)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("Length of resources: ", len(objList.Items))
	for _, obj := range objList.Items {
		fmt.Printf("%+v\n", obj)
	}

	// // resourceList := dh.GetNamespaceScopedAPIResourceList()
	// for _, apiResourceList := range dh.GetNamespaceScopedAPIResourceList() {
	// 	for _, resource := range apiResourceList.APIResources {
	// 		fmt.Println("Resource GV: ", apiResourceList.GroupVersion, " Name: ", resource.Name, " Group: ", resource.Group)
	// 		gv, err := schema.ParseGroupVersion(apiResourceList.GroupVersion)
	// 		if err != nil {
	// 			fmt.Println("Failed to parse gv", apiResourceList.GroupVersion, err)
	// 			continue
	// 		}
	// 		dynamic, _ := df.ClientForGroupVersionResource(gv, resource, "openebs", context.TODO())
	// 		objList, err := dynamic.List(metav1.ListOptions{})
	// 		if err != nil {
	// 			fmt.Println("Failed to list resource", resource.Name, err)
	// 			continue
	// 		}
	// 		fmt.Printf("%+v\n", objList)
	// 	}
	// }

	// if err != nil {
	// 	fmt.Printf("Error while fetching the resource %+v \n", err)
	// 	return
	// }
	// _, err = df.Get("kube-dns-5b6487d4cd-66bhs", metav1.GetOptions{})
}
