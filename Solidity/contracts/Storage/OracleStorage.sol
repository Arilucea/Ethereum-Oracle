// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.10;

library OracleStorage {
    bytes32 constant ORACLE_SLOT = keccak256("arilucea.contracts.storage.oracle");

    struct Layout {
        mapping(address => bool) WhiteList;

        //Petititon cost
        uint256 price;//  = 0.00 ether;
    }

    function oracleStorage() internal pure returns(Layout storage l) {
        bytes32 position = ORACLE_SLOT;
        assembly {
            l.slot := position
        }
    }
}