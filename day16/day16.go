package day16

import (
	"github.com/samber/lo"
	. "github.com/timconinx/AoC2023/util"
)

type (
	tile struct {
		char    string
		fromdir []string
	}
	beam struct {
		position  Coordinate
		direction string
	}
)

var grid = make(map[Coordinate]tile)
var beams []beam
var ypos, xpos int

func ProcessLine(line string) {
	xpos = len(line)
	for x := 0; x < len(line); x++ {
		grid[NewCoordinate(x, ypos)] = tile{char: string(line[x])}
	}
	ypos++
}

func runBeams() {
	for c, t := range grid {
		grid[c] = tile{char: t.char}
	}
	for {
		if len(beams) == 0 {
			break
		}
		var newbeams []beam
		lo.ForEach(beams, func(b beam, _ int) {
			thistile := grid[b.position]
			inloop := lo.Contains(thistile.fromdir, b.direction)
			if !inloop {
				grid[b.position] = tile{thistile.char, append(thistile.fromdir, b.direction)}
				nbeams := calculateNewBeams(b.position, b.direction, thistile.char)
				lo.ForEach(nbeams, func(nb beam, _ int) {
					if _, ok := grid[nb.position]; ok {
						newbeams = append(newbeams, nb)
					}
				})
			}
		})
		beams = newbeams
	}
}

func calculateNewBeams(coord Coordinate, d string, char string) []beam {
	var nbeams []beam
	switch char {
	case ".":
		nbeams = append(nbeams, newBeam(coord, d))
	case "-":
		switch d {
		case "L", "R":
			nbeams = append(nbeams, newBeam(coord, d))
		case "U", "D":
			nbeams = append(nbeams, newBeam(coord, "L"))
			nbeams = append(nbeams, newBeam(coord, "R"))
		}
	case "|":
		switch d {
		case "U", "D":
			nbeams = append(nbeams, newBeam(coord, d))
		case "L", "R":
			nbeams = append(nbeams, newBeam(coord, "U"))
			nbeams = append(nbeams, newBeam(coord, "D"))
		}
	case "\\":
		switch d {
		case "U":
			nbeams = append(nbeams, newBeam(coord, "L"))
		case "D":
			nbeams = append(nbeams, newBeam(coord, "R"))
		case "L":
			nbeams = append(nbeams, newBeam(coord, "U"))
		case "R":
			nbeams = append(nbeams, newBeam(coord, "D"))
		}
	case "/":
		switch d {
		case "U":
			nbeams = append(nbeams, newBeam(coord, "R"))
		case "D":
			nbeams = append(nbeams, newBeam(coord, "L"))
		case "L":
			nbeams = append(nbeams, newBeam(coord, "D"))
		case "R":
			nbeams = append(nbeams, newBeam(coord, "U"))
		}
	}
	return nbeams
}

func newBeam(c Coordinate, d string) beam {
	switch d {
	case "U":
		return beam{NewCoordinate(c.X(), c.Y()-1), d}
	case "D":
		return beam{NewCoordinate(c.X(), c.Y()+1), d}
	case "R":
		return beam{NewCoordinate(c.X()+1, c.Y()), d}
	case "L":
		return beam{NewCoordinate(c.X()-1, c.Y()), d}
	default:
		panic("Unknown direction " + d)
	}

}

func noEnergized() int {
	return lo.Reduce(lo.Values(grid), func(agg int, item tile, _ int) int {
		if len(item.fromdir) > 0 {
			return agg + 1
		} else {
			return agg
		}
	}, 0)
}

func Day16(name string, dorun bool) {
	if dorun {
		ProcessFile("../day16/"+name+".txt", ProcessLine)
		beams = append(beams, beam{position: NewCoordinate(0, 0), direction: "R"})
		runBeams()
		println(noEnergized())
		var max, noe int
		for x := 0; x < xpos; x++ {
			beams = []beam{beam{NewCoordinate(x, 0), "D"}}
			runBeams()
			noe = noEnergized()
			if noe > max {
				max = noe
			}
			beams = []beam{beam{NewCoordinate(x, ypos-1), "U"}}
			runBeams()
			noe = noEnergized()
			if noe > max {
				max = noe
			}
		}
		for y := 0; y < ypos; y++ {
			beams = []beam{beam{NewCoordinate(0, y), "R"}}
			runBeams()
			noe = noEnergized()
			if noe > max {
				max = noe
			}
			beams = []beam{beam{NewCoordinate(xpos-1, y), "L"}}
			runBeams()
			noe = noEnergized()
			if noe > max {
				max = noe
			}
		}
		println(max)
	}
}
