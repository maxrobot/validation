// Copyright (c) 2018 Clearmatics Technologies Ltd

package cli

import (
	"fmt"
	// "math/big"
	"encoding/hex"
	"strconv"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	// "github.com/ethereum/go-ethereum/common"
)

type Header struct {
	ParentHash  string `json:"parentHash"       gencodec:"required"`
	UncleHash   string `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    string `json:"miner"            gencodec:"required"`
	Root        string `json:"stateRoot"        gencodec:"required"`
	TxHash      string `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash string `json:"receiptsRoot"     gencodec:"required"`
	Bloom       string `json:"logsBloom"        gencodec:"required"`
	Difficulty  string `json:"difficulty"		   gencodec:"required"`
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

		encodedBlock := EncodeBlock(blockHeader)
		fmt.Printf("%+v\n", encodedBlock)
	}
}

func EncodeBlock(header Header) (h []byte) {
	// Annoyingly we need to funk around with the fields ugly for the time being
	encParentHash, _ := hex.DecodeString(header.ParentHash[2:])
	encUncleHash, _ := hex.DecodeString(header.UncleHash[2:])
	encCoinbase, _ := hex.DecodeString(header.Coinbase[2:])
	encRoot, _ := hex.DecodeString(header.Root[2:])
	encTxHash, _ := hex.DecodeString(header.TxHash[2:])
	encReceiptHash, _ := hex.DecodeString(header.ReceiptHash[2:])
	encBloom, _ := hex.DecodeString(header.Bloom[2:])
	encDifficulty, _ := hex.DecodeString(header.Difficulty[2:])
	encNumber, _ := hex.DecodeString(header.Number[2:])
	encGasLimit, _ := hex.DecodeString(header.GasLimit[2:])
	encGasUsed, _ := hex.DecodeString(header.GasUsed[2:])
	encTime, _ := hex.DecodeString(header.Time[2:])
	encExtra, _ := hex.DecodeString(header.Extra[2:])
	encMixDigest, _ := hex.DecodeString(header.MixDigest[2:])
	encNonce, _ := hex.DecodeString(header.Nonce[2:])

	x := []interface{}{
		encParentHash,
		encUncleHash,
		encCoinbase,
		encRoot,
		encTxHash,
		encReceiptHash,
		encBloom,
		encDifficulty,
		encNumber,
		encGasLimit,
		encGasUsed,
		encTime,
		encExtra,
		encMixDigest,
		encNonce,
	}

	h, _ = rlp.EncodeToBytes(x)

	return h
}
