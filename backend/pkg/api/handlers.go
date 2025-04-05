package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/korjavin/predictionMarket/backend/pkg/ethereum"
)

// Handler struct holds dependencies for API handlers
type Handler struct {
	ethService *ethereum.Service
}

// NewHandler creates a new Handler instance
func NewHandler(ethService *ethereum.Service) *Handler {
	return &Handler{
		ethService: ethService,
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

// Placeholder for handlers to be implemented
func (h *Handler) DeployOracle(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, Response{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *Handler) GetOracles(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, Response{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *Handler) GetOracle(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, Response{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *Handler) SetOracleResult(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, Response{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *Handler) DeployBet(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, Response{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *Handler) GetBets(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, Response{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *Handler) GetBet(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, Response{
		Success: false,
		Error:   "Not implemented yet",
	})
}

func (h *Handler) CloseBet(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, Response{
		Success: false,
		Error:   "Not implemented yet",
	})
}
