package clientes

import (
	cliente_services "mail_service/internal/services/cliente_services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHandler(clienteService cliente_services.ClienteService) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		data, err := clienteService.GetCliente(ctx)

		var response []interface{}
		if err != nil {
			switch {
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			for _, item := range data {
				response = append(response, struct {
					Client interface{}
					Plan   interface{}
				}{
					Client: item.Client.TOJSON(),
					Plan:   item.Plan.TOJSON(),
				})
			}
		}

		ctx.JSON(http.StatusOK, gin.H{"data": response})
	}
}
