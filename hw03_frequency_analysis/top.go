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

const outputSliceLength = 9

func Top10(input string) []string {
	output := make([]string, 0, 10)
	stringsSlice := prepareString(input)
	sortedSlice := countWords(stringsSlice)
	if len(sortedSlice) > outputSliceLength {
		sortedSlice = sortedSlice[:10]
	}
	for _, val := range sortedSlice {
		output = append(output, val.Key)
	}
	return output
}

func prepareString(inputString string) []string {
	var output []string
	reg := regexp.MustCompile(`[^\w]+$`)
	inputString = strings.ToLower(inputString)
	inputString = strings.TrimSpace(inputString)
	output = strings.Fields(inputString)
	for i, val := range output {
		output[i] = reg.ReplaceAllString(val, "")
	}
	return output
}

func countWords(inputString []string) pairList {
	output := make(map[string]int)
	for _, value := range inputString {
		_, ok := output[value]
		if value == "" {
			continue
		}
		if ok {
			output[value]++
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
