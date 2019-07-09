pragma solidity ^0.5.1;

interface Caller {
    function __callback(string calldata value) external;
    function __callback(uint value) external;
    function __callback(bool value) external;
}

contract OracleEj {
    address payable owner;

    Caller callerInt;

    //Mapea el idUnico con la direccion que envia la peticion
    mapping(bytes => address) Id_Sender;

    //Si true, direcciones autorizas a ejecutar llamadas sin necesidad de pagar
    mapping(address => bool) WhiteList;

    //Petititon cost
    uint256 public price = 0.00 ether;

    event Petition(bytes id, string mType, string message);

    constructor () public {
        owner = msg.sender;
    }

    /**
     * @dev recieves a request an event with the information
     * @param _message information of the request
     * @param _type type of request
     */
    function makePetition(string memory _message, string memory _type) public payable isContract {
        require(WhiteList[msg.sender] == true || msg.value >= price, "No enough value or white list");
        bytes memory id = UniqueID();
        Id_Sender[id] = msg.sender;
        emit Petition(id, _type, _message);
    }

    //--------------------Answers---------------------//
    /**
     * @dev send the answer of the petition to the caller, case string
     * @param _id unique id of the request
     * @param _value answer
     */
    function answer(bytes memory _id, string memory _value) public onlyOwner {
        address answerTo = Id_Sender[_id];
        callerInt = Caller(answerTo);
        callerInt.__callback(_value);
    }

    /**
     * @dev send the answer of the petition to the caller, case uint
     * @param _id unique id of the request
     * @param _value answer
     */
    function answer(bytes memory _id, uint _value) public onlyOwner {
        address answerTo = Id_Sender[_id];
        callerInt = Caller(answerTo);
        callerInt.__callback(_value);
    }

    /**
     * @dev send the answer of the petition to the caller, case boolean
     * @param _id unique id of the request
     * @param _value answer
     */
    function answer(bytes memory _id, bool _value) public onlyOwner {
        address answerTo = Id_Sender[_id];
        callerInt = Caller(answerTo);
        callerInt.__callback(_value);
    }

    /**
     * @dev add or remove an address from the whitelist
     * @param _address address to change
     * @param _state true/false add/remove from the whitelist
     */
    function changeWhiteList(address _address, bool _state) public onlyOwner {
        WhiteList[_address] = _state;
    }

    //--------------------Utils---------------------//

    /*TODO*/
    //This uinque ID only handles one petition per address on each block

    /**
     * @dev make a unique ID for each petition by combianing the address of the caller and the timestamp
     */
    function UniqueID() internal view returns(bytes memory) {
        bytes memory add = bytes(abi.encodePacked(msg.sender));
        bytes memory time = bytes(abi.encodePacked(now));
        string memory abcde = new string(add.length + (time.length-27));
        bytes memory res = bytes(abcde);

        uint i;
        uint t;

        for (i = 0; i<add.length; i++){
            res[t++] = add[i];
        }
        for (i = 27; i<time.length; i++){
            res[t++] = time[i];
        }

        return(res);
    }


    //--------------------Modifiers---------------------//

    /**
     * @dev Returns true if the sender is a contract.
     */
    modifier isContract() {
            uint256 size;
            address sender = msg.sender;
            assembly { size := extcodesize(sender) }
            require(size > 0, "The caller is not a contract");
            _;
    }

    modifier onlyOwner(){
        require(msg.sender == owner, "Not the owner");
        _;
    }


    //--------------------Ether---------------------//

    /**
     * @dev change the price of each the requests
     * @param _newPrice new price in weis
     */
    function changePrize(uint256 _newPrice) public onlyOwner {
        price = _newPrice;
    }

    function getEth() public onlyOwner {
        owner.transfer(address(this).balance);
    }

    //"https://api.coinmarketcap.com/v1/ticker/ethereum/(|0.price_usd)json", "URL/JSON"

}