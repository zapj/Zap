package task

import (
	"context"
	"os/exec"
)

type HttpJob struct {
	Ctx context.Context
}

func (b *HttpJob) Execute() error {
	if err := exec.CommandContext(b.Ctx, "sleep", "5").Run(); err != nil {
		return err
	}
	return nil
}
