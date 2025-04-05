package api

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"sort"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/korjavin/predictionMarket/backend/pkg/ethereum"
	"github.com/korjavin/predictionMarket/backend/pkg/ethereum/contracts"
)

// Use contracts package functions
var (
	DeployTinyOracle = contracts.DeployTinyOracle
	NewTinyOracle    = contracts.NewTinyOracle
	DeployTinyBet    = contracts.DeployTinyBet
	NewTinyBet       = contracts.NewTinyBet
)

// Handler struct holds dependencies for API handlers
type Handler struct {
	ethService *ethereum.Service
	db         *badger.DB
}

// NewHandler creates a new Handler instance
func NewHandler(ethService *ethereum.Service) *Handler {
	// Open BadgerDB
	db, err := badger.Open(badger.DefaultOptions("./badger-data"))
	if err != nil {
		panic(fmt.Sprintf("Failed to open BadgerDB: %v", err))
	}

	return &Handler{
		ethService: ethService,
		db:         db,
	}
}

// Close closes the database
func (h *Handler) Close() {
	if h.db != nil {
		h.db.Close()
	}
}

// Common response structure
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// DeployOracleRequest is the request structure for deploying a new Oracle
type DeployOracleRequest struct {
	TrustedWallets []string `json:"trustedWallets" binding:"required"`
}

// DeployOracleResponse is the response structure after deploying an Oracle
type DeployOracleResponse struct {
	Address        string   `json:"address"`
	TrustedWallets []string `json:"trustedWallets"`
	TxHash         string   `json:"txHash"`
}

// DeployBetRequest is the request structure for deploying a new Bet
type DeployBetRequest struct {
	OracleAddress string `json:"oracleAddress" binding:"required"`
	Question      string `json:"question" binding:"required"`
	ReleaseDate   int64  `json:"releaseDate" binding:"required"`
}

// DeployBetResponse is the response structure after deploying a Bet
type DeployBetResponse struct {
	Address       string `json:"address"`
	OracleAddress string `json:"oracleAddress"`
	Question      string `json:"question"`
	ReleaseDate   int64  `json:"releaseDate"`
	TxHash        string `json:"txHash"`
}

// GetOraclesResponse contains information about deployed oracles
type GetOraclesResponse struct {
	Address        string   `json:"address"`
	TrustedWallets []string `json:"trustedWallets"`
	ResultSet      bool     `json:"resultSet"`
	Result         *bool    `json:"result,omitempty"`
}

// GetBetsResponse contains information about deployed bets
type GetBetsResponse struct {
	Address        string    `json:"address"`
	OracleAddress  string    `json:"oracleAddress"`
	Question       string    `json:"question"`
	ReleaseDate    time.Time `json:"releaseDate"`
	IsClosed       bool      `json:"isClosed"`
	TotalTrueBets  string    `json:"totalTrueBets"`
	TotalFalseBets string    `json:"totalFalseBets"`
	TrueOdds       int       `json:"trueOdds"`
	FalseOdds      int       `json:"falseOdds"`
	ResultSet      bool      `json:"resultSet"`
	Result         *bool     `json:"result,omitempty"`
}

// SetOracleResultRequest is the request structure for setting an oracle result
type SetOracleResultRequest struct {
	Result bool `json:"result" binding:"required"`
}

// storeOracle stores oracle data in the database
func (h *Handler) storeOracle(address string, data []byte) error {
	return h.db.Update(func(txn *badger.Txn) error {
		key := []byte("oracle:" + address)
		return txn.Set(key, data)
	})
}

// getOracle retrieves oracle data from the database
func (h *Handler) getOracle(address string) ([]byte, error) {
	var valCopy []byte
	err := h.db.View(func(txn *badger.Txn) error {
		key := []byte("oracle:" + address)
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		valCopy, err = item.ValueCopy(nil)
		return err
	})
	return valCopy, err
}

// getAllOracles retrieves all oracle addresses from the database
func (h *Handler) getAllOracles() ([]string, error) {
	var addresses []string

	err := h.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = []byte("oracle:")

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := string(item.Key())
			addresses = append(addresses, k[len("oracle:"):])
		}
		return nil
	})

	return addresses, err
}

