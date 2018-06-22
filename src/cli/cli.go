// Copyright (c) 2018 Clearmatics Technologies Ltd

package cli

import (
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

	// Connect to the RPC Client
	client, err := rpc.Dial("http://" + setup.Addr_to + ":" + setup.Port_to)
	if err != nil {
		log.Fatalf("could not create RPC client: %v", err)
	} else {
		shell.Println("Listening on RPC Client: " + setup.Addr_to + ":" + setup.Port_to)
	}

	// Get the latest block number
	shell.AddCmd(&ishell.Cmd{
		Name: "getLatestBlock",
		Help: "Gets the latest block number of Ethereum instance specified as from",
		Func: func(c *ishell.Context) {
			c.Println("===============================================================")
			c.Println("Get latest block number:")
			latestBlock(client)
			c.Println("===============================================================")
		},
	})

	// Get block N
	shell.AddCmd(&ishell.Cmd{
		Name: "getBlock",
		Help: "Gets specific block of Ethereum instance specified as from",
		Func: func(c *ishell.Context) {
			c.Println("===============================================================")
			if len(c.Args) == 0 {
				c.Println("Choose a block.")
			} else if len(c.Args) > 1 {
				c.Println("Too many arguments entered.")
			} else {
				block := strToHex(c.Args[0])
				getBlock(client, block)
			}
			c.Println("===============================================================")
		},
	})

	// Get block N and spew out the RLP encoded block
	shell.AddCmd(&ishell.Cmd{
		Name: "rlpEncodeBlock",
		Help: "Request RLP encoded block [N] of chain [from]",
		Func: func(c *ishell.Context) {
			c.Println("===============================================================")
			if len(c.Args) == 0 {
				c.Println("Choose a block.")
			} else if len(c.Args) > 1 {
				c.Println("Too many arguments entered.")
			} else {
				block := strToHex(c.Args[0])
				c.Println("RLP encode block: " + c.Args[0])
				rlpEncodeBlock(client, block)
			}
			c.Println("===============================================================")
		},
	})

	// Get block N, output three items:
	// * RLP encoded blockHeader
	// * Prefix for signed blockHeader
	// * Prefix for extraData minus signatures
	shell.AddCmd(&ishell.Cmd{
		Name: "getValidBlock",
		Help: "Request block [N] from chain [from], calculates the prefixes required for submission to chain [to]",
		Func: func(c *ishell.Context) {
			c.Println("===============================================================")
			if len(c.Args) == 0 {
				c.Println("Choose a block.")
			} else if len(c.Args) > 1 {
				c.Println("Too many arguments entered.")
			} else {
				block := strToHex(c.Args[0])
				c.Println("RLP encode block: " + c.Args[0])
				calculateRlpEncoding(client, block)
			}
			c.Println("===============================================================")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "submitSignedBlock",
		Help: "Request block [N] from chain [from], calculates the prefixes required for submission to chain [to]",
		Func: func(c *ishell.Context) {
			c.Println("===============================================================")
			if len(c.Args) == 0 {
				c.Println("Choose a block.")
			} else if len(c.Args) > 1 {
				c.Println("Too many arguments entered.")
			} else {
				block := strToHex(c.Args[0])
				c.Println("RLP encode block: " + c.Args[0])
				calculateRlpEncoding(client, block)
			}
			c.Println("===============================================================")
		},
	})

	// run shell
	shell.Run()
}

func strToHex(input string) (output string) {
	val, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("please input decimal:", err)
		return
	}
	output = strconv.FormatInt(int64(val), 16)

	return "0x" + output
}
