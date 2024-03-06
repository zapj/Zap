package task

import "github.com/robfig/cron/v3"

type Scheduler struct {
	ID    int
	Retry int
	C     *cron.Cron
}
