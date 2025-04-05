// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./TinyOracle.sol";

/**
 * @title TinyBet
 * @dev This contract allows users to place bets on a boolean outcome.
 * After the release date, funds will be distributed to winners based on the TinyOracle's result.
 */
contract TinyBet {
    // Oracle contract to get the result from
    TinyOracle public oracle;
    
    // Question being bet on
    string public question;
    
    // Release date after which the bet can be closed
    uint256 public releaseDate;
    
    // Total amount bet on true and false outcomes
    uint256 public totalTrueBets;
    uint256 public totalFalseBets;
    
    // Track bets by address and outcome
    mapping(address => uint256) public trueBets;
    mapping(address => uint256) public falseBets;
    
    // Whether the bet has been closed
    bool public closed = false;
    
    // Array of all bettor addresses for easier iteration
    address[] public trueBettors;
    address[] public falseBettors;
    
    // Events
    event BetPlaced(address indexed bettor, bool prediction, uint256 amount);
    event BetClosed(bool result, uint256 totalPayout);
    event PayoutSent(address indexed bettor, uint256 amount);
    
    /**
     * @dev Constructor that initializes the bet parameters
     * @param _oracle Address of the TinyOracle contract
     * @param _question Description of what is being bet on
     * @param _releaseDate Timestamp after which the bet can be closed
     */
    constructor(address _oracle, string memory _question, uint256 _releaseDate) {
        require(_releaseDate > block.timestamp, "Release date must be in the future");
        oracle = TinyOracle(_oracle);
        question = _question;
        releaseDate = _releaseDate;
    }
    
    /**
     * @dev Allow users to bet on the true outcome
     */
    function betTrue() external payable {
        placeBet(true);
    }
    
    /**
     * @dev Allow users to bet on the false outcome
     */
    function betFalse() external payable {
        placeBet(false);
    }
    
    /**
     * @dev Internal function to handle both types of bets
     * @param prediction The user's bet (true or false)
     */
    function placeBet(bool prediction) internal {
        require(!closed, "Betting is closed");
        require(block.timestamp < releaseDate, "Betting period is over");
        require(msg.value > 0, "Bet amount must be greater than 0");
        
        if (prediction) {
            if (trueBets[msg.sender] == 0) {
                trueBettors.push(msg.sender);
            }
            trueBets[msg.sender] += msg.value;
            totalTrueBets += msg.value;
        } else {
            if (falseBets[msg.sender] == 0) {
                falseBettors.push(msg.sender);
            }
            falseBets[msg.sender] += msg.value;
            totalFalseBets += msg.value;
        }
        
        emit BetPlaced(msg.sender, prediction, msg.value);
    }
    
    /**
     * @dev Close the bet and distribute funds to winners
     */
    function closeBet() external {
        require(!closed, "Bet already closed");
        require(block.timestamp >= releaseDate, "Cannot close bet before release date");
        
        (bool result, bool isSet) = oracle.getResult();
        require(isSet, "Oracle result is not set yet");
        
        closed = true;
        
        uint256 totalPayout = totalTrueBets + totalFalseBets;
        emit BetClosed(result, totalPayout);
        
        // Distribute funds to winners
        if (result) {
            distributeFunds(trueBettors, trueBets, totalTrueBets, totalPayout);
        } else {
            distributeFunds(falseBettors, falseBets, totalFalseBets, totalPayout);
        }
    }
    
    /**
     * @dev Distribute funds to winners proportional to their bets
     * @param bettors Array of addresses that bet on the winning outcome
     * @param bets Mapping of address to bet amount
     * @param totalBets Total amount bet on the winning outcome
     * @param totalPayout Total amount to distribute
     */
    function distributeFunds(
        address[] memory bettors, 
        mapping(address => uint256) storage bets, 
        uint256 totalBets, 
        uint256 totalPayout
    ) internal {
        for (uint i = 0; i < bettors.length; i++) {
            address bettor = bettors[i];
            uint256 bet = bets[bettor];
            
            if (bet > 0) {
                uint256 payout = (bet * totalPayout) / totalBets;
                (bool success, ) = bettor.call{value: payout}("");
                if (success) {
                    emit PayoutSent(bettor, payout);
                }
            }
        }
    }
    
    /**
     * @dev Get the current state of the bet
     * @return Question being bet on
     * @return ReleaseDate timestamp
     * @return IsClosed whether the bet is closed
     * @return TotalTrueBets amount bet on true
     * @return TotalFalseBets amount bet on false
     */
    function getBetInfo() external view returns (
        string memory Question, 
        uint256 ReleaseDate, 
        bool IsClosed, 
        uint256 TotalTrueBets, 
        uint256 TotalFalseBets
    ) {
        return (
            question,
            releaseDate,
            closed,
            totalTrueBets,
            totalFalseBets
        );
    }
    
    /**
     * @dev Get the odds of each outcome
     * @return trueOdds Percentage of bets on true (0-100)
     * @return falseOdds Percentage of bets on false (0-100)
     */
    function getOdds() external view returns (uint256 trueOdds, uint256 falseOdds) {
        uint256 totalBets = totalTrueBets + totalFalseBets;
        
        if (totalBets == 0) {
            return (50, 50); // Default to 50/50 if no bets
        }
        
        trueOdds = (totalTrueBets * 100) / totalBets;
        falseOdds = 100 - trueOdds;
        
        return (trueOdds, falseOdds);
    }
}