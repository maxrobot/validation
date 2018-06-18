// Copyright (c) 2018 Clearmatics Technologies Ltd

package main

import (
	"fmt"
	"os"

	"github.com/validation/cli"
)

func main() {
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)
	// setup, photon := config.ParseParameters(argsWithoutProg)

	// auth, acc := api.Init(setup)

	cli.Launch()
}
