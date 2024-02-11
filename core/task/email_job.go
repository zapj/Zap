package task

import (
	"context"
	"os/exec"
)

type EmailJob struct {
	Ctx context.Context
}

func (b *EmailJob) Execute() error {
	if err := exec.CommandContext(b.Ctx, "sleep", "5").Run(); err != nil {
		return err
	}
	return nil
}
