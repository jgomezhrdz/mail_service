package clientes

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"mail_service/internal/kit/types"
	cliente_services "mail_service/internal/services/cliente_services"

	"github.com/gin-gonic/gin"
)

// CreateHandler returns an HTTP handler for courses creation.
func DeleteHandler(clienteService cliente_services.ClienteService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params, err := url.ParseQuery(ctx.Request.URL.RawQuery)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if ok := params.Has("id"); !ok {
			ctx.JSON(http.StatusInternalServerError, fmt.Errorf("es necesario especificar el elemento a borrar"))
			return
		}

		err = clienteService.DeleteCliente(ctx, params.Get("id"))

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
