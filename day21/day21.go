package day21

import (
	"github.com/samber/lo"
	. "github.com/timconinx/AoC2023/util"
)

var rocks = make(map[Coordinate]bool)
var y int
var start Coordinate
var maxx, maxy int

func ProcessLine(line string) {
	maxx = len(line)
	for x := 0; x < len(line); x++ {
		char := string(line[x])
		switch char {
		case "#":
			rocks[NewCoordinate(x, y)] = true
		case "S":
			start = NewCoordinate(x, y)
		}
	}
	y++
	maxy = y
}

func possibleStepsWith(nr int, poss []Coordinate) int {
	if nr == 0 {
		return len(poss)
	}
	newposs := []Coordinate{}
	for _, p := range poss {
		north := p.North()
		if !rocks[safeCoordinate(north)] {
			newposs = safeAdd(north, newposs)
		}
		south := p.South()
		if !rocks[safeCoordinate(south)] {
			newposs = safeAdd(south, newposs)
		}
		east := p.East()
		if !rocks[safeCoordinate(east)] {
			newposs = safeAdd(east, newposs)
		}
		west := p.West()
		if !rocks[safeCoordinate(west)] {
			newposs = safeAdd(west, newposs)
		}
	}
	return possibleStepsWith(nr-1, newposs)
}

func safeCoordinate(c Coordinate) Coordinate {
	return NewCoordinate(Mod(c.X(), maxx), Mod(c.Y(), maxy))
}

func safeAdd(c Coordinate, cs []Coordinate) []Coordinate {
	if !lo.ContainsBy(cs, func(item Coordinate) bool {
		return item.X() == c.X() && item.Y() == c.Y()
	}) {
		return append(cs, c)
	}
	return cs
}

func Day21(name string, dorun bool) {
	if dorun {
		ProcessFile("../day21/"+name+".txt", ProcessLine)
		println(possibleStepsWith(5000, []Coordinate{start}))
	}
}
