package task

import "time"

type Result struct {
	task       Task
	successful bool
}

func (result *Result) GetTask() Task {
	return result.task
}

func (result *Result) IsSuccessful() bool {
	return result.successful
}

func initResult(t Task) *Result {
	return &Result{
		successful: false,
		task:       t,
	}
}

func getTimeMs() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}
