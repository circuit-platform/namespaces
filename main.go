package main

import (
	"github.com/circuit-platform/namespaces/model"
	utils "github.com/circuit-platform/models-utils"
)

func main() {
	utils.Run(model.CreateNamespacesIndex, nil)
}