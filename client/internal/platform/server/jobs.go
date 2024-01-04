package server

import (
	"log"
	cron_scheduler "mail_service/internal/platform/cron"
	cron_handler "mail_service/internal/platform/cron/handlers"
	"os"
)

func (s *Server) registerJobs() {
	if job, err := cron_scheduler.NewCronjob("HELLO WORLD", "0", "*", "*", "*", "*", "*", cron_handler.HelloWorld); err == nil {
		cron_scheduler.RegisterJob(s.cronScheduler, job)
	} else {
		log.Print(err)
	}

	env := os.Getenv("GIN_MODE")

	if env == "release" {
		cron_scheduler.Start(s.cronScheduler)
	}
}
