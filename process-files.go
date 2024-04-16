package main

import (
	"fmt"
	"mime/multipart"
)

type fileChannel chan multipart.File

func pushFiles(files []*multipart.FileHeader, testChan fileChannel) {
	defer close(testChan)
	for _, fileHeader := range files {
		if fileHeader != nil {
			file, err := fileHeader.Open()
			if err != nil {
				fmt.Println("Error opening the file:", err)
				continue
			}
			testChan <- file
		}
	}
}

func readFiles(testChan fileChannel, timesChan timesChannel) {
	defer close(timesChan)
	for file := range testChan {
		defer file.Close()
		timesChan <- iterateWaves(file)
	}
}
