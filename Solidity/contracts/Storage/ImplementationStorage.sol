// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.10;

library ImplementationStorage {
    bytes32 constant IMPLEMENTATION_SLOT = keccak256("arilucea.contracts.storage.implementation");

    struct Layout {
        address implementation;
    }

    function implementation() internal pure returns(Layout storage l) {
        bytes32 position = IMPLEMENTATION_SLOT;
        assembly {
            l.slot := position
        }
    }

}