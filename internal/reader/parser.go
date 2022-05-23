package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var Frequency = make(map[string]int)
var AvgProcessingTime = make(map[string]float64)

func Parse(fi string) {
	f, err := os.Open(fi)
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, e := Readln(r)
	for e == nil {
		//fmt.Println(s)
		entry := strings.Split(s, ",")
		//format entry
		url := strings.ToUpper(entry[0])

		processingTime, err := strconv.ParseFloat(entry[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		responseStatus := ParseStatusCode(entry[2])

		//Store the Entry Data - Populate frequency and avg map
		if CheckValid(url, responseStatus) {
			AvgProcessingTime[url] = (AvgProcessingTime[url]*float64(Frequency[url]) + processingTime) / float64(Frequency[url]+1)
			Frequency[url] = Frequency[url] + 1
		}
		s, e = Readln(r)
	}
}
