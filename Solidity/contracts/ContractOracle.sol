// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.10;

contract Oracle {
    address payable owner;

    //Addresses authorized to send petition without paying the fee
    mapping(address => bool) WhiteList;

    //Petititon cost
    uint256 public price = 0.00 ether;

    event Petition(address caller, string mType, string message);

    constructor () {
        owner = payable(msg.sender);
    }

    /**
     * @dev recieves a request an event with the information
     * @param _message information of the request
     * @param _type type of request
     */
    function makePetition(string memory _message, string memory _type) external payable isContract {
        require(WhiteList[msg.sender] == true || msg.value >= price, "No enough value or whitelist");
        emit Petition(msg.sender, _type, _message);
    }

    /**
     * @dev add or remove an address from the whitelist
     * @param _address address to change
     * @param _state true/false add/remove from the whitelist
     */
    function changeWhiteList(address _address, bool _state) external onlyOwner {
        WhiteList[_address] = _state;
    }

    //--------------------Utils---------------------//


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
    function changePrize(uint256 _newPrice) external onlyOwner {
        price = _newPrice;
    }

    function getEth() external onlyOwner {
        owner.transfer(address(this).balance);
    }

    //"https://api.coinmarketcap.com/v1/ticker/ethereum/(|0.price_usd)json", "URL/JSON"

}