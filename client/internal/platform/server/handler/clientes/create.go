package clientes

import (
	"errors"
	"net/http"

	"mail_service/internal/kit/types"
	cliente_services "mail_service/internal/services/cliente"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	IDCliente string `json:"idCliente" binding:"required"`
	Nombre    string `json:"nombre"    binding:"required"`
	IDPlan    string `json:"idPlan"  binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(clienteService cliente_services.ClienteService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "request": req})
			return
		}

		err := clienteService.CreateCliente(ctx, req.IDCliente, req.Nombre, req.IDPlan)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err != nil {
			switch {
			case errors.Is(err, types.ErrEmptyString), errors.Is(err, types.ErrInvalidID), errors.Is(err, types.ErrNegativaString):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
