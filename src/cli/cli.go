// Copyright (c) 2018 Clearmatics Technologies Ltd

package cli

import (
	// "strings"
	"fmt"
	"log"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/crypto/sha3"

	"github.com/validation/src/config"
)

type BlockNonce [8]byte

const (
	// BloomByteLength represents the number of bytes used in a header log bloom.
	BloomByteLength = 256

	// BloomBitLength represents the number of bits used in a header log bloom.
	BloomBitLength = 8 * BloomByteLength
)

// Bloom represents a 2048 bit bloom filter.
type Bloom [BloomByteLength]byte

type Block struct {
    Number string
}

type Header struct {
	Hash  common.Hash
	ParentHash  common.Hash
	UncleHash   common.Hash
	Coinbase    common.Address
	Root        common.Hash
	TxHash      common.Hash
	ReceiptHash common.Hash
	Bloom       Bloom
	Difficulty  string
	Number      string
	GasLimit    string
	GasUsed     string
	Time        uint64
	Extra       []byte
	MixDigest   common.Hash
	Nonce       string
}

// type Header struct {
// 		Hash string
//     ParentHash string
//     UncleHash string
//     Coinbase string
//     Root string
//     TxHash string
//     ReceiptHash string
//     Bloom string
//     Difficulty string
//     Number string
//     GasLimit string
//     GasUsed string
//     Time string
//     ExtraData string
//     MixDigest string
//     Nonce string
// }

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

	// Get block N
	shell.AddCmd(&ishell.Cmd{
		Name: "getBlock",
		Help: "Gets specific block of Ethereum instance specified as from",
		Func: func(c *ishell.Context) {
			c.Println("===============================================================")
			if len(c.Args) > 1 {
				c.Println("Too many arguments entered.")
			} else {
				// block, err := strconv.Atoi(c.Args[0])
				// fmt.Println(strconv.FormatInt(c.Args[0], 16))
				// fmt.Println(block,strconv.FormatInt(block, 16))
				// if err != nil {
				// 	log.Fatalf("could not create RPC client: %v", err)
				// }

				c.Println("Get block " + c.Args[0] + " :")

				var blockHeader Header
				err = client.Call(&blockHeader, "eth_getBlockByNumber", c.Args[0], true)
				if err != nil {
					fmt.Println("can't get latest block:", err)
					return
					} else {
						// Print events from the subscription as they arrive.
						// k, _ := strconv.ParseInt(blockHeader.Number, 0, 64)
						// fmt.Printf("latest block: %v\n", k)
						fmt.Printf("%+v\n", blockHeader)
					}
					hasher := sha3.NewKeccak256()
					rlp.Encode(hasher, []interface{}{
						blockHeader.ParentHash,
						blockHeader.UncleHash,
						blockHeader.Coinbase,
						blockHeader.Root,
						blockHeader.TxHash,
						blockHeader.ReceiptHash,
						blockHeader.Bloom,
						blockHeader.Difficulty,
						blockHeader.Number,
						blockHeader.GasLimit,
						blockHeader.GasUsed,
						blockHeader.Time,
						blockHeader.Extra, // Yes, this will panic if extra is too short
						blockHeader.MixDigest,
						blockHeader.Nonce,
					})
					// fmt.Println(hasher)
					var hash common.Hash
					fmt.Printf("%x\n", hasher.Sum(hash[:0]))
					fmt.Printf("%x\n", blockHeader.Hash)
			}

			c.Println("===============================================================")
		},
	})

	// run shell
	shell.Run()
}
