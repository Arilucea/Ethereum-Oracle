// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.10;

import {ImplementationStorage} from "./Storage/ImplementationStorage.sol";
import {OwnableStorage} from "./Storage/OwnableStorage.sol";

/**
 * @title Proxy contract
 */
contract Proxy {

    constructor(address implementation) payable {
        OwnableStorage.Layout storage l = OwnableStorage.ownableStorage();
        l.owner = (payable(msg.sender));
        _setImplementation(implementation);
    }

    /**
     * @notice set address of implementation contract
     * @param newImplementation address of the new implementation
     */
    function setImplementation(address newImplementation) external onlyOwner {
        _setImplementation(newImplementation);
    }
    
    /**
     * @notice get address of implementation contract
     * @return implementation address
     */
    function getImplementation() external view returns(address implementation) {
        return _getImplementation();
    }

    function _getImplementation() internal view returns(address implementation) {
        return ImplementationStorage.implementation().implementation;
    }

    function _setImplementation(address implementation) internal {
        ImplementationStorage.implementation().implementation = implementation;
    }

    modifier onlyOwner {
        require(msg.sender == OwnableStorage.ownableStorage().owner, "Not the owner of the contract");
        _;
    }

    fallback() external payable {
        address implementation = _getImplementation();

        assembly {
            calldatacopy(0, 0, calldatasize())
            let result := delegatecall(gas(), implementation, 0, calldatasize(), 0, 0)
            returndatacopy(0, 0, returndatasize())
            switch result
            case 0 {
                revert(0, returndatasize())
            }
            default {
                return(0, returndatasize())
            }
        }
    }

    receive() external payable {}

}