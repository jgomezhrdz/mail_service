package server

import (
	"log"
	cron_scheduler "mail_service/internal/platform/cron"
	cron_handler "mail_service/internal/platform/cron/handlers"
)

func (s *Server) registerJobs() {
	if job, err := cron_scheduler.NewCronjob("HELLO WORLD", "*", "*", "*", "asd", "*", cron_handler.HelloWorld); err == nil {
		cron_scheduler.RegisterJob(s.cronScheduler, job)
	} else {
		log.Print(err)
	}
}
