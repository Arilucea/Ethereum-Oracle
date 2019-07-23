package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	readConfig("../config.json", &config)

	senderAddr := common.HexToAddress(config.SenderAddr)
	oracleAddr := common.HexToAddress(config.OracleAddr)
	addressesList := [2]common.Address{senderAddr, oracleAddr}

	ctx, client, socket, nonce, oracleCont := networkInit(addressesList)

	response(oracleCont, nonce)

	captureEvents(ctx, oracleAddr, socket)

	balance, _ := client.BalanceAt(ctx, senderAddr, nil)

	fmt.Println(balance)

}

//--------------------------------- Read configuration parameters ---------------------------------//
func readConfig(fileName string, config *Configuration) {
	file, err := os.Open(fileName)
	fatalError("Error opening file", err)

	json.NewDecoder(file).Decode(&config)
}

//--------------------------------- System init ---------------------------------//
func networkInit(adresses [2]common.Address) (context.Context, *ethclient.Client, *ethclient.Client, uint64, *contract.OracleGo) {
	ctx := context.Background()

	rpcURL := config.RPCURL + os.Getenv("InfuraOracle")
	client, err := ethclient.Dial(rpcURL)
	fatalError("Error connecting to the network", err)

	socket, err := ethclient.Dial(config.SocketURL)
	fatalError("Error connecting to the socket", err)

	nonce, err := client.NonceAt(context.Background(), adresses[0], nil)

	oracleCont, err := contract.NewOracleGo(adresses[1], client)

	return ctx, client, socket, nonce, oracleCont
}

//--------------------------------- Event elements ---------------------------------//
type eventFormat struct {
	Id      []byte
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

	contractAbi, err := abi.JSON(strings.NewReader(contract.OracleGoABI))
	fatalError("Abi problem", err)

	fmt.Println("System initiated")

	for {
		select {
		case err := <-SubEvent.Err():
			log.Println(err)
		case vLog := <-logs:
			err := contractAbi.Unpack(&event, "Petition", vLog.Data)

			printError("Error parsing event", err)

			id := hex.EncodeToString(event.Id)
			id = "0x" + id

			parsePetition(string(event.Message[:]))

			fmt.Println(id)
			fmt.Println(string(event.MType[:]))
			fmt.Println(string(event.Message[:]))
		}
	}

}

func parsePetition(url string) {

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

	fmt.Println(params)
}

func processParams(params string) []string {

	paramsSplit := strings.Split(params, ".")

	/*for _, item := range paramsSplit {
		if strings.Index(item, "|") != -1 {
			fmt.Println("---------")
		}
	}*/

	//Check if the firs element of the parameters in the petition is a "."
	if params[0] == byte(46) {
		paramsSplit = paramsSplit[1:]
	}

	return paramsSplit

	/*paramsSplit = params.split(".") #Split the string by the dots
	  #Create an array of numbers and string based on the characters "|", "."
	  for v in paramsSplit:
	      if v.find("|") != -1:
	          r = v.split("|")
	          for i in range(len(r)):
	              if (i==0  and v.find("|")!=0):
	                  if (r[i]!=""):
	                     paramsList.append(r[i])
	              elif (i != 0 or v.find("|")==0):
	                  if (r[i]!=""):
	                      paramsList.append(int(r[i]))
	      else:
	          if (v!=""):
	              paramsList.append(v)

	  #Return the wanted value from the GET response based on the parameters of the event
	  #response - json recover from the API
	  #paramList - list of parameters from the event, to get the desired field of the API response
	  value = response
	  for item in paramsList:
	      value = value[item]

	  return(value)*/

}

func response(*contract.OracleGo, uint64) {

	/*privateKey, err := crypto.HexToECDSA(os.Getenv("PK1"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}*/
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
