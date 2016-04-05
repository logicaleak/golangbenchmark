package benchmark
import (
	"time"
)

func continuousProcessRequest(jobType string, numberPerSecond int) {
	for {
		for i := 0; i < numberPerSecond; i++ {
				ExecutePerJob(jobType, make(chan int), false)
		}
		time.Sleep(time.Second)
	}
}