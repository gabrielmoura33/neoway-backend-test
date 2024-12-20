package handler

import (
	"net/http"

	"github.com/gabrielmoura33/neoway-backend-test/config"
	"github.com/gabrielmoura33/neoway-backend-test/domain"
	"github.com/gabrielmoura33/neoway-backend-test/usecase"
	"github.com/gin-gonic/gin"
	"strings"
)

type ClientHandler struct {
	uc usecase.ClientUseCase
}

func NewClientHandler(uc usecase.ClientUseCase) *ClientHandler {
	return &ClientHandler{uc: uc}
}

func (h *ClientHandler) CreateClient(c *gin.Context) {
	config.IncrementRequestCount()

	var input struct {
		Document  string `json:"document" binding:"required"`
		Name      string `json:"name" binding:"required"`
		IsBlocked bool   `json:"is_blocked"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	docClean := removeNonDigits(input.Document)
	var clientType domain.ClientType
	if len(docClean) == 11 {
		clientType = domain.ClientTypeIndividual
	} else if len(docClean) == 14 {
		clientType = domain.ClientTypeCompany
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid document length"})
		return
	}

	client := domain.Client{
		Document:  input.Document,
		Name:      input.Name,
		IsBlocked: input.IsBlocked,
		Type:      clientType,
	}

	if err := h.uc.CreateClient(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, client)
}

func (h *ClientHandler) GetAllClients(c *gin.Context) {
	config.IncrementRequestCount()

	filterName := c.Query("name")
	clients, err := h.uc.GetAllClients(filterName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clients)
}

func (h *ClientHandler) GetClientByDocument(c *gin.Context) {
	config.IncrementRequestCount()

	document := c.Param("document")
	client, err := h.uc.GetClientByDocument(document)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if client == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "client not found"})
		return
	}

	c.JSON(http.StatusOK, client)
}

func (h *ClientHandler) ClientExists(c *gin.Context) {
	config.IncrementRequestCount()

	document := c.Query("document")
	exists, err := h.uc.ClientExists(document)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"exists": exists})
}

func (h *ClientHandler) Status(c *gin.Context) {
	config.IncrementRequestCount()

	c.JSON(http.StatusOK, gin.H{
		"uptime":         config.Uptime().String(),
		"request_count":  config.RequestCount,
	})
}

func removeNonDigits(s string) string {
	r := strings.NewReplacer(".", "", "-", "", "/", "")
	return r.Replace(s)
}
