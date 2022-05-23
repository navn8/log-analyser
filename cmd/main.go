package main

import (
	"fmt"
	reader "log-analyser"
)

func main() {
	fmt.Println("yo")
	filePath := "/Users/trip/Work/mf/hello/log-analyser/log.txt"
	reader.Parse(filePath)
	findTop := 10
	result := reader.RankByFrequency(reader.AvgProcessingTime, findTop)
	fmt.Printf("%v", result)
}
