package day04

import (
	"regexp"

	"github.com/samber/lo"
	"github.com/timconinx/AoC2023/util"
)

var score int
var total int
var next []int

func Scorecards(line string) {
	scorepower := winningNumbers(line)
	score += int(util.Pow(2, scorepower-1))
}

func NumberOfCards(line string) {
	nocards := 1
	if len(next) != 0 {
		nocards += next[0]
		next = next[1:]
	}
	total += nocards
	scorepower := winningNumbers(line)
	for i := 0; i < scorepower; i++ {
		if len(next) > i {
			next[i] += nocards
		} else {
			next = append(next, nocards)
		}
	}
}

func winningNumbers(line string) int {
	line = regexp.MustCompile(`Card\s+\d+:\s+`).ReplaceAllString(line, "")
	parts := regexp.MustCompile(`\s+\|\s+`).Split(line, -1)
	swinning := regexp.MustCompile(`\s+`).Split(parts[0], -1)
	snumbers := regexp.MustCompile(`\s+`).Split(parts[1], -1)
	winning := lo.Map(swinning, func(s string, _ int) int {
		return util.Atoi(s)
	})
	numbers := lo.Map(snumbers, func(s string, _ int) int {
		return util.Atoi(s)
	})
	var wnumbers int
	lo.ForEach(numbers, func(item int, index int) {
		if lo.Contains(winning, item) {
			wnumbers++
		}
	})
	return wnumbers
}

func Day04(name string, dorun bool) {
	if dorun {
		util.ProcessFile("../day04/"+name+".txt", Scorecards)
		println(score)
		util.ProcessFile("../day04/"+name+".txt", NumberOfCards)
		println(total)
	}
}
