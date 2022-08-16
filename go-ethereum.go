package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("//./pipe/geth.ipc")
	if err != nil {
		log.Fatal(err)
	}

	//addr := common.HexToAddress("0xc365c3315cf926351ccaf13fa7d19c8c4058c8e1")
	balance, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println(balance)
}
