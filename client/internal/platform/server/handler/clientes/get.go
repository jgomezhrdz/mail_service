package clientes

import (
	cliente_services "mail_service/internal/services/cliente_services"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetHandler(clienteService cliente_services.ClienteService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params, err := url.ParseQuery(ctx.Request.URL.RawQuery)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		// Process the parsed query parameters
		data, err := clienteService.GetCliente(ctx, params)

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
			if len(response) == 0 {
				response = []interface{}{}
			}
		}
		ctx.JSON(http.StatusOK, gin.H{"data": response})
	}
}
