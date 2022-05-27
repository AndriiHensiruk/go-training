// SPDX-License-Identifier: MIT
pragma solidity ^0.8.14;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";


// ====================================== Start main contract ======================================

contract TrainingNFT1155 is ERC1155, Ownable {

    using Address for address;
    using Counters for Counters.Counter;

    Counters.Counter private _tokenIdTracker;


    constructor(string memory uri_) ERC1155(uri_) {
        // start with 1
        _tokenIdTracker.increment();
    }

    function getCurrentTokenId() external view returns (uint256) {
        return _tokenIdTracker.current();
    }


    /**
 * @dev Creates `amount` new tokens for `to`, of token type `id`.
     *
     * See {ERC1155-_mint}.
     *
     * Requirements:
     *
     * - the caller must have the `MINTER_ROLE`.
     */

    function mint(
        address to,
        uint256 id,
        uint256 amount,
        bytes memory data
    ) public onlyOwner {
        _mint(to, id, amount, data);
    }

    /**
     * @dev See {IERC1155-safeTransferFrom}.
     */
    function safeTransferFrom(
        address from,
        address to,
        uint256 id,
        uint256 amount,
        bytes memory data
    ) public override {
        require(
            from == _msgSender() || isApprovedForAll(from, _msgSender()),
            "ERC1155: caller is not owner nor approved"
        );
        _safeTransferFrom(from, to, id, amount, data);
    }


    function owner() public view override (Ownable) returns (address) {
        return super.owner();
    }
}