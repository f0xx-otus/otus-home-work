package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

type pair struct {
	Key   string
	Value int
}

type pairList []pair

func Top10(input string) []string {
	var output []string
	stringsSlice, err := prepareString(input)
	if err != nil {
		return output
	}
	sortedSlice := countWords(stringsSlice)
	for i, item := range sortedSlice {
		output = append(output, item.Key)
		if i == 9 {
			break
		}
	}
	return output
}

func prepareString(inputString string) ([]string, error) {
	var output []string
	reg := regexp.MustCompile("[^\\wа-яА-Я0-9]+")
	inputString = strings.ToLower(inputString)
	inputString = strings.TrimSpace(inputString)
	output = strings.Fields(inputString)
	for i, val := range output {
		output[i] = reg.ReplaceAllString(val, "")
	}
	return output, nil
}

func countWords(inputString []string) pairList {
	output := make(map[string]int)
	for _, value := range inputString {
		_, ok := output[value]
		if value == "" {
			continue
		}
		if ok {
			output[value] += 1
		} else {
			output[value] = 1
		}
	}
	return rankByWordCount(output)
}

func rankByWordCount(wordsAndCounts map[string]int) pairList {
	pairList := make(pairList, len(wordsAndCounts))
	i := 0
	for k, v := range wordsAndCounts {
		pairList[i] = pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pairList))
	return pairList
}

func (p pairList) Len() int           { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
