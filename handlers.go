package main

import (
	"html/template"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, "homepage")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	files := r.MultipartForm.File["folderInput"]

	testChan := make(fileChannel)
	timesChan := make(timesChannel)
	go pushFiles(files, testChan)
	go readFiles(testChan, timesChan)
	pbTimes, sumOfBestTimes := readTimes(timesChan)
	splitsData := Splits{
		PB:        pbTimes.formatTimes(),
		SumOfBest: sumOfBestTimes.formatTimes(),
	}
	readyChannel <- true
	handleChannel <- splitsData
}

func splitsHandler(w http.ResponseWriter, r *http.Request) {
	<-readyChannel
	data := <-handleChannel
	tmpl, _ := template.ParseFiles("splits.html")
	tmpl.Execute(w, data)
}
