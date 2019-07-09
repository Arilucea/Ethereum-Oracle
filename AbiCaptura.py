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
		"payable": "false",
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": "false",
		"inputs": [
			{
				"name": "_id",
				"type": "bytes"
			},
			{
				"name": "_value",
				"type": "string"
			}
		],
		"name": "answer",
		"outputs": [],
		"payable": "false",
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": "false",
		"inputs": [
			{
				"name": "_id",
				"type": "bytes"
			},
			{
				"name": "_value",
				"type": "uint256"
			}
		],
		"name": "answer",
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
				"name": "id",
				"type": "bytes"
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