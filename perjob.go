package benchmark



func ExecutePerJob(jobType string, doneChannel chan int, useChannel bool) {
	switch jobType {
	case "SUMUNTIL10000":
		go sumUntil10thousand(doneChannel, useChannel)
	}
}

