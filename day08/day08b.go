package day08

import (
	"github.com/samber/lo"
	"github.com/timconinx/AoC2023/util"
)

var startnodes []string

func stepsToAllReachZ() []int {
	var allsteps []int
	for _, snode := range startnodes {
		var steps int
		node := nodes[snode]
		for {
			dir := string(directions[steps%len(directions)])
			steps++
			var target string
			switch dir {
			case "L":
				target = node.Left
			case "R":
				target = node.Right
			}
			if string(target[len(target)-1]) == "Z" {
				allsteps = append(allsteps, steps)
				break
			}
			node = nodes[target]
		}
	}
	return allsteps
}

func kgv(n []int) int {
	return lo.Reduce(n, func(agg int, i int, _ int) int {
		return util.Kgv(agg, i)
	}, 1)
}

func Day08b(name string, dorun bool) {
	if dorun {
		util.ProcessFile("../day08/"+name+".txt", ProcessLine)
		println(kgv(stepsToAllReachZ()))
	}
}
