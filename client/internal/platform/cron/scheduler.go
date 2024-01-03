package cron_scheduler

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/robfig/cron"
)

type CronJob struct {
	minute   string
	hour     string
	monthDay string
	month    string
	weekDay  string
	cmdName  string
	cmd      func()
}

var ErrPatternInvalid = errors.New("invalid cron pattern")

func NewCronjob(cmdName, minute, hour, monthDay, month, weekDay string, cmd func()) (CronJob, error) {
	validValue := "^(\\*|[0-9]{1,2}|[0-9]{1,2}-[0-9]{1,2}|[0-9]{1,2},[0-9]{1,2}(,[0-9]{1,2})*)$"
	if !isValid(minute, validValue) ||
		!isValid(hour, validValue) ||
		!isValid(monthDay, validValue) ||
		!isValid(month, validValue) ||
		!isValid(weekDay, validValue) {
		return CronJob{}, fmt.Errorf("%s: "+ErrPatternInvalid.Error(), cmdName)
	}

	return CronJob{
		cmdName:  cmdName,
		minute:   minute,
		hour:     hour,
		monthDay: monthDay,
		month:    month,
		weekDay:  weekDay,
		cmd:      cmd,
	}, nil
}

func isValid(value string, pattern string) bool {
	matched, _ := regexp.MatchString(pattern, value)
	return matched
}

func (c CronJob) parseSchedule() string {
	return c.minute + " " + c.hour + " " + c.monthDay + " " + c.month + " " + c.weekDay
}

func InitializeCron() *cron.Cron {
	loc, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}
	return cron.NewWithLocation(loc)
}

func RegisterJob(cronManager *cron.Cron, jobs ...CronJob) {
	for _, job := range jobs {
		cronManager.AddFunc(job.parseSchedule(), job.cmd)
	}
}

func Start(cronManager *cron.Cron) {
	cronManager.Start()
}
