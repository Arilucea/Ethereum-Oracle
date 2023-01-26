// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.10;

import {OracleStorage} from "./Storage/OracleStorage.sol";
import {OwnableStorage} from "./Storage/OwnableStorage.sol";

contract Oracle {

    event Petition(address caller, string message, string mType);

    /**
     * @dev recieves a request an event with the information
     * @param _message information of the request
     * @param _type type of request
     */
    function makePetition(string memory _message, string memory _type) external payable isContract {
        OracleStorage.Layout storage l = OracleStorage.oracleStorage();
        require(l.WhiteList[msg.sender] == true || msg.value >= l.price, "No enough value or whitelist");
        emit Petition(msg.sender, _message, _type);
    }

    /**
     * @dev add or remove an address from the whitelist
     * @param _address address to change
     * @param _state true/false add/remove from the whitelist
     */
    function changeWhiteList(address _address, bool _state) external onlyOwner {
        OracleStorage.Layout storage l = OracleStorage.oracleStorage();
        l.WhiteList[_address] = _state;
    }

    //--------------------Utils---------------------//

    event OwnershipChanged(address previousOwner, address newOwner);

    function owner() external view returns(address) {
        return _owner();
    }

    function _owner() internal view returns(address) {
        return OwnableStorage.ownableStorage().owner;
    }

    function transferOwnership(address _newOwner) external onlyOwner {
        OwnableStorage.Layout storage l = OwnableStorage.ownableStorage();
        address previousOwner = l.owner;
        l.owner = payable(_newOwner);
        emit OwnershipChanged(previousOwner, _newOwner);
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
        require(msg.sender == _owner(), "Not the owner");
        _;
    }


    //--------------------Ether---------------------//

    /**
     * @dev change the price of each the requests
     * @param _newPrice new price in weis
     */
    function changePrize(uint256 _newPrice) external onlyOwner {
        OracleStorage.Layout storage l = OracleStorage.oracleStorage();
        l.price = _newPrice;
    }

    function getEth() external onlyOwner {
        (bool success, ) = (_owner()).call{value: address(this).balance}("");
        require(success == true, "Currency transfer fail");
    }

    //"https://api.coinmarketcap.com/v1/ticker/ethereum/(|0.price_usd)json", "URL/JSON"

}