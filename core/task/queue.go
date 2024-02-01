package task

var GlobalTaskQueue = make(chan Job, 10)
