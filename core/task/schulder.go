package task

import "github.com/robfig/cron/v3"

type Schulder struct {
	ID    int
	Retry int
	C     *cron.Cron
}
