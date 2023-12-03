package day03

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"github.com/timconinx/AoC2023/util"
)

var symbols map[util.Coordinate]string = make(map[util.Coordinate]string)
var linenr int
var sum int

var gears map[util.Coordinate][]int = make(map[util.Coordinate][]int)

func ScanSymbols(line string) {
	linenr++
	lo.ForEach(strings.Split(line, ""), func(c string, i int) {
		if !regexp.MustCompile(`[\d\.]`).MatchString(c) {
			symbols[util.NewCoordinate(linenr, i+1)] = c
		}
	})
}

func ScanNumbers(line string) {
	linenr++
	var currentnumber int
	var startindex int
	lo.ForEach(strings.Split(line, ""), func(c string, i int) {
		n, err := strconv.Atoi(c)
		if err == nil {
			if currentnumber > 0 {
				currentnumber = 10*currentnumber + n
			} else {
				startindex = i
				currentnumber = n
			}
		} else {
			if currentnumber > 0 {
				if nextToSymbol(currentnumber, linenr, startindex, i+1) {
					//println(fmt.Sprintf("%v next to symbol", currentnumber))
					sum += currentnumber
				}
			}
			currentnumber = 0
		}
	})
	if currentnumber > 0 {
		if nextToSymbol(currentnumber, linenr, startindex, len(line)-1) {
			//println(fmt.Sprintf("%v next to symbol", currentnumber))
			sum += currentnumber
		}
	}
}

func nextToSymbol(number int, x int, start int, end int) bool {
	for i := start; i <= end; i++ {
		up := util.NewCoordinate(x-1, i)
		down := util.NewCoordinate(x+1, i)
		if symbols[up] != "" || symbols[down] != "" {
			if symbols[up] == "*" {
				gears[up] = append(gears[up], number)
			}
			if symbols[down] == "*" {
				gears[down] = append(gears[down], number)
			}
			return true
		}
	}
	left := util.NewCoordinate(x, start)
	right := util.NewCoordinate(x, end)
	if symbols[left] != "" || symbols[right] != "" {
		if symbols[left] == "*" {
			gears[left] = append(gears[left], number)
		}
		if symbols[right] == "*" {
			gears[right] = append(gears[right], number)
		}
		return true
	}
	return false
}

func calculateGearRatio() int {
	return lo.Sum[int](
		lo.Reduce[[]int](
			lo.Values[util.Coordinate, []int](gears),
			func(agg []int, item []int, i int) []int {
				if len(item) == 2 {
					return append(agg, item[0]*item[1])
				}
				return agg
			}, []int{}),
	)
}

func Day03(name string, dorun bool) {
	if dorun {
		util.ProcessFile("../day03/"+name+".txt", ScanSymbols)
		linenr = 0
		util.ProcessFile("../day03/"+name+".txt", ScanNumbers)
		println(sum)
		println(calculateGearRatio())
	}
}
