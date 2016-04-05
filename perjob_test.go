package benchmark
import (
	"testing"
)



func BenchmarkPerjob(b *testing.B) {
	go continuousProcessRequest("SUMUNTIL10000", 10000)
	for i := 0; i < b.N; i++ {
		//There is overhead of channel creation
		doneChannel := make(chan int)
		ExecutePerJob("SUMUNTIL10000", doneChannel, true)
		<- doneChannel
	}
}

fu