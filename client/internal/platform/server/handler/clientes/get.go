package clientes

import (
	mailing "mail_service/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHandler(clienteReposiroty mailing.ClienteRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := clienteReposiroty.Get(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, "")
	}
}
