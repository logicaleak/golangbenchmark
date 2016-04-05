package benchmark

type Job struct {
	useChannel bool
	doneChannel chan int
	typ string
}

func sumUntil10thousand(doneChannel chan int, useChannel bool) {
	var sum int
	for i := 0; i < 100000; i++ {
		sum += i
	}
	if useChannel {
		doneChannel <- 1
	}
}

func sumUntilNumber(doneChannel chan int, useChannel bool, until int) {
	var sum int
	for i := 0; i < until; i++ {
		sum += i
	}
	if useChannel {
		doneChannel <- 1
	}
}