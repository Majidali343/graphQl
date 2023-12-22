package handlers

import (
	"net/http"

	jokesData "graphQl/assets"
	utils "graphQl/utils"

	"github.com/gin-gonic/gin"

	"github.com/graphql-go/handler"
)

func GetJokesHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"data": jokesData.JokesData})
}

func GraphqlHandler(c *gin.Context) {
	h := handler.New(&handler.Config{
		Schema:   &utils.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	h.ServeHTTP(c.Writer, c.Request)
}
