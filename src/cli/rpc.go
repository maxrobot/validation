// Copyright (c) 2018 Clearmatics Technologies Ltd

package cli

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
)

type Header struct {
	ParentHash  string `json:"parentHash"       gencodec:"required"`
	UncleHash   string `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    string `json:"miner"            gencodec:"required"`
	Root        string `json:"stateRoot"        gencodec:"required"`
	TxHash      string `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash string `json:"receiptsRoot"     gencodec:"required"`
	Bloom       string `json:"logsBloom"        gencodec:"required"`
	Difficulty  string `json:"difficulty"		  	gencodec:"required"`
	Number      string `json:"number"           gencodec:"required"`
	GasLimit    string `json:"gasLimit"         gencodec:"required"`
	GasUsed     string `json:"gasUsed"          gencodec:"required"`
	Time        string `json:"timestamp"        gencodec:"required"`
	Extra       string `json:"extraData"        gencodec:"required"`
	MixDigest   string `json:"mixHash"          gencodec:"required"`
	Nonce       string `json:"nonce"            gencodec:"required"`
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
	err := client.Call(&blockHeader, "eth_getBlockByNumber", block, true)
	if err != nil {
		fmt.Println("can't get requested block:", err)
		return
	} else {
		fmt.Printf("%+v\n", blockHeader)
	}
}

func rlpEncodeBlock(client *rpc.Client, block string) {
	var blockHeader Header
	err := client.Call(&blockHeader, "eth_getBlockByNumber", block, true)
	if err != nil {
		fmt.Println("can't get requested block:", err)
		return
	} else {
		blockInterface := GenerateInterface(blockHeader)
		encodedBlock := EncodeBlock(blockInterface)
		fmt.Printf("%+x\n", encodedBlock)
	}
}

func calculateRlpEncoding(client *rpc.Client, block string) {
	var blockHeader Header
	err := client.Call(&blockHeader, "eth_getBlockByNumber", block, true)
	if err != nil {
		fmt.Println("can't get requested block:", err)
		return
	} else {
		// Generate an interface to encode the standard block header
		blockInterface := GenerateInterface(blockHeader)
		encodedBlock := EncodeBlock(blockInterface)
		fmt.Printf("%+x\n", encodedBlock)

		// Generate an interface to encode the blockheader without the signature in the extraData
		blockHeader.Extra = blockHeader.Extra[:len(blockHeader.Extra)-130]
		blockInterface = GenerateInterface(blockHeader)
		encodedBlock = EncodeBlock(blockInterface)
		fmt.Printf("\n%+x\n", encodedBlock[1:3])

		// Generate an interface to encode the blockheader without the signature in the extraData
		encExtra, _ := hex.DecodeString(blockHeader.Extra[2:])
		encodedBlock = EncodeBlock(encExtra)
		fmt.Printf("\n%+x\n", encodedBlock[0:1])
	}
}

func GenerateInterface(blockHeader Header) (rest interface{}) {
	blockInterface := []interface{}{}
	s := reflect.ValueOf(&blockHeader).Elem()

	// Append items into the interface
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i).String()
		element, _ := hex.DecodeString(f[2:])
		blockInterface = append(blockInterface, element)
	}

	return blockInterface
}

func EncodeBlock(blockInterface interface{}) (h []byte) {
	// Encode the block
	h, _ = rlp.EncodeToBytes(blockInterface)

	return h
}
