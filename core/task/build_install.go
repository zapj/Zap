package task

import (
	"context"
	"os/exec"
)

type BuildInstallJob struct {
	Ctx context.Context
}

func (b *BuildInstallJob) Execute() error {
	if err := exec.CommandContext(b.Ctx, "sleep", "60").Run(); err != nil {
		return err
	}
	return nil
}
