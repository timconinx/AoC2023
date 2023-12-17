package day10

import (
	. "github.com/timconinx/AoC2023/util"
)

type (
	pipe struct {
		painted    bool
		connectsTo []Coordinate
	}
)

var grid = make(map[Coordinate]pipe)
var start Coordinate
var ypos int

func ProcessLine(line string) {
	for x := 0; x < len(line); x++ {
		neighbours := []Coordinate{}
		switch string(line[x]) {
		case "|":
			neighbours = append(neighbours, NewCoordinate(x, ypos-1))
			neighbours = append(neighbours, NewCoordinate(x, ypos+1))
		case "-":
			neighbours = append(neighbours, NewCoordinate(x-1, ypos))
			neighbours = append(neighbours, NewCoordinate(x+1, ypos))
		case "L":
			neighbours = append(neighbours, NewCoordinate(x, ypos-1))
			neighbours = append(neighbours, NewCoordinate(x+1, ypos))
		case "J":
			neighbours = append(neighbours, NewCoordinate(x, ypos-1))
			neighbours = append(neighbours, NewCoordinate(x-1, ypos))
		case "7":
			neighbours = append(neighbours, NewCoordinate(x, ypos+1))
			neighbours = append(neighbours, NewCoordinate(x-1, ypos))
		case "F":
			neighbours = append(neighbours, NewCoordinate(x, ypos+1))
			neighbours = append(neighbours, NewCoordinate(x+1, ypos))
		case "S":
			start = NewCoordinate(x, ypos)
		}
		grid[NewCoordinate(x, ypos)] = pipe{connectsTo: neighbours}
	}
	ypos++
}

func completeStart() {
	neighbours := []Coordinate{}
	startnode := pipe{painted: true}
	for c, p := range grid {
		if len(p.connectsTo) > 0 && (p.connectsTo[0].Equals(start) || p.connectsTo[1].Equals(start)) {
			neighbours = append(neighbours, c)
		}
	}
	if len(neighbours) != 2 {
		panic("premisse of exactly two connecting pipes is false")
	}
	startnode.connectsTo = neighbours
	grid[start] = startnode
}

func maxPaintFlow() int {
	var length int
	streams := []Coordinate{
		grid[start].connectsTo[0],
		grid[start].connectsTo[1],
	}
	for {
		if len(streams) == 0 {
			break
		}
		newstreams := []Coordinate{}
		for _, s := range streams {
			p := grid[s]
			if !p.painted {
				np := pipe{painted: true, connectsTo: p.connectsTo}
				grid[s] = np
				if !grid[p.connectsTo[0]].painted {
					newstreams = append(newstreams, p.connectsTo[0])
				} else if !grid[p.connectsTo[1]].painted {
					newstreams = append(newstreams, p.connectsTo[1])
				}

			}
		}
		streams = newstreams
		length++
	}
	return length
}

func Day10(name string, dorun bool) {
	if dorun {
		ProcessFile("../day10/"+name+".txt", ProcessLine)
		completeStart()
		println(maxPaintFlow())
	}
}
