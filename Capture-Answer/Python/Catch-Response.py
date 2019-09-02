#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
@author: javier
"""
import requests, json, os
from web3 import Web3, HTTPProvider, WebsocketProvider
from ContractAbi import abiOracle, abiAnswer

ConfigData = []
with open("../config.json", "r") as file:
    ConfigData = json.load(file)

#Establish connection with the network and the socket, API key for infura is on a enviaroment variable
InfuraAPI = ConfigData["rpcURL"] + os.environ['InfuraOracle']
w3 = Web3(HTTPProvider(InfuraAPI))
w3Socket = Web3(WebsocketProvider(ConfigData["socketURL"]))

#Contract address
oracleAddr = w3.toChecksumAddress(ConfigData["oracleAddr"])
oracleCont = w3.eth.contract(address=oracleAddr, abi=abiOracle()) #Contract definition using address and abi

#Contract definition to capture the events
oracleSocket = w3Socket.eth.contract(address=oracleAddr, abi=abiOracle())

pk = os.environ['PK1'] #The private key is on a enviroment variable

# Split the parameters of the petition from the event and make the GET requests to the API
def proces(dataEvent):
    response = ""
    dBytes = bytes.fromhex(dataEvent[2:]) #Strip the 0x and make the content bytes
    dString = dBytes.replace(b'\x00', b"") #Delete empty values x00
    dHttp = dString[dString.find(b"http"):].decode("utf-8") #Make it string
    
    #Get the URL, this are all the characters before "("
    httpURL, paramsComp = dHttp.split("(")
    #Geth the parameters, everything before ")"
    params, _ = paramsComp.split(")")
    
    #Check if the API response is a json, then make the call
    if (paramsComp.find("json") != -1):
        response = requests.get(httpURL).json()
    paramsList = []
    
    #Process the data from the API petition and get the request information
    #A "." before a value indicates that it is a string
    #A "|" indicates it is a number
    
    paramsSplit = params.split(".") #Split the string by the dots   
    
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
        
    return(value)
#Recover the unique ID from the event to identify the petition
def recoverCaller(dataEvent):
    dataEvent = dataEvent[2:] #Trim the "0x"
    q32bytes = int(len(dataEvent)/64) #Split the chain in sets of 32bytes
    a = []
    for i in range(q32bytes):
        a.append(dataEvent[i*64:(i+1)*64])
    return("0x" + a[0][-40:]) #Generate the unique ID

r = []
def log_loop(event_filter):
    while True:
        r = []
        nonce = w3.eth.getTransactionCount(Web3.toChecksumAddress(ConfigData["senderAddr"]))
        try:
            for event in event_filter.get_new_entries():
                #print(event)
                data = event["data"]
                #print(data)
                
                resp = proces(data)
                addressCaller = recoverCaller(data)
                #print(resp)
                
                #print(oracleCont.call().addSen(idPeticion))
                if (type(resp) == float):
                    resp = int(resp)
                
                contractAnswer = w3.eth.contract(address=w3.toChecksumAddress(addressCaller), abi=abiAnswer())
                
                #Transaction answers
                tx = contractAnswer.functions.__callback(resp).buildTransaction({
                        'gas': 4700000,
                        'gasPrice': w3.toWei("1", "gwei"),
                        'nonce': nonce})
                signed_tx = w3.eth.account.signTransaction(tx, private_key=pk)            
                txHash = w3.eth.sendRawTransaction(signed_tx.rawTransaction)
                r.append(txHash)
                print(r)
                nonce = nonce + 1
                                
                event_filter = oracleSocket.events.Petition.createFilter(fromBlock='latest')
        except Exception as e:
            print("Error for events", e)
            event_filter = oracleSocket.events.Petition.createFilter(fromBlock='latest')



print("Iniciando captura")
event_filter = oracleSocket.events.Petition.createFilter(fromBlock='latest')
log_loop(event_filter)


"https://reqres.in/api/products/3(data.year)json"
'''  
"https://api.coinmarketcap.com/v1/ticker/ethereum(|0.price_usd)json"
"https://api.coinbase.com/v2/prices/spot?currency=USD(.data.amount)json"
'''
