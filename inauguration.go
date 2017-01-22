package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"path/filepath"
	"sort"
	"strings"
)

// Speech describes a file/speech text
type Speech struct {
	Name           string
	Raw            string
	NumWords       int
	NumUniqueWords int
	Tokens         []string
	FL             FreqList

	TFIDF FreqList
}

// Freq is a simple k/v that can be sorted
type Freq struct {
	Key string
	Val float64
}

// FreqList is a simple k/v list that can be sorted
type FreqList []Freq

func (f FreqList) Len() int           { return len(f) }
func (f FreqList) Less(i, j int) bool { return f[i].Val < f[j].Val }
func (f FreqList) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f FreqList) Has(search string) bool {
	for _, item := range f {
		if item.Key == search {
			return true
		}
	}
	return false
}

func handleError(errstring string, e error) {
	if e != nil {
		log.Fatalf("[ERROR] %s: (%v)\n", errstring, e)
	}
}

func tokenize(contents string) []string {
	result := []string{}
	contents = strings.Replace(contents, "\t", " ", -1)
	contents = strings.Replace(contents, "\n", " ", -1)
	contents = strings.Replace(contents, "--", " ", -1)
	contents = strings.Replace(contents, "’", "'", -1)
	contents = strings.Replace(contents, "“", "\"", -1)
	contents = strings.Replace(contents, "”", "\"", -1)

	words := strings.Split(contents, " ")
	for _, w := range words {
		w = strings.Trim(w, " ,.'\")(;-_“…:’”—-?")
		w = strings.ToLower(w)
		if w != "" {
			result = append(result, w)
		}
	}
	return result
}

func getFrequencies(words []string) FreqList {
	result := FreqList{}
	n := len(words) // n == number of words in document
	m := map[string]int{}
	for _, w := range words {
		if val, ok := m[w]; ok {
			m[w] = val + 1
		} else {
			m[w] = 1
		}
	}

	for k, v := range m {
		result = append(result, Freq{k, float64(v) / float64(n)})
	}
	sort.Sort(result)

	return result
}

func tfidf(input *Speech, corpus []*Speech) FreqList {
	rel := FreqList{}
	n := len(corpus) // number of documents
	for _, w := range input.FL {
		nt := 0 // number of documents with the word
		for _, s := range corpus {
			if s.FL.Has(w.Key) {
				nt++ // if w (word) appears in s (speech), add 1 to nt
			}
		}
		idf := math.Log(float64(n / nt))
		tfidf := w.Val * idf
		t := Freq{w.Key, tfidf}
		rel = append(rel, t)
	}
	sort.Sort(sort.Reverse(rel))
	return rel
}

func main() {
	speeches := []*Speech{}
	speechFiles, err := filepath.Glob("./speeches/*.txt")
	handleError("Failed to glob", err)

	for _, filename := range speechFiles {
		temp := &Speech{
			Name: filename,
		}
		data, err := ioutil.ReadFile(filename)
		handleError("Failed to read file", err)
		temp.Raw = string(data)
		temp.Tokens = tokenize(temp.Raw)
		temp.NumWords = len(temp.Tokens)
		temp.FL = getFrequencies(temp.Tokens)
		temp.NumUniqueWords = len(temp.FL)
		speeches = append(speeches, temp)
	}

	for _, s := range speeches {
		t := tfidf(s, speeches)
		s.TFIDF = t
		fmt.Printf("Speech: %s; Words: %d\nTFIDF:\n", s.Name, s.NumWords)
		for i, temp := range s.TFIDF {
			if i > 10 {
				break
			}
			fmt.Printf("\t%s: %f\n", temp.Key, temp.Val)
		}
		fmt.Println()
	}

}
