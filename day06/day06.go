package day06

import (
	"regexp"
	"strconv"

	"github.com/samber/lo"
	"github.com/timconinx/AoC2023/util"
)

var times []int
var distances []int

func ProcessLine(line string) {
	timepattern := regexp.MustCompile(`Time:\s+`)
	distancepattern := regexp.MustCompile(`Distance:\s+`)
	if timepattern.MatchString(line) {
		times = util.SplitIntArray(timepattern.ReplaceAllString(line, ""))
	} else if distancepattern.MatchString(line) {
		distances = util.SplitIntArray(distancepattern.ReplaceAllString(line, ""))
	} else {
		panic("line: " + line)
	}
}

func findProduct() int {
	product := 1
	for i := 0; i < len(times); i++ {
		product *= findNoRaces(times[i], uint64(distances[i]))
	}
	return product
}

func findNoRaces(time int, distance uint64) int {
	firstrace := 0
	lastrace := time
	for {
		firstrace++
		if calculateDistance(time, firstrace) > distance {
			break
		}
	}
	for {
		lastrace--
		if calculateDistance(time, lastrace) > distance {
			break
		}
	}
	return lastrace - firstrace + 1
}

func calculateDistance(totaltime int, presstime int) uint64 {
	tt := uint64(totaltime)
	pt := uint64(presstime)
	return (tt - pt) * pt
}

func combine(arr []int) uint64 {
	return lo.ReduceRight(arr, func(agg uint64, item int, index int) uint64 {
		if index == len(arr)-1 {
			return uint64(item)
		} else {
			return (uint64(item) * util.Pow(10, nodigits(agg))) + agg
		}
	}, uint64(0))
}

func nodigits(i uint64) int {
	return len(strconv.FormatUint(i, 10))
}

func Day06(name string, dorun bool) {
	if dorun {
		util.ProcessFile("../day06/"+name+".txt", ProcessLine)
		//println(findProduct())
		println(findNoRaces(int(combine(times)), combine(distances)))
	}
}
