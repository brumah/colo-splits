package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type fileChannel chan *os.File

func pushFiles(folderPath string, fileChan fileChannel) {
	defer close(fileChan)
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".txt" {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			fileChan <- file
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
	}
}

func readFiles(fileChan fileChannel, timesChan timesChannel) {
	defer close(timesChan)
	for file := range fileChan {
		defer file.Close()
		timesChan <- iterateWaves(file)
	}
}
