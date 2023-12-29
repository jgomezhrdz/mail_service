package clientes

import (
	"bytes"
	"encoding/json"
	mailing "mail_service/internal"
	"mail_service/internal/kit/event"
	"mail_service/internal/kit/event/eventmocks"
	"mail_service/internal/platform/storage/storagemocks"
	cliente_services "mail_service/internal/services/cliente_services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T) {
	repositoryMock := new(storagemocks.ClienteRepository)
	repositoryMock.On(
		"Save",
		mock.Anything,
		mock.Anything,
	).Return(nil)

	eventBusMock := new(eventmocks.Bus)

	createClienteSrv := cliente_services.NewClienteService(repositoryMock, eventBusMock)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/clientes", CreateHandler(createClienteSrv))

	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		createCourseReq := createRequest{
			IDCliente: uuid.NewString(),
			Nombre:    "Test Cliente",
		}

		b, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/clientes", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		eventBusMock.On("Publish", mock.Anything, mock.MatchedBy(func(events []event.Event) bool {
			evt := events[0].(mailing.ClienteCreatedEvent)
			return evt.CourseName() == createCourseReq.Nombre
		})).Return(nil)

		eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		createCourseReq := createRequest{
			IDCliente: uuid.NewString(),
			Nombre:    "Demo Course",
			IDPlan:    uuid.NewString(),
		}

		b, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/clientes", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
