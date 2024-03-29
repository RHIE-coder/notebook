//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MyContract {

    bytes16 private constant _SYMBOLS = "0123456789abcdef";

    /**
     * @dev Return the log in base 10, rounded down, of a positive value.
     * Returns 0 if given 0.
     */
    function log10(uint256 value) internal pure returns (uint256) {
        uint256 result = 0;
        unchecked {
            if (value >= 10 ** 64) {
                value /= 10 ** 64;
                result += 64;
            }
            if (value >= 10 ** 32) {
                value /= 10 ** 32;
                result += 32;
            }
            if (value >= 10 ** 16) {
                value /= 10 ** 16;
                result += 16;
            }
            if (value >= 10 ** 8) {
                value /= 10 ** 8;
                result += 8;
            }
            if (value >= 10 ** 4) {
                value /= 10 ** 4;
                result += 4;
            }
            if (value >= 10 ** 2) {
                value /= 10 ** 2;
                result += 2;
            }
            if (value >= 10 ** 1) {
                result += 1;
            }
        }
        return result;
    }

    /**
     * @dev Converts a `uint256` to its ASCII `string` decimal representation.
     */
    function toString(uint256 value) internal pure returns (string memory) {
        unchecked {
            uint256 length = log10(value) + 1;
            string memory buffer = new string(length);
            uint256 ptr;
            /// @solidity memory-safe-assembly
            assembly {
                ptr := add(buffer, add(32, length))
            }
            while (true) {
                ptr--;
                /// @solidity memory-safe-assembly
                assembly {
                    mstore8(ptr, byte(mod(value, 10), _SYMBOLS))
                }
                value /= 10;
                if (value == 0) break;
            }
            return buffer;
        }
    }

    function doSomething(uint256 myFunctionParam) public pure returns(string memory){
        return string.concat("hello there: ", toString(myFunctionParam));
    }
}

contract MyOtherContract {
    function doTheThing(address myContractAddress, uint256 myFunctionParam) public returns(string memory){
        (bool success, bytes memory result) = myContractAddress.call(
            abi.encodeWithSignature(
                "doSomething(uint256)",
                myFunctionParam
            )
        );

        if(!success){
            revert("no no no");
        }

        string memory message = abi.decode(result, (string));
        return message;
    }
}


contract Relayer {

    address resource;
    event Response(bool success, bytes data);

    constructor(address _targetAddress) {
        resource = _targetAddress;
    }

    function relay(address _sender, uint _value) public {
        (bool isSuccess, bytes memory data) = resource.call{}(abi.encodeWithSignature("callMe(uint)", _value));

        emit Response(isSuccess, data);
    }
}


