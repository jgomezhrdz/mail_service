package clientes

import (
	mailing "mail_service/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHandler(clienteReposiroty mailing.ClienteRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := clienteReposiroty.Get(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		var jsonData []struct {
			Client map[string]interface{}
			Plan   map[string]interface{}
		}
		for _, instance := range result {
			clientJSON := mailing.Cliente.TOJSON(instance.Client) // Assuming ToJSON is a method in mailing.Cliente
			planJSON := mailing.Plan.TOJSON(instance.Plan)        // Assuming ToJSON is a method in mailing.Plan

			jsonData = append(jsonData, struct {
				Client map[string]interface{}
				Plan   map[string]interface{}
			}{Client: clientJSON, Plan: planJSON})
		}

		ctx.JSON(http.StatusOK, gin.H{"data": jsonData})
	}
}