// storeBet stores bet data in the database
func (h *Handler) storeBet(address string, data []byte) error {
	return h.db.Update(func(txn *badger.Txn) error {
		key := []byte("bet:" + address)
		return txn.Set(key, data)
	})
}

// getBet retrieves bet data from the database
func (h *Handler) getBet(address string) ([]byte, error) {
	var valCopy []byte
	err := h.db.View(func(txn *badger.Txn) error {
		key := []byte("bet:" + address)
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		valCopy, err = item.ValueCopy(nil)
		return err
	})
	return valCopy, err
}

// getAllBets retrieves all bet addresses from the database
func (h *Handler) getAllBets() ([]string, error) {
	var addresses []string

	err := h.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = []byte("bet:")

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := string(item.Key())
			addresses = append(addresses, k[len("bet:"):])
		}
		return nil
	})

	return addresses, err
}

// Health check endpoint
func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    map[string]string{"status": "ok"},
	})
}

// Setup API routes
func SetupRouter(handler *Handler) *gin.Engine {
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Health check endpoint
	router.GET("/health", handler.Health)

	// API group with version
	v1 := router.Group("/api/v1")
	{
		// Oracle endpoints
		v1.POST("/oracles", handler.DeployOracle)
		v1.GET("/oracles", handler.GetOracles)
		v1.GET("/oracles/:address", handler.GetOracle)
		v1.POST("/oracles/:address/set-result", handler.SetOracleResult)

		// Bet endpoints
		v1.POST("/bets", handler.DeployBet)
		v1.GET("/bets", handler.GetBets)
		v1.GET("/bets/:address", handler.GetBet)
		v1.POST("/bets/:address/close", handler.CloseBet)
	}

	return router
}

// DeployOracle deploys a new TinyOracle contract
func (h *Handler) DeployOracle(c *gin.Context) {
	var req DeployOracleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request: " + err.Error(),
		})
		return
	}

	// Convert wallet addresses to Ethereum addresses
	var ethAddresses []common.Address
	for _, wallet := range req.TrustedWallets {
		if !common.IsHexAddress(wallet) {
			c.JSON(http.StatusBadRequest, Response{
				Success: false,
				Error:   fmt.Sprintf("Invalid Ethereum address: %s", wallet),
			})
			return
		}
		ethAddresses = append(ethAddresses, common.HexToAddress(wallet))
	}

	ctx := c.Request.Context()

	// Create transaction options
	auth, err := h.ethService.CreateTransactOpts(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to create transaction options: " + err.Error(),
		})
		return
	}

	// Deploy TinyOracle contract
	// Note: This assumes you have a TinyOracle contract binding from the generated Go code
	// You may need to modify this code to match your actual contract deployment method
	address, tx, err := DeployTinyOracle(auth, h.ethService.GetClient(), ethAddresses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to deploy oracle contract: " + err.Error(),
		})
		return
	}

	// Store oracle in database
	oracleData := DeployOracleResponse{
		Address:        address.Hex(),
		TrustedWallets: req.TrustedWallets,
		TxHash:         tx.Hash().Hex(),
	}

	oracleBytes, err := json.Marshal(oracleData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to serialize oracle data: " + err.Error(),
		})
		return
	}

	if err := h.storeOracle(address.Hex(), oracleBytes); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to store oracle in database: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    oracleData,
	})
}

// GetOracles retrieves all deployed oracle contracts
func (h *Handler) GetOracles(c *gin.Context) {
	addresses, err := h.getAllOracles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to get oracles: " + err.Error(),
		})
		return
	}

	var oracles []GetOraclesResponse
	for _, addr := range addresses {
		// Get oracle data from database
		oracleBytes, err := h.getOracle(addr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Error:   "Failed to get oracle data: " + err.Error(),
			})
			return
		}

		var oracleData DeployOracleResponse
		if err := json.Unmarshal(oracleBytes, &oracleData); err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Error:   "Failed to deserialize oracle data: " + err.Error(),
			})
			return
		}

		// Get current oracle state from blockchain
		oracle, err := NewTinyOracle(common.HexToAddress(addr), h.ethService.GetClient())
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Error:   "Failed to connect to oracle contract: " + err.Error(),
			})
			return
		}

		// Get the result and whether it's set
		result, isSet, err := oracle.GetResult(nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Error:   "Failed to get oracle result: " + err.Error(),
			})
			return
		}

		// Only include result if it has been set
		var resultPtr *bool
		if isSet {
			resultPtr = &result
		}

		oracles = append(oracles, GetOraclesResponse{
			Address:        addr,
			TrustedWallets: oracleData.TrustedWallets,
			ResultSet:      isSet,
			Result:         resultPtr,
		})
	}

	// Sort oracles by address for consistent ordering
	sort.Slice(oracles, func(i, j int) bool {
		return oracles[i].Address < oracles[j].Address
	})

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    oracles,
	})
}

