package day11

import (
	"github.com/samber/lo"
	"github.com/timconinx/AoC2023/util"
)

var notemptycolumns []int
var emptycolumns []int
var notemptyrows []int
var emptyrows []int
var galaxy []util.Coordinate
var y, ylength, xlength int

func ProcessLine(line string) {
	xlength = len(line)
	for x, dot := range line {
		if string(dot) == "#" {
			galaxy = append(galaxy, util.NewCoordinate(x, y))
			if !lo.Contains(notemptyrows, y) {
				notemptyrows = append(notemptyrows, y)
			}
			if !lo.Contains(notemptycolumns, x) {
				notemptycolumns = append(notemptycolumns, x)
			}
		}
	}
	y++
}

func expandTheGalaxy() {
	es := 999999
	for x := 0; x < xlength; x++ {
		if !lo.Contains(notemptycolumns, x) {
			emptycolumns = append(emptycolumns, x)
		}
	}
	for y := 0; y < ylength; y++ {
		if !lo.Contains(notemptyrows, y) {
			emptyrows = append(emptyrows, y)
		}
	}
	xlength += (len(emptycolumns) * es)
	ylength += (len(emptyrows) * es)
	for {
		if len(emptycolumns) == 0 {
			break
		}
		x := emptycolumns[0]
		emptycolumns = lo.Drop(emptycolumns, 1)
		expandRight(x, es)
		for i := 0; i < len(emptycolumns); i++ {
			emptycolumns[i] += es
		}
	}
	for {
		if len(emptyrows) == 0 {
			break
		}
		y := emptyrows[0]
		emptyrows = lo.Drop(emptyrows, 1)
		expandDown(y, es)
		for i := 0; i < len(emptyrows); i++ {
			emptyrows[i] += es
		}
	}
}

func expandRight(column int, size int) {
	for i := 0; i < len(galaxy); i++ {
		if galaxy[i].X() > column {
			galaxy[i] = util.NewCoordinate(galaxy[i].X()+size, galaxy[i].Y())
		}
	}
}

func expandDown(row int, size int) {
	for i := 0; i < len(galaxy); i++ {
		if galaxy[i].Y() > row {
			galaxy[i] = util.NewCoordinate(galaxy[i].X(), galaxy[i].Y()+size)
		}
	}
}

func sumOfDistances() int {
	var sum int
	for i := 0; i < len(galaxy)-1; i++ {
		for j := i + 1; j < len(galaxy); j++ {
			sum += galaxy[i].MDistanceTo(galaxy[j])
		}
	}
	return sum
}

func printGalaxy() {
	for y := 0; y < ylength; y++ {
		for x := 0; x < xlength; x++ {
			if galaxyContains(util.NewCoordinate(x, y)) {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
	println()
}

func galaxyContains(c util.Coordinate) bool {
	return lo.ContainsBy(galaxy, func(item util.Coordinate) bool {
		return c.X() == item.X() && c.Y() == item.Y()
	})
}

func Day11(name string, dorun bool) {
	if dorun {
		util.ProcessFile("../day11/"+name+".txt", ProcessLine)
		ylength = y
		//printGalaxy()
		expandTheGalaxy()
		//printGalaxy()
		println(sumOfDistances())
	}
}
