// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

// 0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2
// 0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db
contract PointDev is ERC20 {
    constructor() ERC20("PointDev","PTD") {
        _mint(msg.sender,10_000 * 1_000_000_000_000_000_000);
    }
}
// 소피: 0xafb6dC45682507d8D6C6C493968FF0e57413560d