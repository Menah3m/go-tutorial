package v1

import (
	"github.com/menah3m/go-tutorial/k8s-crd/pkg/apis/mycrd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

/*
   @Auth: menah3m
   @Desc:
*/

var SchemeGroupVersion = schema.GroupVersion{
	Group:   mycrd.GroupName,
	Version: mycrd.Version,
}

func adKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, &Network{}, &NetworkList{})
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
