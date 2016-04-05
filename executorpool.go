package benchmark

type ExecutorPoolTestContext struct {
	jobChannel chan Job
	jobType string
	executorNumber int
}

func NewExecutorPoolTestContext(jobType string, executorNumber int) *ExecutorPoolTestContext {
	eptc := ExecutorPoolTestContext{jobType:jobType, jobChannel:make(chan Job), executorNumber:executorNumber}
	return &eptc
}

func (selfPtr *ExecutorPoolTestContext) StartExecutors() {
	for i := 0; i < selfPtr.executorNumber; i++ {
		go Executor(selfPtr)
	}
}


func Executor(eptc *ExecutorPoolTestContext) {
	for {
		job := <- eptc.jobChannel
		jobType := eptc.jobType

		useChannel := job.useChannel
		doneChannel := job.doneChannel
		switch jobType {
		case "SUMUNTIL10000":
			sumUntil10thousand(doneChannel, useChannel)
		}

	}
}

func (selfPtr *ExecutorPoolTestContext) ExecuteWithPool(jobType string, doneChannel chan int, useChannel bool) {
	selfPtr.jobChannel <- Job{useChannel:useChannel, typ:jobType, doneChannel:doneChannel}
}
