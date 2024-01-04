package cron_scheduler

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/robfig/cron"
)

type CronJob struct {
	second   string
	minute   string
	hour     string
	monthDay string
	month    string
	weekDay  string
	cmdName  string
	cmd      func()
}

const (
	VALIDPATTERN = "^(\\*|[0-5]?\\d)(?:-(\\*|[0-5]?\\d))?(?:\\/(\\d+))?$"
)

var ErrPatternInvalid = errors.New("invalid cron pattern")

func NewCronjob(cmdName, second, minute, hour, monthDay, month, weekDay string, cmd func()) (CronJob, error) {
	if !isValid(second) ||
		!isValid(minute) ||
		!isValid(hour) ||
		!isValid(monthDay) ||
		!isValid(month) ||
		!isValid(weekDay) {
		return CronJob{}, fmt.Errorf("%s: "+ErrPatternInvalid.Error(), cmdName)
	}

	return CronJob{
		cmdName:  cmdName,
		second:   second,
		minute:   minute,
		hour:     hour,
		monthDay: monthDay,
		month:    month,
		weekDay:  weekDay,
		cmd:      cmd,
	}, nil
}

func isValid(value string) bool {
	matched, _ := regexp.MatchString(VALIDPATTERN, value)
	return matched
}

func (c CronJob) parseSchedule() string {
	return c.second + " " + c.minute + " " + c.hour + " " + c.monthDay + " " + c.month + " " + c.weekDay
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
		log.Print("Defining Job " + job.cmdName + " at: " + job.parseSchedule())
		cronManager.AddFunc(job.parseSchedule(), job.cmd)
	}
}

func Start(cronManager *cron.Cron) {
	cronManager.Start()
}
