// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title TinyOracle
 * @dev This contract serves as a simple oracle that stores a boolean value
 * which can be set only once by trusted wallets.
 */
contract TinyOracle {
    // Boolean value that represents the state/answer
    // Initially undefined (not set)
    bool private result;
    bool private resultIsSet = false;
    
    // Mapping of trusted wallets that are allowed to set the result
    mapping(address => bool) public trustedWallets;
    
    // Event emitted when result is set
    event ResultSet(bool indexed value, address setter);
    
    /**
     * @dev Constructor that initializes trusted wallets
     * @param _trustedWallets Comma-separated list of addresses that can set the result
     */
    constructor(address[] memory _trustedWallets) {
        for (uint i = 0; i < _trustedWallets.length; i++) {
            trustedWallets[_trustedWallets[i]] = true;
        }
    }
    
    /**
     * @dev Sets the result. Can only be called once by a trusted wallet.
     * @param _result The boolean value to set
     * @return success Whether the operation was successful
     */
    function setResult(bool _result) external returns (bool success) {
        require(trustedWallets[msg.sender], "Only trusted wallets can set the result");
        require(!resultIsSet, "Result has already been set");
        
        result = _result;
        resultIsSet = true;
        
        emit ResultSet(_result, msg.sender);
        return true;
    }
    
    /**
     * @dev Gets the current result
     * @return value The current result value
     * @return isSet Whether the result has been set
     */
    function getResult() external view returns (bool value, bool isSet) {
        return (result, resultIsSet);
    }
    
    /**
     * @dev Checks if an address is a trusted wallet
     * @param _wallet The address to check
     * @return Whether the address is a trusted wallet
     */
    function isTrustedWallet(address _wallet) external view returns (bool) {
        return trustedWallets[_wallet];
    }
}