package main

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/foundVanting/opensea-stream-go/entity"
	"github.com/foundVanting/opensea-stream-go/opensea"
	"github.com/foundVanting/opensea-stream-go/types"
	"github.com/mitchellh/mapstructure"
	"github.com/nshafer/phx"
	"github.com/stretchr/testify/require"
	"math/big"
)

func main() {
	client := opensea.NewStreamClient(types.TESTNET, "", phx.LogInfo, func(err error) {
		fmt.Println("NewStreamClient err:", err)
	})
	client.Connect()

	go client.OnItemListed("untitled-collection-14425637", func(response any) {
		var itemListedEvent entity.ItemListedEvent
		err := mapstructure.Decode(response, &itemListedEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemListedEvent)
	})
	go client.OnItemCancelled("untitled-collection-14425637", func(response any) {
		var itemListedEvent entity.ItemListedEvent
		err := mapstructure.Decode(response, &itemListedEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemListedEvent)
	})
	//select {}

	result := encodePacked(
		common.HexToAddress(to).Bytes(),
		encodeUint256Array(ids),
		encodeUint256Array(amounts),
		ticket,
		encodeUint256(big.NewInt(timestamp)),
	)
	hash := crypto.Keccak256(result)
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hash), hash)
	hash = crypto.Keccak256([]byte(msg))
	signature, err := crypto.Sign(hash, privateKey)
	require.NoError(t, err)
	signature[64] += 27
	t.Logf("Signature: %s", hexutil.Encode(signature))
}

func encodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}
