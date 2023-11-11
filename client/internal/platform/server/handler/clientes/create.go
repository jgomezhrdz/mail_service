package clientes

import (
	"net/http"

	mailing "mail_service/internal"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	IDCliente string `json:"idCliente" binding:"required"`
	Nombre    string `json:"nombre"    binding:"required"`
	IDPlan    string `json:"idPlan"  binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(clienteReposiroty mailing.ClienteRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "request": req})
			return
		}

		cliente, err := mailing.NewCliente(req.IDCliente, req.Nombre, req.IDPlan)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := clienteReposiroty.Save(ctx, cliente); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, cliente.ID().Value())
	}
}
