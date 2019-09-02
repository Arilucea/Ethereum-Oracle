pragma solidity ^0.5.1;

interface Oracle {
    function makePetition(string calldata _message, string calldata _type) external payable;
}

contract EjemploSolicitud {

    Oracle oracleInt;
    uint public datoInt;
    string public datoString;
    address owner;


    constructor(address _oracle) public {
        oracleInt = Oracle(_oracle);
        owner = msg.sender;

    }

    //Ejecuta la peticion al oraculo
    function callInt(string memory _message, string memory _type) public payable {
        oracleInt.makePetition.value(msg.value)(_message, _type);
    }

    //Respuesta del oraculo, si la respuesta es una cadena
    function __callback(string calldata value) external {
        require(msg.sender == address(0x5Df13792d16D353e07db601560A55Df0ac498fD1), "No oracle");
        datoString = value;
    }

    //Respuesta del oraculo, si la respuesta es un entero
    function __callback(uint value) external {
        require(msg.sender == address(0x5Df13792d16D353e07db601560A55Df0ac498fD1), "No oracle");
        datoInt = value;
    }

    function changeOracle(address _oracle) public {
        require(msg.sender == owner, "Not the owner");
        oracleInt = Oracle(_oracle);
    }

}

//"https://api.coinbase.com/v2/prices/spot?currency=USD(.data.amount)json", "URL/GET"