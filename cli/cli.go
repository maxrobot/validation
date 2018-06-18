// Copyright (c) 2018 Clearmatics Technologies Ltd

package cli

import (
	// "strings"

	"github.com/abiosoft/ishell"
	// "gitlab.clearmatics.net/dev/boe-poc/src/api"
	// "gitlab.clearmatics.net/dev/boe-poc/src/config"
)

func Launch() {
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println("Block Validation CLI Tool")

	// display specific account balance
	shell.AddCmd(&ishell.Cmd{
		Name: "getBlock",
		Help: "get an ethereum block from a chain",
		Func: func(c *ishell.Context) {
			c.Println("===============================================================")
			c.Println("Get block:")
			if len(c.Args) > 1 {
				c.Println("Too many arguments entered.")
			} else {
				c.Println(c.Args[0])
			}
			c.Println("===============================================================")
		},
	})

	// run shell
	shell.Run()
}
