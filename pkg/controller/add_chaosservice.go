package controller

import (
	"github.com/skornehl/cc-chaos/pkg/controller/chaosservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, chaosservice.Add)
}
