package day18

import (
	"regexp"
	"strings"

	. "github.com/timconinx/AoC2023/util"
)

var curx, cury int = 1, 1

var figure = []Coordinate{NewCoordinate(1, 1)}

func ProcessLine(line string) {
	parts := strings.Split(line, " ")
	dir := parts[0]
	length := Atoi(parts[1])
	adapt(length, dir)
	figure = append([]Coordinate{NewCoordinate(curx, cury)}, figure...)
}

func adapt(length int, dir string) {
	switch dir {
	case "U", "3":
		cury -= length
	case "D", "1":
		cury += length
	case "L", "2":
		curx -= length
	case "R", "0":
		curx += length
	}
}

func ProcessLine2(line string) {
	matches := regexp.MustCompile(`\(#(.+?)\)$`).FindStringSubmatch(line)
	code := matches[1]
	dir := string(code[5])
	length := StringAsHex(code[0:5])
	adapt(length, dir)
	figure = append([]Coordinate{NewCoordinate(curx, cury)}, figure...)
}

func Day18(name string, dorun bool) {
	if dorun {
		ProcessFile("../day18/"+name+".txt", ProcessLine)
		println(AreaOf(figure) + CircumferenceOf(figure)/2 + 1)
		figure = []Coordinate{NewCoordinate(1, 1)}
		curx, cury = 1, 1
		ProcessFile("../day18/"+name+".txt", ProcessLine2)
		println(AreaOf(figure) + CircumferenceOf(figure)/2 + 1)
	}
}
