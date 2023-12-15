package day09

import (
	"github.com/samber/lo"
	"github.com/timconinx/AoC2023/util"
)

var endings int
var previouses int

func ProcessLine(line string) {
	history := util.SplitIntArray(line)
	next, previous := nextAndPreviousValue(history)
	endings += history[len(history)-1] + next
	previouses += history[0] - previous
}

func nextAndPreviousValue(row []int) (int, int) {
	newrow := make([]int, len(row)-1)
	for i := 0; i < len(newrow); i++ {
		newrow[i] = row[i+1] - row[i]
	}
	if allZero(newrow) {
		return 0, 0
	} else {
		next, previous := nextAndPreviousValue(newrow)
		return newrow[len(newrow)-1] + next, newrow[0] - previous
	}
}

func allZero(row []int) bool {
	return lo.Reduce(row, func(agg bool, item int, _ int) bool {
		return agg && (item == 0)
	}, true)
}

func Day09(name string, dorun bool) {
	if dorun {
		util.ProcessFile("../day09/"+name+".txt", ProcessLine)
		println(endings)
		println(previouses)
	}
}
