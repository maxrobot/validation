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
	ParentHash  string `json:"parentHash"`
	UncleHash   string `json:"sha3Uncles"`
	Coinbase    string `json:"miner"`
	Root        string `json:"stateRoot"`
	TxHash      string `json:"transactionsRoot"`
	ReceiptHash string `json:"receiptsRoot"`
	Bloom       string `json:"logsBloom"`
	Difficulty  string `json:"difficulty"`
	Number      string `json:"number"`
	GasLimit    string `json:"gasLimit"`
	GasUsed     string `json:"gasUsed"`
	Time        string `json:"timestamp"`
	Extra       string `json:"extraData"`
	MixDigest   string `json:"mixHash"`
	Nonce       string `json:"nonce"`
}

func latestBlock(client *rpc.Client) {
	var lastBlock Block
	err := client.Call(&lastBlock, "eth_getBlockByNumber", "latest", true)
	if err != nil {
		fmt.Println("can't get latest block:", err)
		return
	}
	// Print events from the subscription as they arrive.
	k, _ := strconv.ParseInt(lastBlock.Number, 0, 64)
	fmt.Printf("latest block: %v\n", k)
}

func getBlock(client *rpc.Client, block string) {
	var blockHeader Header
	err := client.Call(&blockHeader, "eth_getBlockByNumber", block, true)
	if err != nil {
		fmt.Println("can't get requested block:", err)
		return
	}
	fmt.Printf("%+v\n", blockHeader)
}

func rlpEncodeBlock(client *rpc.Client, block string) {
	var blockHeader Header
	err := client.Call(&blockHeader, "eth_getBlockByNumber", block, true)
	if err != nil {
		fmt.Println("can't get requested block:", err)
		return
	}
	blockInterface := GenerateInterface(blockHeader)
	encodedBlock := EncodeBlock(blockInterface)
	fmt.Printf("%+x\n", encodedBlock)
}

func calculateRlpEncoding(client *rpc.Client, block string) {
	var blockHeader Header
	err := client.Call(&blockHeader, "eth_getBlockByNumber", block, true)
	if err != nil {
		fmt.Println("can't get requested block:", err)
		return
	}
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

// Creates an interface for a block
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

// Encodes a block
func EncodeBlock(blockInterface interface{}) (h []byte) {
	h, _ = rlp.EncodeToBytes(blockInterface)

	return h
}
