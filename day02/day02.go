package day02

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	"github.com/timconinx/AoC2023/util"
)

var linenr int
var sum int
var powersum int
var totalballs map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func ProcessLine(line string) {
	linenr++
	line, found := strings.CutPrefix(line, fmt.Sprintf("Game %v: ", linenr))
	if !found {
		panic("Game pattern not found")
	}
	games := strings.Split(line, "; ")
	gamevalid := true
	var gamepower map[string]int = make(map[string]int)
	lo.ForEach(games, func(game string, _ int) {
		balls := strings.Split(game, ", ")
		lo.ForEach(balls, func(ball string, _ int) {
			parts := strings.Split(ball, " ")
			number := util.Atoi(parts[0])
			color := parts[1]
			if number > totalballs[color] {
				gamevalid = false
			}
			if number > gamepower[color] {
				gamepower[color] = number
			}
		})
	})
	if gamevalid {
		sum += linenr
	}
	powersum += lo.Reduce[int](lo.Values[string, int](gamepower), func(agg int, item int, _ int) int {
		return agg * item
	}, 1)
}

func Day02(name string, dorun bool) {
	if dorun {
		util.ProcessFile("../day02/"+name+".txt", ProcessLine)
		println(sum)
		println(powersum)
	}
}
