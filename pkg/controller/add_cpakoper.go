package controller

import (
	"github.com/vbudi000/cpakoper/pkg/controller/cpakoper"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, cpakoper.Add)
}
