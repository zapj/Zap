package task

type BuildInstallJob struct {
	ZapJob
}

func (b *BuildInstallJob) Execute() {
	b.ZapJob.Execute()
}
