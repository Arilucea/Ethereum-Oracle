package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	contractCaller "./CallerGo"
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
	readConfig("../config.json", &config)

	senderAddr := common.HexToAddress(config.SenderAddr)
	oracleAddr := common.HexToAddress(config.OracleAddr)
	addressesList := [2]common.Address{senderAddr, oracleAddr}

	networkInit(addressesList)

	balance, _ := client.BalanceAt(ctx, senderAddr, nil)
	fmt.Println(balance)

	captureEvents(ctx, oracleAddr, socket)

}

//--------------------------------- Read configuration parameters ---------------------------------//
func readConfig(fileName string, config *Configuration) {
	file, err := os.Open(fileName)
	fatalError("Error opening file", err)

	json.NewDecoder(file).Decode(&config)
}

//Network values
var ctx context.Context
var client *ethclient.Client
var socket *ethclient.Client
var nonce uint64
var oracleCont *contract.Oracle

var err error

//--------------------------------- System init ---------------------------------//
func networkInit(adresses [2]common.Address) {
	ctx = context.Background()

	rpcURL := config.RPCURL + os.Getenv("InfuraOracle")
	client, err = ethclient.Dial(rpcURL)
	fatalError("Error connecting to the network", err)

	socket, err = ethclient.Dial(config.SocketURL)
	fatalError("Error connecting to the socket", err)

	nonce, err = client.NonceAt(context.Background(), adresses[0], nil)
	fatalError("Error reading nonce", err)

	oracleCont, err = contract.NewOracle(adresses[1], client)
	fatalError("Error initiating the contract", err)

}

//--------------------------------- Event elements ---------------------------------//
type eventFormat struct {
	Caller  common.Address
	MType   string
	Message string
}

var event eventFormat

func captureEvents(ctx context.Context, oracleAddr common.Address, socket *ethclient.Client) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{oracleAddr},
	}

	logs := make(chan types.Log)
	SubEvent, err := socket.SubscribeFilterLogs(ctx, query, logs)

	contractAbi, err := abi.JSON(strings.NewReader(contract.OracleABI))
	fatalError("Abi problem", err)

	fmt.Println("System initiated")

	for {
		select {
		case err := <-SubEvent.Err():
			log.Println(err)
		case vLog := <-logs:
			err := contractAbi.Unpack(&event, "Petition", vLog.Data)
			printError("Error parsing event", err)

			answerVal := parsePetition(string(event.Message[:]))
			//v, _ := strconv.Atoi(answerVal)
			//v2 := big.NewInt(int64(v))
			sendResponse(client, oracleCont, nonce, event.Caller, answerVal, config)
		}
	}
}

func parsePetition(url string) string {

	//-------------------- Event data process ----------------------------//
	tmpSplit := strings.Split(url, "(")
	httpURL := tmpSplit[0]
	paramsComp := tmpSplit[1]
	tmpSplit = strings.Split(paramsComp, ")")
	params := tmpSplit[0]

	//-------------------- Get petition ----------------------------//
	resp, _ := http.Get(httpURL)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	raw := map[string]json.RawMessage{}    //Mapping to store the final result
	rawTmp := map[string]json.RawMessage{} //Mapping to use on during the for process

	json.Unmarshal(body, &raw) //Asign the value of the get response to raw mapping

	lista := processParams(params) //Process the params obtained from the event, it returns an array of strings

	var value []byte //Store the result and the steps to reach it

	for _, ele := range lista {
		rawTmp = map[string]json.RawMessage{}
		json.Unmarshal(raw[ele], &rawTmp)
		value, _ = json.Marshal(raw[ele])
		raw = rawTmp
	}

	// , -> 44
	// [ -> 91

	//Remove " at the begining and the end
	if value[0] == 34 && value[len(value)-1] == 34 {
		value = value[1 : len(value)-1]
	}

	fmt.Println("Acabado", string(value))
	return (string(value))
}

func processParams(params string) []string {

	paramsSplit := strings.Split(params, ".")

	//Check if the firs element of the parameters in the petition is a "."
	if params[0] == byte(46) {
		paramsSplit = paramsSplit[1:]
	}

	return paramsSplit
}

func sendResponse(client *ethclient.Client, contract *contract.Oracle, nonce uint64, callerAddress common.Address, value interface{}, config Configuration) {

	privateKey, err := crypto.HexToECDSA(os.Getenv("PK1"))
	printError("Error with reading private key", err)

	nonce, _ = client.PendingNonceAt(context.Background(), common.HexToAddress(config.SenderAddr))

	gasPrice, err := client.SuggestGasPrice(ctx)
	printError("Error gas price", err)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	contractCaller.SetAbi(value)
	cCaller, _ := contractCaller.NewOracle(callerAddress, client)

	//Answer0 is the function to send strings, answer for bools and answer1 for ints
	tx, err := cCaller.Callback(auth, value)
	printError("Error response transaction", err)

	//nonce++

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
}

func fatalError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func printError(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
	}
}
