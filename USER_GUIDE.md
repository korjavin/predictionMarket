# TinyBet Prediction Market - User Guide

Welcome to TinyBet, a simple, decentralized prediction market platform built on the Ethereum blockchain. This guide will help you understand how to use the platform, even if you're new to cryptocurrency and blockchain applications.

## What is a Prediction Market?

A prediction market is a platform where users can bet on the outcome of future events. In TinyBet, these outcomes are binary (yes/no or true/false). For example, "Will it rain in New York on December 25th?" or "Will candidate X win the election?".

Users place bets with cryptocurrency (in this case, Ethereum/ETH) on the outcome they believe will happen. After the event occurs, those who bet correctly receive rewards proportional to their bet amounts.

## Key Concepts

### Oracle

An oracle is a trusted source of truth that determines the outcome of the event. In TinyBet, an oracle is a smart contract with designated "trusted wallets" that can set the result (TRUE or FALSE) after the event occurs.

### Prediction Market (Bet)

A prediction market (or "bet") is where users place their bets on either the TRUE or FALSE outcome of a question. Each bet is connected to an oracle that will determine its outcome.

### Release Date

Each prediction market has a release date after which betting is closed and the oracle's result can be used to distribute rewards.

## Getting Started

### Prerequisites

1. **Ethereum Wallet**: You need a wallet like MetaMask to interact with the platform.
2. **Test ETH**: For the testnet version, you'll need test ETH (not real money).

### Setting Up MetaMask

1. Install the [MetaMask browser extension](https://metamask.io/download.html)
2. Create or import a wallet
3. Connect to the Ethereum testnet (we use Sepolia)

### Getting Test ETH

You can get test ETH for Sepolia testnet from:
- [Sepolia Faucet](https://sepoliafaucet.com/)
- [Alchemy Sepolia Faucet](https://sepoliafaucet.com/)
- [QuickNode Sepolia Faucet](https://faucet.quicknode.com/ethereum/sepolia)

For more information about Ethereum networks and faucets, visit the [Ethereum Developer Documentation](https://ethereum.org/en/developers/docs/networks/).

## Using TinyBet

### Connecting Your Wallet

1. Visit the TinyBet website
2. Click the "Connect Wallet" button in the top right
3. Approve the connection in your MetaMask

### Creating an Oracle

Oracles should be created for events whose outcome can be definitively determined:

1. Go to the "Oracles" page
2. Fill out the "Create New Oracle" form:
   - Enter wallet addresses that are trusted to report the outcome (comma-separated)
   - These addresses will have permission to set the True/False result
3. Click "Deploy Oracle"
4. Confirm the transaction in MetaMask
5. Wait for the transaction to be mined
6. The new oracle will appear in the "Deployed Oracles" list

### Creating a Prediction Market

1. Go to the "Bets" page
2. Fill out the "Create New Prediction Market" form:
   - Oracle Address: The address of an existing oracle
   - Question: What is being predicted (e.g., "Will BTC price be above $50K on Jan 1?")
   - Release Date: When betting will close and results can be determined
3. Click "Deploy Prediction Market"
4. Confirm the transaction in MetaMask
5. Wait for the transaction to be mined
6. The new market will appear in the "Active Prediction Markets" list

### Placing a Bet

1. Find a prediction market you want to participate in
2. Enter the amount of ETH you want to bet
3. Click either "Bet TRUE" or "Bet FALSE"
4. Confirm the transaction in MetaMask
5. Wait for the transaction to be mined
6. Your bet will be added to the market's statistics

### Setting an Oracle Result (for Trusted Wallets only)

If your wallet is designated as a trusted wallet for an oracle:

1. Go to the "Oracles" page
2. Find the oracle you're trusted for
3. Once the event has occurred, click either "Set TRUE" or "Set FALSE"
4. Confirm the transaction in MetaMask
5. Wait for the transaction to be mined
6. The oracle's result will be updated

### Closing a Bet and Distributing Funds

After the release date has passed and the oracle result has been set:

1. Find the corresponding prediction market
2. Click "Close Bet & Distribute Funds"
3. Confirm the transaction in MetaMask
4. Wait for the transaction to be mined
5. Funds will be automatically distributed to winners

### Checking Your Winnings

Winnings are automatically sent to the ETH address that placed the winning bet. You can check your MetaMask wallet to see if you've received any ETH from winning bets.

## Understanding the Interface

### Odds and Statistics

Each prediction market displays:
- The percentage of funds bet on TRUE vs FALSE
- Total amount of ETH bet on each outcome
- A visual representation of the current odds
- Current status (Active, Release Date Reached, or Closed)

### Status Indicators

- **Active**: Betting is open
- **Release Date Reached**: Betting is closed but funds haven't been distributed yet
- **Closed**: The bet has been finalized and funds have been distributed

## FAQ

**Q: How are winnings calculated?**

A: Winnings are calculated proportionally to your bet amount. If you bet on the winning outcome, you'll receive a share of the total pool proportional to your contribution to the winning side.

**Q: What happens if nobody bets on the winning outcome?**

A: In the unlikely event that no one bets on the winning outcome, the contract code ensures that no funds are distributed.

**Q: Can I cancel my bet?**

A: No, once a bet is placed, it cannot be cancelled or withdrawn until the market is closed.

**Q: What if the oracle never sets a result?**

A: If the trusted wallets never set a result for the oracle, the funds will remain locked in the contract indefinitely. This is why it's important to only participate in markets with reliable oracles.

**Q: Are there fees for using TinyBet?**

A: There are no platform fees, but you'll pay Ethereum network transaction fees (gas) for each action you take.

## Safety Tips

1. Only bet with funds you can afford to lose
2. Verify the oracle address and trusted wallets before participating
3. Double-check the release date to ensure it aligns with the event
4. Remember that testnet ETH has no real value and is only for testing
5. When using the mainnet version, all transactions involve real ETH

## Need Help?

If you have any questions or need assistance, please visit our GitHub repository at [github.com/korjavin/predictionMarket](https://github.com/korjavin/predictionMarket) and create an issue.