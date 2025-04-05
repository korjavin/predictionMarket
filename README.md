# TinyBet Prediction Market

A simple, decentralized prediction market platform built on Ethereum. This application allows users to create oracles, set up prediction markets, and participate in betting on binary outcomes.

## Project Structure

- `contracts/`: Ethereum smart contracts
  - `src/`: Contract source code (Solidity)
  - `artifacts/`: Compiled contract artifacts
- `backend/`: Go backend service
  - `cmd/server/`: Main server entry point
  - `pkg/`: Package code (API handlers, Ethereum service, etc.)
- `web/`: Web frontend
  - `src/`: HTML, CSS, and JavaScript files
- `.github/workflows/`: CI/CD pipeline configuration

## Technical Overview

### Smart Contracts

1. **TinyOracle**
   - Stores a boolean value (true/false)
   - Can only be set once by trusted wallets defined at deploy time
   - Once set, the value is immutable

2. **TinyBet**
   - References a TinyOracle to determine the outcome
   - Allows users to place ETH bets on true/false outcomes
   - After the release date, funds can be distributed to winners based on the Oracle's value

### Backend Service

A Go (Golang) server that provides:
- REST API for deploying and interacting with smart contracts
- Endpoints for retrieving oracle and bet information
- CORS-enabled for web frontend communication

### Web Interface

A simple UI built with HTML, CSS, and vanilla JavaScript that provides:
- Oracle creation and management
- Prediction market creation
- Betting interface
- Visual representation of current bets and odds

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Node.js and npm (for contract development)
- Go (for backend development)
- An Ethereum wallet (like MetaMask)
- Access to Ethereum testnet (Goerli, Sepolia, etc.)
- API key for an Ethereum RPC provider (Infura, Alchemy, etc.)

### Environment Variables

Create a `.env` file in the project root:

```
ETHEREUM_RPC=https://goerli.infura.io/v3/your-api-key
NETWORK_ID=5
PRIVATE_KEY=your-private-key-without-0x-prefix
```

### Running with Docker Compose

```bash
docker-compose up -d
```

This will start both the backend API service and the web frontend.

### Local Development

#### Smart Contracts

```bash
cd contracts
npm install
npx hardhat compile
```

#### Backend

```bash
cd backend
go mod download
go run cmd/server/main.go
```

#### Frontend

```bash
cd web/src
# Use any simple HTTP server to serve the static files
python -m http.server 8000
```

## API Documentation

### Oracles

- `POST /api/v1/oracles`: Deploy a new oracle
- `GET /api/v1/oracles`: List all oracles
- `GET /api/v1/oracles/:address`: Get oracle details
- `POST /api/v1/oracles/:address/set-result`: Set oracle result

### Bets

- `POST /api/v1/bets`: Deploy a new bet
- `GET /api/v1/bets`: List all bets
- `GET /api/v1/bets/:address`: Get bet details
- `POST /api/v1/bets/:address/close`: Close bet and distribute funds

## Testing

### Smart Contracts

```bash
cd contracts
npx hardhat test
```

### Backend

```bash
cd backend
go test ./...
```

## Deployment

The GitHub Actions CI/CD pipeline will build and push Docker images when code is pushed to the main branch.

For manual deployment:

```bash
# Build and push backend image
docker build -t korjavin/prediction-market-backend:latest ./backend
docker push korjavin/prediction-market-backend:latest

# Build and push frontend image
docker build -t korjavin/prediction-market-frontend:latest ./web
docker push korjavin/prediction-market-frontend:latest
```

## License

MIT