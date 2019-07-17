package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	contract "./OracleGo"
)

//Configuration json
type Configuration struct {
	RPCURL     string `json:"rpcURL"`
	SocketURL  string `json:"socketURL"`
	OracleAddr string `json:"oracleAddr"`
	SenderAddr string `json:"senderAddr"`
}

var config Configuration

func main() {
	setConfig("../config.json", &config)

	ctx := context.Background()
	senderAddr := common.HexToAddress(config.SenderAddr)

	rpcURL := config.RPCURL + os.Getenv("InfuraOracle")
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal("Error connecting to the network", err)
	}

	socket, err := ethclient.Dial(config.SocketURL)
	if err != nil {
		log.Fatal("Error connecting to the socket", err)
	}

	oracleAddr := common.HexToAddress(config.OracleAddr)
	//oracleCont, err := contract.NewOracleGo(oracleAddr, client)

	if err != nil {
		log.Fatal("Error with the contract", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{oracleAddr},
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(contract.OracleGoABI)))
	if err != nil {
		log.Fatal(err)
	}

	logs := make(chan types.Log)
	sub, err := socket.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:

			event := struct {
				Id      []byte
				MType   string
				Message string
			}{}

			err := contractAbi.Unpack(&event, "Petition", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			id := hex.EncodeToString(event.Id)
			id = "0x" + id

			fmt.Println(id)
			fmt.Println(string(event.MType[:]))
			fmt.Println(string(event.Message[:]))
		}
	}

	//value, err := oracleCont.

	balance, err := client.BalanceAt(ctx, senderAddr, nil)

	fmt.Println(balance)

}

//Read configuration parameters
func setConfig(fileName string, config *Configuration) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file", err)
	}

	json.NewDecoder(file).Decode(&config)
}
