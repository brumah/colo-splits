package main

func main() {
	fileChan := make(fileChannel)
	timesChan := make(timesChannel)
	go pushFiles("splits/", fileChan)
	go readFiles(fileChan, timesChan)
	readTimes(timesChan)
}
