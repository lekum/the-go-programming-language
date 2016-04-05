package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Wordcount struct {
	Words  []string
	Counts map[string]int
}

func (wc Wordcount) Len() int {
	return len(wc.Words)
}

func (wc Wordcount) Less(i, j int) bool {
	return wc.Counts[wc.Words[i]] < wc.Counts[wc.Words[j]]
}

func (wc Wordcount) Swap(i, j int) {
	wc.Words[i], wc.Words[j] = wc.Words[j], wc.Words[i]
}

func main() {
	wc := Wordcount{[]string{}, make(map[string]int)}
	var total int
	in := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		wc.Counts[word]++
		total++
	}

	for k, _ := range wc.Counts {
		wc.Words = append(wc.Words, k)
	}

	sort.Sort(sort.Reverse(wc))

	for _, word := range wc.Words {
		fmt.Printf("%s -> %%%d\n", word, int(100*float64(wc.Counts[word])/float64(total)))
	}

}
