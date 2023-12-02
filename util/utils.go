package util

import (
	"bufio"
	"os"
	"strconv"
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
