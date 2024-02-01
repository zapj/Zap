package task

import "github.com/robfig/cron/v3"

type InstallApp struct {
	ID    int
	Retry int
	C     *cron.Cron
}

type Job struct {
	ID    int
	Retry int
	C     *cron.Cron
}
