package main

import (
	"fmt"
	"log-analyser/internal/reader"
)

func main() {
	fmt.Println("########### Empty ###########")
	filePath := "/Users/trip/Work/mf/hello/log-analyser/log_empty.txt"
	reader.Parse(filePath)
	findTop := 10
	result := reader.RankByFrequency(reader.AvgProcessingTime, findTop)
	fmt.Printf("%v", result)

	fmt.Println("\n \n ############# Normal ##########")
	filePath = "/Users/trip/Work/mf/hello/log-analyser/log.txt"
	reader.Parse(filePath)
	findTop = 10
	result = reader.RankByFrequency(reader.AvgProcessingTime, findTop)
	fmt.Printf("%v", result)

	fmt.Println("\n \n ############# All Error Status != 200 ##########")
	filePath = "/Users/trip/Work/mf/hello/log-analyser/log_error.txt"
	reader.Parse(filePath)
	findTop = 10
	result = reader.RankByFrequency(reader.AvgProcessingTime, findTop)
	fmt.Printf("%v", result)

	fmt.Println("\n \n ############# Ending with Gif ##########")
	filePath = "/Users/trip/Work/mf/hello/log-analyser/log_gif.txt"
	reader.Parse(filePath)
	findTop = 10
	result = reader.RankByFrequency(reader.AvgProcessingTime, findTop)
	fmt.Printf("%v", result)
}