// GetOracle retrieves a specific oracle contract by address
func (h *Handler) GetOracle(c *gin.Context) {
	address := c.Param("address")

	if !common.IsHexAddress(address) {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid Ethereum address",
		})
		return
	}

	// Get oracle data from database
	oracleBytes, err := h.getOracle(address)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   "Oracle not found: " + err.Error(),
		})
		return
	}

	var oracleData DeployOracleResponse
	if err := json.Unmarshal(oracleBytes, &oracleData); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to deserialize oracle data: " + err.Error(),
		})
		return
	}

	// Get current oracle state from blockchain
	oracle, err := NewTinyOracle(common.HexToAddress(address), h.ethService.GetClient())
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to connect to oracle contract: " + err.Error(),
		})
		return
	}

	// Get the result and whether it's set
	result, isSet, err := oracle.GetResult(nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to get oracle result: " + err.Error(),
		})
		return
	}

	// Only include result if it has been set
	var resultPtr *bool
	if isSet {
		resultPtr = &result
	}

	oracleResponse := GetOraclesResponse{
		Address:        address,
		TrustedWallets: oracleData.TrustedWallets,
		ResultSet:      isSet,
		Result:         resultPtr,
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    oracleResponse,
	})
}

// SetOracleResult sets the result of an oracle
func (h *Handler) SetOracleResult(c *gin.Context) {
	address := c.Param("address")

	if !common.IsHexAddress(address) {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid Ethereum address",
		})
		return
	}

	var req SetOracleResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request: " + err.Error(),
		})
		return
	}

	// Get oracle instance
	oracle, err := NewTinyOracle(common.HexToAddress(address), h.ethService.GetClient())
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to connect to oracle contract: " + err.Error(),
		})
		return
	}

	// Check if result is already set
	_, isSet, err := oracle.GetResult(nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to check if result is already set: " + err.Error(),
		})
		return
	}

	if isSet {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Oracle result is already set",
		})
		return
	}

	// Check if sender is trusted
	isTrusted, err := oracle.IsTrustedWallet(nil, h.ethService.GetFromAddress())
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to check if sender is trusted: " + err.Error(),
		})
		return
	}

	if !isTrusted {
		c.JSON(http.StatusForbidden, Response{
			Success: false,
			Error:   "Sender is not a trusted wallet for this oracle",
		})
		return
	}

	// Set the result
	ctx := c.Request.Context()
	auth, err := h.ethService.CreateTransactOpts(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to create transaction options: " + err.Error(),
		})
		return
	}

	tx, err := oracle.SetResult(auth, req.Result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to set oracle result: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data: map[string]interface{}{
			"address": address,
			"result":  req.Result,
			"txHash":  tx.Hash().Hex(),
		},
	})
}

