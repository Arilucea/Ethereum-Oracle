#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
@author: javier
"""

def abiOracle():
    abi = [
	{
		"constant": "false",
		"inputs": [
			{
				"name": "_newPrice",
				"type": "uint256"
			}
		],
		"name": "changePrize",
		"outputs": [],
		"payable": "false",
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": "false",
		"inputs": [
			{
				"name": "_address",
				"type": "address"
			},
			{
				"name": "_state",
				"type": "bool"
			}
		],
		"name": "changeWhiteList",
		"outputs": [],
		"payable": "false",
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": "false",
		"inputs": [
			{
				"name": "_message",
				"type": "string"
			},
			{
				"name": "_type",
				"type": "string"
			}
		],
		"name": "makePetition",
		"outputs": [],
		"payable": "true",
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"constant": "true",
		"inputs": [],
		"name": "price",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": "false",
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": "false",
		"inputs": [],
		"name": "getEth",
		"outputs": [],
		"payable": "false",
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"payable": "false",
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": "false",
		"inputs": [
			{
				"indexed": "false",
				"name": "caller",
				"type": "address"
			},
			{
				"indexed": "false",
				"name": "mType",
				"type": "string"
			},
			{
				"indexed": "false",
				"name": "message",
				"type": "string"
			}
		],
		"name": "Petition",
		"type": "event"
	}
]
    return abi




def abiAnswer():
    abi = [
	{
		"constant": "false",
		"inputs": [
			{
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "__callback",
		"outputs": [],
		"payable": "false",
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": "false",
		"inputs": [
			{
				"name": "value",
				"type": "bool"
			}
		],
		"name": "__callback",
		"outputs": [],
		"payable": "false",
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": "false",
		"inputs": [
			{
				"name": "value",
				"type": "string"
			}
		],
		"name": "__callback",
		"outputs": [],
		"payable": "false",
		"stateMutability": "nonpayable",
		"type": "function"
	}
]
    return abi