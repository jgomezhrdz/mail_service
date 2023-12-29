package services

import (
	mailing "mail_service/internal"
	"mail_service/internal/kit/event"
)

// CourseService is the default CourseService interface
// implementation returned by creating.NewCourseService.
type PlanService struct {
	clienteRepository mailing.ClienteRepository
	eventBus          event.Bus
}

// NewCourseService returns the default Service interface implementation.
func NewClienteService(clienteRepository mailing.ClienteRepository, eventBus event.Bus) PlanService {
	return PlanService{
		clienteRepository: clienteRepository,
		eventBus:          eventBus,
	}
}

type CourseCounterService struct{}

func NewCourseCounterService() CourseCounterService {
	return CourseCounterService{}
}
