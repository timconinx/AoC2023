package day01

import (
	"strconv"
	"strings"

	"github.com/samber/lo"

	"github.com/timconinx/AoC2023/util"
)

var sum int

func ProcessLine(line string) {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	chars := strings.Split(line, "")
	word := ""
	digits := lo.FilterMap(chars, func(s string, _ int) (int, bool) {
		d, err := strconv.Atoi(s)
		if err != nil {
			word = word + s
			found := -1
			lo.ForEach(words, func(w string, i int) {
				if strings.Contains(word, w) {
					found = i + 1
				}
			})
			if found < 0 {
				return 0, false
			} else {
				word = word[len(word)-1:]
				return found, true
			}
		}
		word = ""
		return d, true
	})
	if len(digits) > 0 {
		sum += digits[0]*10 + digits[len(digits)-1]
	}
}

func Day01(name string, dorun bool) {
	if dorun {
		util.ProcessFile("../day01/"+name+".txt", ProcessLine)
		println(sum)
	}
}
