package controller

import (
	"github.com/yaoliu/yao-operator/pkg/controller/yaopod"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, yaopod.Add)
}
