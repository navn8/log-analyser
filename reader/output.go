package reader

import (
	"sort"
)

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
