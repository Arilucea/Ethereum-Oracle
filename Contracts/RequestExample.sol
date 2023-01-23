// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.10;

interface Oracle {
    function makePetition(string calldata _message, string calldata _type) external payable;
}

contract PetitionExample {

    Oracle oracleInt;
    uint public datoInt;
    string public datoString;
    address owner;

    constructor(address _oracle) {
        oracleInt = Oracle(_oracle);
        owner = msg.sender;

    }

    //Trigger oracle petition
    function callInt(string memory _message, string memory _type) external payable {
        oracleInt.makePetition{value: msg.value}(_message, _type);
    }

    //Oracle answer if the value is a string
    function __callback(string calldata value) external {
        require(msg.sender == address(0x5Df13792d16D353e07db601560A55Df0ac498fD1), "Answer is not from the oracle");
        datoString = value;
    }

    //Oracle answer if the value is a uint
    function __callback(uint value) external {
        require(msg.sender == address(0x5Df13792d16D353e07db601560A55Df0ac498fD1), "Answer is not from the oracle");
        datoInt = value;
    }

    function changeOracle(address _oracle) public {
        require(msg.sender == owner, "Not the owner");
        oracleInt = Oracle(_oracle);
    }

}

//"https://api.coinbase.com/v2/prices/spot?currency=USD(.data.amount)json", "URL/GET"