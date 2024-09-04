package controllers

import (
	"net/http"
	"web-api/src/repository"

	"github.com/gin-gonic/gin"
)

func ListCryptocurrencies(c *gin.Context) {
	cryptos := repository.GetAllCryptocurrencies()
	c.JSON(200, cryptos)
}

func SearchCryptocurrency(c *gin.Context) {
	term := c.Query("term")
	if term == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search term is required"})
		return
	}

	cryptocurrencies, err := repository.SearchCryptocurrencies(term)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cryptocurrencies)
}