// DeployBet deploys a new TinyBet contract
func (h *Handler) DeployBet(c *gin.Context) {
	var req DeployBetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request: " + err.Error(),
		})
		return
	}

	// Validate oracle address
	if !common.IsHexAddress(req.OracleAddress) {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid oracle address",
		})
		return
	}

	// Validate release date is in the future
	if req.ReleaseDate <= time.Now().Unix() {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Release date must be in the future",
		})
		return
	}

	// Check if oracle exists and can be interacted with
	oracleAddr := common.HexToAddress(req.OracleAddress)
	oracle, err := NewTinyOracle(oracleAddr, h.ethService.GetClient())
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to connect to oracle contract: " + err.Error(),
		})
		return
	}

	// Just call a function to ensure the contract exists
	_, _, err = oracle.GetResult(nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Oracle contract does not exist or cannot be accessed: " + err.Error(),
		})
		return
	}

	// Create transaction options for deploying the bet contract
	ctx := c.Request.Context()
	auth, err := h.ethService.CreateTransactOpts(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to create transaction options: " + err.Error(),
		})
		return
	}

	// Deploy new bet contract
	// Note: This assumes you have a TinyBet contract binding from the generated Go code
	// You may need to modify this code to match your actual contract deployment method
	betAddr, tx, err := DeployTinyBet(
		auth,
		h.ethService.GetClient(),
		oracleAddr,
		req.Question,
		big.NewInt(req.ReleaseDate),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to deploy bet contract: " + err.Error(),
		})
		return
	}

	// Store bet details in database
	betData := DeployBetResponse{
		Address:       betAddr.Hex(),
		OracleAddress: req.OracleAddress,
		Question:      req.Question,
		ReleaseDate:   req.ReleaseDate,
		TxHash:        tx.Hash().Hex(),
	}

	betBytes, err := json.Marshal(betData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to serialize bet data: " + err.Error(),
		})
		return
	}

	if err := h.storeBet(betAddr.Hex(), betBytes); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to store bet in database: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    betData,
	})
}

// GetBets retrieves all deployed bet contracts
func (h *Handler) GetBets(c *gin.Context) {
	addresses, err := h.getAllBets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to get bets: " + err.Error(),
		})
		return
	}

	var bets []map[string]interface{}
	for _, addr := range addresses {
		// Get bet data from database
		betBytes, err := h.getBet(addr)
		if err != nil {
			continue // Skip if bet cannot be retrieved
		}

		var betData DeployBetResponse
		if err := json.Unmarshal(betBytes, &betData); err != nil {
			continue // Skip if bet data cannot be deserialized
		}

		// Get bet contract instance
		bet, err := NewTinyBet(common.HexToAddress(addr), h.ethService.GetClient())
		if err != nil {
			continue // Skip if contract cannot be accessed
		}

		// Get bet information from the contract
		betInfo, err := bet.GetBetInfo(nil)
		if err != nil {
			continue // Skip if bet info cannot be fetched
		}

		// Get bet odds
		trueOdds, falseOdds, err := bet.GetOdds(nil)
		if err != nil {
			// Use default odds if there's an error
			trueOdds = big.NewInt(50)
			falseOdds = big.NewInt(50)
		}

		// Get oracle information
		oracle, err := NewTinyOracle(common.HexToAddress(betData.OracleAddress), h.ethService.GetClient())
		if err != nil {
			continue // Skip if oracle cannot be accessed
		}

		// Get the oracle result and whether it's set
		oracleResult, resultSet, err := oracle.GetResult(nil)
		if err != nil {
			continue // Skip if oracle result cannot be fetched
		}

		// Only include result if it has been set
		var resultPtr *bool
		if resultSet {
			resultPtr = &oracleResult
		}

		// Convert release date to time.Time for better serialization
		releaseDateTime := time.Unix(betInfo.ReleaseDate.Int64(), 0)

		// Create response object
		betResponse := map[string]interface{}{
			"address":        addr,
			"oracleAddress":  betData.OracleAddress,
			"question":       betInfo.Question,
			"releaseDate":    releaseDateTime,
			"isClosed":       betInfo.IsClosed,
			"totalTrueBets":  betInfo.TotalTrueBets.String(),
			"totalFalseBets": betInfo.TotalFalseBets.String(),
			"trueOdds":       trueOdds.Int64(),
			"falseOdds":      falseOdds.Int64(),
			"resultSet":      resultSet,
		}

		if resultPtr != nil {
			betResponse["result"] = *resultPtr
		}

		bets = append(bets, betResponse)
	}

	// Sort bets by release date (newest first)
	sort.Slice(bets, func(i, j int) bool {
		iTime := bets[i]["releaseDate"].(time.Time)
		jTime := bets[j]["releaseDate"].(time.Time)
		return iTime.After(jTime)
	})

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    bets,
	})
}

