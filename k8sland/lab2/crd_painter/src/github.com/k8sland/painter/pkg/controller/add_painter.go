// Copyright 2018 Imhotep Software LLC. Apache 2.0 Licence

package controller

import (
	"github.com/k8sland/painter/pkg/controller/painter"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, painter.Add)
}
