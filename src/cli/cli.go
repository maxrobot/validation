// Copyright (c) 2018 Clearmatics Technologies Ltd

package cli

import (
	// "strings"
	"fmt"
	"log"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/validation/src/config"
)

type Block struct {
    Number string
}

func Launch(setup config.Setup) {
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

			// Connect the client
		  client, err := rpc.Dial("http://" + setup.Addr_to + ":" + setup.Port_to)
		  if err != nil {
		    log.Fatalf("could not create ipc client: %v", err)
		  }

			var lastBlock Block
		  err = client.Call(&lastBlock, "eth_getBlockByNumber", "latest", true)
		  if err != nil {
		      fmt.Println("can't get latest block:", err)
		      return
		  } else {
				// Print events from the subscription as they arrive.
				k, _ := strconv.ParseInt(lastBlock.Number, 0, 64)
			  fmt.Printf("latest block: %v\n", k)
			}

			c.Println("===============================================================")
		},
	})

	// run shell
	shell.Run()
}
