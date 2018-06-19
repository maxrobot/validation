// Copyright (c) 2018 Clearmatics Technologies Ltd

package cli

import (
	"fmt"
  "math/big"
	"strconv"
	"encoding/hex"

  "github.com/ethereum/go-ethereum/rpc"
  "github.com/ethereum/go-ethereum/rlp"
  "github.com/ethereum/go-ethereum/common"
)

type Header struct {
    ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
    UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
    Coinbase    common.Address `json:"miner"            gencodec:"required"`
    Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
    TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
    ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
    Bloom       Bloom          `json:"logsBloom"        gencodec:"required"`
    Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
    Number      *big.Int       `json:"number"           gencodec:"required"`
    GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
    GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
    Time        *big.Int       `json:"timestamp"        gencodec:"required"`
    Extra       string				 `json:"extraData"        gencodec:"required"`
    MixDigest   common.Hash    `json:"mixHash"          gencodec:"required"`
    Nonce       BlockNonce     `json:"nonce"            gencodec:"required"`
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

func EncodeBlock(header Header) (h []byte) {
	// Annoyingly we need to funk around with the extraData field
	encodedExtraData, err := hex.DecodeString(header.Extra[2:])
	if err != nil {
	    panic(err)
	}

  x := []interface{}{
    header.ParentHash,
    header.UncleHash,
    header.Coinbase,
    header.Root,
    header.TxHash,
    header.ReceiptHash,
    header.Bloom,
    header.Difficulty,
    header.Number,
    header.GasLimit,
    header.GasUsed,
    header.Time,
    encodedExtraData,
    header.MixDigest,
    header.Nonce,
  }

	h, _ = rlp.EncodeToBytes(x)

  return h
}
