package day05

import (
	"strings"

	"github.com/samber/lo"
	"github.com/timconinx/AoC2023/util"
)

var readingstate = "seeds"
var seeds []int
var sts [][]int
var stf [][]int
var ftw [][]int
var wtl [][]int
var ltt [][]int
var tth [][]int
var htl [][]int

func ProcessLine(line string) {
	// I am not proud
	switch readingstate {
	case "seeds":
		if line, found := strings.CutPrefix(line, "seeds: "); found {
			seeds = util.SplitIntArray(line)
		} else if line == "" {
			readingstate = "sts"
		} else {
			panic("wrong state: was seeds but got " + line)
		}
	case "sts":
		if strings.HasPrefix(line, "seed-to-soil") {
			// do nothing
		} else if util.StartsWithDigit(line) {
			sts = append(sts, util.SplitIntArray(line))
		} else if line == "" {
			readingstate = "stf"
		} else {
			panic("wrong state: was sts but got " + line)
		}
	case "stf":
		if strings.HasPrefix(line, "soil-to-fertilizer") {
			// do nothing
		} else if util.StartsWithDigit(line) {
			stf = append(stf, util.SplitIntArray(line))
		} else if line == "" {
			readingstate = "ftw"
		} else {
			panic("wrong state: was stf but got " + line)
		}
	case "ftw":
		if strings.HasPrefix(line, "fertilizer-to-water") {
			// do nothing
		} else if util.StartsWithDigit(line) {
			ftw = append(ftw, util.SplitIntArray(line))
		} else if line == "" {
			readingstate = "wtl"
		} else {
			panic("wrong state: was ftw but got " + line)
		}
	case "wtl":
		if strings.HasPrefix(line, "water-to-light") {
			// do nothing
		} else if util.StartsWithDigit(line) {
			wtl = append(wtl, util.SplitIntArray(line))
		} else if line == "" {
			readingstate = "ltt"
		} else {
			panic("wrong state: was wtl but got " + line)
		}
	case "ltt":
		if strings.HasPrefix(line, "light-to-temperature") {
			// do nothing
		} else if util.StartsWithDigit(line) {
			ltt = append(ltt, util.SplitIntArray(line))
		} else if line == "" {
			readingstate = "tth"
		} else {
			panic("wrong state: was ltt but got " + line)
		}
	case "tth":
		if strings.HasPrefix(line, "temperature-to-humidity") {
			// do nothing
		} else if util.StartsWithDigit(line) {
			tth = append(tth, util.SplitIntArray(line))
		} else if line == "" {
			readingstate = "htl"
		} else {
			panic("wrong state: was tth but got " + line)
		}
	case "htl":
		if strings.HasPrefix(line, "humidity-to-location") {
			// do nothing
		} else if util.StartsWithDigit(line) {
			htl = append(htl, util.SplitIntArray(line))
		} else if line == "" {
			readingstate = "end"
		} else {
			panic("wrong state: was htl but got " + line)
		}
	default:
		panic("wrong state " + readingstate + " on " + line)
	}
}

func calculateMinimumLocation() int {
	var lowest int = int(^uint(0) >> 1)
	lo.ForEach(seeds, func(seed int, _ int) {
		location := moveWithAllTheRanges(seed)
		if location < lowest {
			lowest = location
		}
	})
	return lowest
}

func calculateMinimumLocationWithRange() int {
	var lowest int = int(^uint(0) >> 1)
	locations := make(chan int, 10)
	defer close(locations)
	var totalthreads int
	for {
		totalthreads++
		go func(seed int, length int) {
			locations <- moveBetterWithAllTheRanges(seed, seed+length)
		}(seeds[0], seeds[1])
		if len(seeds) > 2 {
			seeds = seeds[2:]
		} else {
			break
		}
	}
	for i := 0; i < totalthreads; i++ {
		location := <-locations
		if location < lowest {
			lowest = location
		}
	}
	return lowest
}

func moveWithAllTheRanges(s int) int {
	lo.ForEach([][][]int{sts, stf, ftw, wtl, ltt, tth, htl}, func(item [][]int, _ int) {
		s = moveWithRange(s, item)
	})
	return s
}

func moveWithRange(s int, rr [][]int) int {
	for _, r := range rr {
		if s >= r[1] && s <= r[1]+r[2] {
			return r[0] + s - r[1]
		}
	}
	return s
}

func moveBetterWithAllTheRanges(s int, e int) int {
	stack := [][]int{{s, e}}
	var newstack [][]int
	lo.ForEach([][][]int{sts, stf, ftw, wtl, ltt, tth, htl}, func(item [][]int, _ int) {
		newstack = nil
		for {
			block := stack[0]
			stack = lo.Drop(stack, 1)
			s, e := moveWithRangeBetter(block[0], block[1], item)
			if e > block[1] {
				e = block[1]
			}
			length := e - block[0]
			newstack = append(newstack, []int{s, s + length})
			if e < block[1] {
				stack = append(stack, []int{e + 1, block[1]})
			}
			if len(stack) == 0 {
				stack = newstack
				break
			}
		}
	})
	var result = int(^uint(0) >> 1)
	lo.ForEach(stack, func(item []int, _ int) {
		if item[0] < result {
			result = item[0]
		}
	})
	return result
}

// return (destination, last in range to match)
func moveWithRangeBetter(s int, e int, rr [][]int) (int, int) {
	for _, r := range rr {
		if s >= r[1] && s <= r[1]+r[2] {
			return r[0] + s - r[1], r[1] + r[2]
		}
	}
	return s, s
}

func Day05(name string, dorun bool) {
	if dorun {
		util.ProcessFile("../day05/"+name+".txt", ProcessLine)
		println(calculateMinimumLocation())
		println(calculateMinimumLocationWithRange())
	}
}
