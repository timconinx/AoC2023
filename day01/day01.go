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
	file, _ := os.Open(filename)
	defer file.Close()
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		result += processLine(reader.Text())
	}
	return result
}

func processLine(line string) int {
	chars := strings.Split(line, "")
	digits := lo.Filter(chars, func(c string, _ int) bool {
		_, err := strconv.Atoi(c)
		return err == nil
	})
	if len(digits) > 1 {
		return digits[0]*10 + digits[len(digits)-1]
	}
	return 0
}

func Day01() {
	println(processFile("test.txt"))
}
