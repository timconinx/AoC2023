package util

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

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

func Pow(base int, exp int) uint64 {
	if exp < 0 {
		return 0
	}
	b := uint64(base)
	pow := uint64(1)
	for i := 0; i < exp; i++ {
		pow *= b
	}
	return pow
}

func SplitIntArray(line string) []int {
	sarray := regexp.MustCompile(`\s+`).Split(line, -1)
	return lo.Map(sarray, func(s string, _ int) int {
		return Atoi(s)
	})
}

func StartsWithDigit(line string) bool {
	return regexp.MustCompile(`^\d`).MatchString(line)
}
