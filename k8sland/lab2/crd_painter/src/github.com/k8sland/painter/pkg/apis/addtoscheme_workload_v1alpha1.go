// Copyright 2018 Imhotep Software LLC. Apache 2.0 Licence

package apis

import (
	"github.com/k8sland/painter/pkg/apis/workload/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
}
