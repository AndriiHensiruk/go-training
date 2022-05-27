// SPDX-License-Identifier: MIT
pragma solidity ^0.8.14;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract TrainingERC20 is ERC20 {

    constructor(uint256 initialSupply) ERC20("Training ERC20", "T20") {
        _mint(msg.sender, initialSupply);
    }
}