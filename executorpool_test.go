package benchmark
import (
	"testing"
	"time"
)

func continuousRequests(numberPerSecond int, eptc *ExecutorPoolTestContext) {
	dummyChannel := make(chan int)
	for {
		for i := 0; i < numberPerSecond; i++ {
			eptc.ExecuteWithPool("SUMUNTIL10000", dummyChannel, false)
		}
		time.Sleep(time.Second)
	}
}

func BenchmarkSumUntil10000ExecutorPool(b *testing.B) {
	jobType := "SUMUNTIL10000"

	executorPoolTestContext := NewExecutorPoolTestContext(jobType, 10)
	executorPoolTestContext.StartExecutors()

	go continuousRequests(10000, executorPoolTestContext)
	for i := 0; i < b.N; i++ {
		doneChannel := make(chan int)
		executorPoolTestContext.ExecuteWithPool(jobType, doneChannel, true)
		<- doneChannel
	}
}