// GetBet retrieves a specific bet contract by address
func (h *Handler) GetBet(c *gin.Context) {
	address := c.Param("address")

	if !common.IsHexAddress(address) {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid Ethereum address",
		})
		return
	}

	// Get bet data from database
	betBytes, err := h.getBet(address)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   "Bet not found: " + err.Error(),
		})
		return
	}

	var betData DeployBetResponse
	if err := json.Unmarshal(betBytes, &betData); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to deserialize bet data: " + err.Error(),
		})
		return
	}

	// Get bet contract instance
	bet, err := NewTinyBet(common.HexToAddress(address), h.ethService.GetClient())
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to connect to bet contract: " + err.Error(),
		})
		return
	}

	// Get bet information from the contract
	betInfo, err := bet.GetBetInfo(nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to get bet info: " + err.Error(),
		})
		return
	}

	// Get bet odds
	trueOdds, falseOdds, err := bet.GetOdds(nil)
	if err != nil {
		// Use default odds if there's an error
		trueOdds = big.NewInt(50)
		falseOdds = big.NewInt(50)
	}

	// Get oracle information
	oracle, err := NewTinyOracle(common.HexToAddress(betData.OracleAddress), h.ethService.GetClient())
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to connect to oracle contract: " + err.Error(),
		})
		return
	}

	// Get the oracle result and whether it's set
	oracleResult, resultSet, err := oracle.GetResult(nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to get oracle result: " + err.Error(),
		})
		return
	}

	// Only include result if it has been set
	var resultPtr *bool
	if resultSet {
		resultPtr = &oracleResult
	}

	// Convert release date to time.Time for better serialization
	releaseDateTime := time.Unix(betInfo.ReleaseDate.Int64(), 0)

	betResponse := map[string]interface{}{
		"address":        address,
		"oracleAddress":  betData.OracleAddress,
		"question":       betInfo.Question,
		"releaseDate":    releaseDateTime,
		"isClosed":       betInfo.IsClosed,
		"totalTrueBets":  betInfo.TotalTrueBets.String(),
		"totalFalseBets": betInfo.TotalFalseBets.String(),
		"trueOdds":       trueOdds.Int64(),
		"falseOdds":      falseOdds.Int64(),
		"resultSet":      resultSet,
	}

	if resultPtr != nil {
		betResponse["result"] = *resultPtr
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    betResponse,
	})
}

// CloseBet closes a bet and distributes funds to winners
func (h *Handler) CloseBet(c *gin.Context) {
	address := c.Param("address")

	if !common.IsHexAddress(address) {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid Ethereum address",
		})
		return
	}

	// Get bet contract instance
	bet, err := NewTinyBet(common.HexToAddress(address), h.ethService.GetClient())
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to connect to bet contract: " + err.Error(),
		})
		return
	}

	// Get bet information from the contract
	betInfo, err := bet.GetBetInfo(nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to get bet info: " + err.Error(),
		})
		return
	}

	// Check if bet is already closed
	if betInfo.IsClosed {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Bet is already closed",
		})
		return
	}

	// Check if we're past the release date
	currentTime := time.Now().Unix()
	if currentTime < betInfo.ReleaseDate.Int64() {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Cannot close bet before release date",
		})
		return
	}

	// Get the oracle address
	betData := DeployBetResponse{}
	betBytes, err := h.getBet(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to get bet data: " + err.Error(),
		})
		return
	}

	if err := json.Unmarshal(betBytes, &betData); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to deserialize bet data: " + err.Error(),
		})
		return
	}

	// Check if the oracle result is set
	oracle, err := NewTinyOracle(common.HexToAddress(betData.OracleAddress), h.ethService.GetClient())
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to connect to oracle contract: " + err.Error(),
		})
		return
	}

	_, isSet, err := oracle.GetResult(nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to get oracle result: " + err.Error(),
		})
		return
	}

	if !isSet {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Oracle result is not set yet",
		})
		return
	}

	// Close the bet
	ctx := c.Request.Context()
	auth, err := h.ethService.CreateTransactOpts(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to create transaction options: " + err.Error(),
		})
		return
	}

	tx, err := bet.CloseBet(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to close bet: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data: map[string]interface{}{
			"address": address,
			"txHash":  tx.Hash().Hex(),
			"message": "Bet closed successfully. Funds are being distributed to winners.",
		},
	})
}
