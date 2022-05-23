package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var Frequency = make(map[string]int)
var AvgProcessingTime = make(map[string]float64)

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
func ParseStatusCode(s string) string {
	parsed := strings.Split(s, ":")
	return parsed[1]
}
func Parse(fi string) {
	f, err := os.Open(fi)
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, e := Readln(r)
	for e == nil {
		fmt.Println(s)
		entry := strings.Split(s, ",")
		//format entry
		url := strings.ToUpper(entry[0])
		processingTime, err := strconv.ParseFloat(entry[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		responseStatus := ParseStatusCode(entry[2])

		//Store the Entry Data - Populate frequency and avg map
		if responseStatus == "200" {
			AvgProcessingTime[url] = (AvgProcessingTime[url]*float64(Frequency[url]) + processingTime) / float64(Frequency[url]+1)
			Frequency[url] = Frequency[url] + 1
		}
		s, e = Readln(r)
	}
}

// data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value float64
}

// slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// function to turn a map into a PairList, then sort and return it.
func RankByFrequency(urlFrequency map[string]float64, size int) PairList {
	pl := make(PairList, len(urlFrequency))
	i := 0
	for k, v := range urlFrequency {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	if size > len(pl) {
		size = len(pl)
	}
	return pl[0:size]
}
