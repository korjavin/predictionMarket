package ethereum

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/korjavin/predictionMarket/backend/pkg/config"
)

// Service provides methods to interact with Ethereum blockchain
type Service struct {
	client      *ethclient.Client
	config      *config.Config
	privateKey  *ecdsa.PrivateKey
	fromAddress common.Address
}

// NewService creates a new Ethereum service
func NewService(cfg *config.Config) (*Service, error) {
	client, err := ethclient.Dial(cfg.EthereumRPC)
	if err != nil {
		return nil, err
	}

	service := &Service{
		client: client,
		config: cfg,
	}

	// Set up private key if provided
	if cfg.PrivateKey != "" {
		if err := service.setPrivateKey(cfg.PrivateKey); err != nil {
			return nil, err
		}
	}

	return service, nil
}

// setPrivateKey configures the service with a private key for transaction signing
func (s *Service) setPrivateKey(privateKeyHex string) error {
	// Remove 0x prefix if present
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	// Decode private key from hex
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return err
	}

	// Create private key
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return err
	}

	s.privateKey = privateKey
	s.fromAddress = crypto.PubkeyToAddress(privateKey.PublicKey)
	return nil
}

// CreateTransactOpts creates transaction options for contract interactions
func (s *Service) CreateTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	if s.privateKey == nil {
		return nil, errors.New("private key not set")
	}

	chainID, err := s.client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	nonce, err := s.client.PendingNonceAt(ctx, s.fromAddress)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(s.privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // gas limit
	auth.GasPrice = gasPrice

	return auth, nil
}

// GetClient returns the underlying ethclient
func (s *Service) GetClient() *ethclient.Client {
	return s.client
}

// GetFromAddress returns the address derived from the private key
func (s *Service) GetFromAddress() common.Address {
	return s.fromAddress
}

// Close closes the Ethereum client connection
func (s *Service) Close() {
	s.client.Close()
}
