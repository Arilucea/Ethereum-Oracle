pragma solidity ^0.5.1;

interface Oracle {
    function makePetition(string calldata _message, string calldata _type) external payable;
}

contract RequestExample {

    Oracle oracleInt;
    uint public datoInt;
    string public datoString;
    address owner;


    constructor(address _oracle) public {
        oracleInt = Oracle(_oracle);
        owner = msg.sender;
    }

    /**
     * @dev make the request to the oracle
     * @param _message information of the request
     * @param _type type of request
     */
    function callInt(string memory _message, string memory _type) public payable {
        oracleInt.makePetition.value(msg.value)(_message, _type);
    }

    /**
     * @dev recieves the answer from the oracle
     */
    function __callback(string calldata value) external onlyOracle {
        datoString = value;
    }

    /**
     * @dev recieves the answer from the oracle
     */
    function __callback(uint value) external onlyOracle {
        datoInt = value;
    }
    /**
     * @dev changes the oracle from the contract
     */
    function changeOracle(address _oracle) public {
        require(msg.sender == owner, "Not the owner");
        oracleInt = Oracle(_oracle);
    }

    modifier onlyOracle() {
        require(msg.sender == address(oracleInt), "Not the oracle");
        _;
    }

}