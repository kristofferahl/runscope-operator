package controller

import (
	"github.com/kristofferahl/runscope-operator/pkg/controller/runscopeagent"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, runscopeagent.Add)
}
