# Simple Ethereum-Oracle
Given an URL to the contract using the "makePetition" function, it will answer with the requested information from the API and store it on the caller contract using the "__callback" function. **Currently the system only supports GET requests.**

## Configuration
In the config.json file change the the parameters needed:
- **rpcURL**: endpoint of the network (in the example using Infura).
- **socketURL**: endpoint of the socket to capture the event (in the example using Infura).
- **oracleAddr**: address of the oracle contract.
- **senderAddr**: address used to send the answers.

The private key of the address used to send the answers and API key of infura, both are used in the scripts are keep as enviroment variables.

## How to use it
1. Deploy ```ContractOracle```
2. Change the parameters on ```config.json```
3. Start the script ```Catch-Response```
4. Caller contract must implement the "__callback" function.

### Petition examples
First must be indicate the the url, after it there must be an "(" containing the chain of elements to get the one we want from the response if the field is a string a "." should be before the value, if the field is a number there should be a "|". After the last element the brackest have to be closed and inicate if the response is a json.

> "https://api.coinbase.com/v2/prices/spot?currency=USD(.data.amount)json", "URL/GET"

We want to get the Bitcoin price in USD an store it in our contract. We call this API
```https://api.coinbase.com/v2/prices/spot?currency=USD``` json as response.
```
{
    'data': 
            {
                'base': 'BTC', 
                'currency': 'USD', 
                'amount': '11221.685'
            }
}
```
To make the petition we have to indicate the field we want from the json ("amount"). The petition to the oracle should be the following.
>"https://api.coinbase.com/v2/prices/spot?currency=USD(.data.amount)json", "URL/GET"

The fields we want to filter are "data" and inside it "amount" both are strings so both are precided by a "."

The "__callback" function of the caller will recieve a '11221.685' in the form of a string.
