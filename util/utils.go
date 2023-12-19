package util

import (
	"bufio"
	"fmt"
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

func CutPrefix(s string, p string) string {
	result, found := strings.CutPrefix(s, p)
	if !found {
		panic(fmt.Sprintf("%v prefix not found in %v", p, s))
	}
	return result
}

func CutSuffix(s string, p string) string {
	result, found := strings.CutSuffix(s, p)
	if !found {
		panic(fmt.Sprintf("%v suffix not found in %v", p, s))
	}
	return result
}

func Ggd(x int, y int) int {
	var a, b int
	if x < y {
		a = x
		b = y
	} else {
		a = y
		b = x
	}
	for {
		if b == 0 {
			break
		}
		t := b
		b = a % b
		a = t
	}
	return a
}

func Kgv(x int, y int) int {
	return (x / Ggd(x, y)) * y
}

func Abs(i int) int {
	if i < 0 {
		return 0 - i
	} else {
		return i
	}
}

func StringAsHex(s string) int {
	i := -1
	return lo.ReduceRight(strings.Split(s, ""), func(agg int, b string, _ int) int {
		c, err := strconv.Atoi(b)
		if err != nil {
			switch b {
			case "a":
				c = 10
			case "b":
				c = 11
			case "c":
				c = 12
			case "d":
				c = 13
			case "e":
				c = 14
			case "f":
				c = 15
			default:
				panic("got character " + b)
			}
		}
		i++
		return agg + c*int(Pow(16, i))
	}, 0)
}
