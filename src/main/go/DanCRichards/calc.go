package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type StationData struct {
	min   int
	max   int
	total int
	count int
}

func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	started := time.Now()
	run()
	fmt.Printf("%0.6f", time.Since(started).Seconds())
}

func run() {
	if len(os.Args) < 2 {
		panic("No arguments")
	}
	fileName := os.Args[1]

	// Open File
	file, fileError := os.Open(fileName)
	defer file.Close() // This will close the file after the function has been run.
	if fileError != nil {
		panic(fileError)
	}

	// Start processing the records

	stations := make(map[string]*StationData)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		parts := strings.Split(line, ";")

		stationName := parts[0]
		// To do, change float to int
		temp, parseError := strconv.Atoi(strings.ReplaceAll(parts[1], ".", ""))

		if parseError != nil {
			fmt.Printf("Error parsing float on line, " + line)
			panic(parseError)
		}

		if stations[stationName] == nil {
			stations[stationName] = &StationData{temp, temp, temp, 1}
		} else {
			stations[stationName].count += 1
			stations[stationName].total += temp
			stations[stationName].min = getMin(stations[stationName].min, temp)
			stations[stationName].max = getMax(stations[stationName].min, temp)
		}

	}

	for key, value := range stations {
		average := value.total / value.count
		fmt.Printf("%s=%f/%f/%f\n", key, value.min/10, value.max/10, average/10)
	}

}
