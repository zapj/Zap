package task

import (
	"context"
	"os/exec"

	"github.com/zapj/zap/core/models"
)

type InstallJob struct {
	Ctx      context.Context
	TaskData *models.ZapTask
}

func (b *InstallJob) Execute() error {
	if err := exec.CommandContext(b.Ctx, "sleep", "30").Run(); err != nil {
		return err
	}
	return nil
}
