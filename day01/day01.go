package day01

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func processFile(filename string) int {
	result := 0
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		result += processLine(reader.Text())
	}
	return result
}

func processLine(line string) int {
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
		return digits[0]*10 + digits[len(digits)-1]
	}
	return 0
}

func Day01(name string) {
	println(processFile("../day01/" + name + ".txt"))
}
