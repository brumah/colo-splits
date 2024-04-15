package main

import (
	"html/template"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, "homepage")
}

func splitsHandler(w http.ResponseWriter, r *http.Request) {
	fileChan := make(fileChannel)
	timesChan := make(timesChannel)
	go pushFiles("splits/", fileChan)
	go readFiles(fileChan, timesChan)
	pbTimes, sumOfBestTimes := readTimes(timesChan)
	splitsData := Splits{
		PB:        pbTimes.formatTimes(),
		SumOfBest: sumOfBestTimes.formatTimes(),
	}

	tmpl, _ := template.ParseFiles("splits.html")
	tmpl.Execute(w, splitsData)
}
