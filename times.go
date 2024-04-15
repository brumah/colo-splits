package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Times struct {
	wave1          int
	wave2          int
	wave3          int
	wave4          int
	wave5          int
	wave6          int
	wave7          int
	wave8          int
	wave9          int
	wave10         int
	wave11         int
	wave12         int
	total          int
	completedWaves int
}

type FormattedTimes struct {
	Wave1  string
	Wave2  string
	Wave3  string
	Wave4  string
	Wave5  string
	Wave6  string
	Wave7  string
	Wave8  string
	Wave9  string
	Wave10 string
	Wave11 string
	Wave12 string
	Total  string
}

type Splits struct {
	PB        FormattedTimes
	SumOfBest FormattedTimes
}

type timesChannel chan Times

func iterateWaves(file *os.File) Times {
	scanner := bufio.NewScanner(file)

	var times Times
	var completedWaves int
	var total int
	for scanner.Scan() {
		currentWave := strings.ReplaceAll(strings.Split(scanner.Text(), ":")[0], " ", "")
		waveTime := strings.ReplaceAll(strings.Split(strings.Split(scanner.Text(), ":")[1], "/")[0], " ", "")
		ticks, _ := strconv.Atoi(waveTime)
		times.setWaveTime(currentWave, ticks)
		total += ticks
		completedWaves++
	}
	times.total = total
	times.completedWaves = completedWaves

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
	}
	return times
}

func (t *Times) setWaveTime(currentWave string, waveTime int) {
	switch currentWave {
	case "Wave1":
		t.wave1 = waveTime
	case "Wave2":
		t.wave2 = waveTime
	case "Wave3":
		t.wave3 = waveTime
	case "Wave4":
		t.wave4 = waveTime
	case "Wave5":
		t.wave5 = waveTime
	case "Wave6":
		t.wave6 = waveTime
	case "Wave7":
		t.wave7 = waveTime
	case "Wave8":
		t.wave8 = waveTime
	case "Wave9":
		t.wave9 = waveTime
	case "Wave10":
		t.wave10 = waveTime
	case "Wave11":
		t.wave11 = waveTime
	case "Wave12":
		t.wave12 = waveTime
	}
}

func (t *Times) getBestWaveTime(currentTimes Times) {
	if currentTimes.wave1 < t.wave1 && currentTimes.wave1 != 0 {
		t.wave1 = currentTimes.wave1
	}
	if currentTimes.wave2 < t.wave2 && currentTimes.wave2 != 0 {
		t.wave2 = currentTimes.wave2
	}
	if currentTimes.wave3 < t.wave3 && currentTimes.wave3 != 0 {
		t.wave3 = currentTimes.wave3
	}
	if currentTimes.wave4 < t.wave4 && currentTimes.wave4 != 0 {
		t.wave4 = currentTimes.wave4
	}
	if currentTimes.wave5 < t.wave5 && currentTimes.wave5 != 0 {
		t.wave5 = currentTimes.wave5
	}
	if currentTimes.wave6 < t.wave6 && currentTimes.wave6 != 0 {
		t.wave6 = currentTimes.wave6
	}
	if currentTimes.wave7 < t.wave7 && currentTimes.wave7 != 0 {
		t.wave7 = currentTimes.wave7
	}
	if currentTimes.wave8 < t.wave8 && currentTimes.wave8 != 0 {
		t.wave8 = currentTimes.wave8
	}
	if currentTimes.wave9 < t.wave9 && currentTimes.wave9 != 0 {
		t.wave9 = currentTimes.wave9
	}
	if currentTimes.wave10 < t.wave10 && currentTimes.wave10 != 0 {
		t.wave10 = currentTimes.wave10
	}
	if currentTimes.wave11 < t.wave11 && currentTimes.wave11 != 0 {
		t.wave11 = currentTimes.wave11
	}
	if currentTimes.wave12 < t.wave12 && currentTimes.wave12 != 0 {
		t.wave12 = currentTimes.wave12
	}
	t.total = t.wave1 + t.wave2 + t.wave3 + t.wave4 + t.wave5 + t.wave6 + t.wave7 + t.wave8 + t.wave9 + t.wave10 + t.wave11 + t.wave12
}

func convertTicks(ticks int) string {
	minutes := int((float64(ticks) * 0.6) / 60.0)
	seconds := int(math.Mod(float64(ticks)*0.6, 60.0))
	remainderTick := math.Mod(float64(ticks)*0.6, 60.0) - math.Floor(math.Mod(float64(ticks)*0.6, 60.0))
	remainderTick = math.Round(remainderTick * 10)
	return fmt.Sprintf("%02d:%02d.%v0\n", minutes, seconds, int(remainderTick))
}

func (t Times) formatTimes() FormattedTimes {
	return FormattedTimes{
		Wave1:  convertTicks(t.wave1),
		Wave2:  convertTicks(t.wave2),
		Wave3:  convertTicks(t.wave3),
		Wave4:  convertTicks(t.wave4),
		Wave5:  convertTicks(t.wave5),
		Wave6:  convertTicks(t.wave6),
		Wave7:  convertTicks(t.wave7),
		Wave8:  convertTicks(t.wave8),
		Wave9:  convertTicks(t.wave9),
		Wave10: convertTicks(t.wave10),
		Wave11: convertTicks(t.wave11),
		Wave12: convertTicks(t.wave12),
		Total:  convertTicks(t.total),
	}
}

func readTimes(timesChan timesChannel) (Times, Times) {
	pbTimes := Times{
		total: 100000,
	}
	sumOfBestTimes := Times{
		wave1:  100000,
		wave2:  100000,
		wave3:  100000,
		wave4:  100000,
		wave5:  100000,
		wave6:  100000,
		wave7:  100000,
		wave8:  100000,
		wave9:  100000,
		wave10: 100000,
		wave11: 100000,
		wave12: 100000,
	}
	for times := range timesChan {
		if times.completedWaves == 12 && times.total < pbTimes.total {
			pbTimes = times
		}
		sumOfBestTimes.getBestWaveTime(times)
	}
	return pbTimes, sumOfBestTimes
}
