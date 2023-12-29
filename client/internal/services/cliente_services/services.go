package cliente_services

import (
	mailing "mail_service/internal"
	"mail_service/internal/kit/event"
)

// CourseService is the default CourseService interface
// implementation returned by creating.NewCourseService.
type ClienteService struct {
	clienteRepository mailing.ClienteRepository
	eventBus          event.Bus
}

// NewCourseService returns the default Service interface implementation.
func NewClienteService(clienteRepository mailing.ClienteRepository, eventBus event.Bus) ClienteService {
	return ClienteService{
		clienteRepository: clienteRepository,
		eventBus:          eventBus,
	}
}

type CourseCounterService struct{}

func NewCourseCounterService() CourseCounterService {
	return CourseCounterService{}
}
