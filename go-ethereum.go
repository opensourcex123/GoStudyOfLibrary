package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("//./pipe/geth.ipc")
	if err != nil {
		log.Fatal(err)
	}

	addr := common.HexToAddress("0x59F01C809c2262ed2d651c4160d804A4d1e8001e")
	balance, err := client.BlockNumber(context.Background())
	chainId, _ := client.ChainID(context.Background())
	// 获取合约账户的代码
	block, _ := client.CodeAt(context.Background(), addr, nil)
	networkId, _ := client.NetworkID(context.Background())
	// 账户的交易数
	txCount, _ := client.NonceAt(context.Background(), addr, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(balance)
	log.Println(chainId)
	log.Println(block)
	log.Println(networkId)
	log.Println(txCount)
}
