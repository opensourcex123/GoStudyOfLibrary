package main

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
)

func main() {
	//client, err := ethclient.Dial("//./pipe/geth.ipc")
	client, err := ethclient.DialContext(context.Background(), "http://192.168.0.106:8545")
	if err != nil {
		log.Fatal(err)
	}

	contractAddr := common.HexToAddress("0x35D1A29E20AdEF025e72909938D3d803B804f28D")
	addr := common.HexToAddress("0x3f53893bdd70220f167f920a519de830bf2a3a84")
	block, _ := client.BlockNumber(context.Background())
	chainId, _ := client.ChainID(context.Background())
	// 获取合约账户的代码
	//code, _ := client.CodeAt(context.Background(), contractAddr, nil)
	networkId, _ := client.NetworkID(context.Background())
	// 账户的交易数
	txCount, _ := client.NonceAt(context.Background(), addr, nil)

	pendingCount, _ := client.PendingTransactionCount(context.Background())

	log.Println(block)
	log.Println(chainId)
	//log.Println(string(code))
	log.Println(networkId)
	log.Println(txCount)
	log.Println("等待的交易数：", pendingCount)

	//tipCap, _ := client.SuggestGasTipCap(context.Background())
	//feeCap, _ := client.SuggestGasPrice(context.Background())

	json := `[
	{
		"inputs": [],
		"name": "retrieve",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "num",
				"type": "uint256"
			}
		],
		"name": "store",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
	abi, _ := abi.JSON(strings.NewReader(json))
	data, _ := abi.Pack("retrieve")
	msg := ethereum.CallMsg{
		To: &contractAddr,
		//Gas:       uint64(21064),
		//GasTipCap: tipCap,
		//GasFeeCap: feeCap,
		Data: data,
	}
	resp, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		log.Fatal("err:", err)
	}
	log.Println(resp)
}
