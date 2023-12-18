package day18

import (
	"sort"
	"strings"

	"github.com/samber/lo"
	. "github.com/timconinx/AoC2023/util"
)

type (
	hole struct {
		dug   bool
		dir   string
		color string
	}
)

var curx, cury int = 1, 1
var maxx, maxy int = -10000, -10000
var minx, miny int = 10000, 10000

var grid = make(map[Coordinate]hole)
var xindex = make(map[int][]Coordinate)
var yindex = make(map[int][]Coordinate)
var laceholes = []Coordinate{NewCoordinate(1, 1)}

func ProcessLine(line string) {
	parts := strings.Split(line, " ")
	dir := parts[0]
	length := Atoi(parts[1])
	col := parts[2]
	var l int
	switch dir {
	case "U":
		for {
			grid[NewCoordinate(curx, cury-l)] = hole{true, dir, col}
			l++
			if l == length {
				cury -= length
				break
			}
		}
	case "D":
		for {
			grid[NewCoordinate(curx, cury+l)] = hole{true, dir, col}
			l++
			if l == length {
				cury += length
				break
			}
		}
	case "L":
		for {
			grid[NewCoordinate(curx-l, cury)] = hole{true, dir, col}
			l++
			if l == length {
				curx -= length
				break
			}
		}
	case "R":
		for {
			grid[NewCoordinate(curx+l, cury)] = hole{true, dir, col}
			l++
			if l == length {
				curx += length
				break
			}
		}
	}
	if curx > maxx {
		maxx = curx
	}
	if curx < minx {
		minx = curx
	}
	if cury > maxy {
		maxy = cury
	}
	if cury < miny {
		miny = cury
	}
	laceholes = append([]Coordinate{NewCoordinate(curx, cury)}, laceholes...)
}

func IndexAll() {
	lo.ForEach(lo.Keys(grid), func(c Coordinate, _ int) {
		cs := yindex[c.Y()]
		cs = append(cs, c)
		sort.Slice(cs, func(i, j int) bool {
			return cs[i].Y() < cs[j].Y()
		})
		yindex[c.Y()] = cs

		cs = xindex[c.X()]
		cs = append(cs, c)
		sort.Slice(cs, func(i, j int) bool {
			return cs[i].X() < cs[j].X()
		})
		yindex[c.X()] = cs
	})
}

func DigFurther() {
	for y := miny; y <= maxy; y++ {

	}
}

func shoelaces() int {
	var sum int
	for i := 0; i < len(laceholes)-1; i++ {
		sum = sum + (laceholes[i].X() * laceholes[i+1].Y()) - (laceholes[i].Y() * laceholes[i+1].X())
	}
	return sum / 2
}

func abs(i int) int {
	if i < 0 {
		return 0 - i
	} else {
		return i
	}
}

func Day18(name string, dorun bool) {
	if dorun {
		ProcessFile("../day18/"+name+".txt", ProcessLine)
		//IndexAll()
		//DigFurther()
		linelength := len(lo.Values(grid))
		println(abs(shoelaces()) + linelength/2 + 1)
	}
}
