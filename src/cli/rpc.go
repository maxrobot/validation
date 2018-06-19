// Copyright (c) 2018 Clearmatics Technologies Ltd

package cli

import (
	"fmt"
	// "log"
	"strconv"

	// "github.com/validation/src/config"
  "github.com/ethereum/go-ethereum/rpc"
)

type Header struct {
    ParentHash string
    UncleHash string
    Coinbase string
    Root string
    TxHash string
    ReceiptHash string
    Bloom string
    Difficulty string
    Number string
    GasLimit string
    GasUsed string
    Time string
    ExtraData string
    MixDigest string
    Nonce string
}

func latestBlock(client *rpc.Client) {
  var lastBlock Block
  err := client.Call(&lastBlock, "eth_getBlockByNumber", "latest", true)
  if err != nil {
      fmt.Println("can't get latest block:", err)
      return
  } else {
    // Print events from the subscription as they arrive.
    k, _ := strconv.ParseInt(lastBlock.Number, 0, 64)
    fmt.Printf("latest block: %v\n", k)
  }
}

func getBlock(client *rpc.Client, block string) {
  var blockHeader Header
  fmt.Printf(block)
  err := client.Call(&blockHeader, "eth_getBlockByNumber", block, true)
  if err != nil {
  	fmt.Println("can't get requested block:", err)
  	return
  } else {
  	fmt.Printf("%+v\n", blockHeader)
  }
}
