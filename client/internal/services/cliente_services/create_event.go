package cliente_services

import (
	"context"
	"errors"
	mailing "mail_service/internal"
	"mail_service/internal/kit/event"
)

type IncreaseCoursesCounterOnCourseCreated struct {
	increasingService CourseCounterService
}

func NewIncreaseCoursesCounterOnCourseCreated(increaserService CourseCounterService) IncreaseCoursesCounterOnCourseCreated {
	return IncreaseCoursesCounterOnCourseCreated{
		increasingService: increaserService,
	}
}

func (e IncreaseCoursesCounterOnCourseCreated) Handle(_ context.Context, evt event.Event) error {
	courseCreatedEvt, ok := evt.(mailing.ClienteCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}

	return e.increasingService.Increase(courseCreatedEvt.ID())
}
