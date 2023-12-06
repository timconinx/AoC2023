package util

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func ProcessFile(filename string, processfunc func(string)) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		processfunc(reader.Text())
	}
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Pow(base int, exp int) int {
	if exp < 0 {
		return 0
	}
	pow := 1
	for i := 0; i < exp; i++ {
		pow *= base
	}
	return pow
}

func SplitIntArray(line string) []int {
	sarray := strings.Split(line, " ")
	return lo.Map(sarray, func(s string, _ int) int {
		return Atoi(s)
	})
}

func StartsWithDigit(line string) bool {
	return regexp.MustCompile(`^\d`).MatchString(line)
}
