package task

import "context"

type BashJob struct {
	Ctx context.Context
}

func (b *BashJob) Execute() error {
	return nil
}
