// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.10;

library OwnableStorage {
    bytes32 constant OWNABLE_SLOT = keccak256("arilucea.contracts.storage.ownable");

    struct Layout {
        address payable owner;
    }

    function ownableStorage() internal pure returns(Layout storage l) {
        bytes32 position = OWNABLE_SLOT;
        assembly {
            l.slot := position
        }
    }

}