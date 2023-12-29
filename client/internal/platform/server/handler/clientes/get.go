package clientes

import (
	cliente_services "mail_service/internal/services/cliente_services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHandler(clienteService cliente_services.ClienteService) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		data, err := clienteService.GetCliente(ctx)

		if err != nil {
			switch {
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.JSON(http.StatusOK, gin.H{"data": data})
	}
}
