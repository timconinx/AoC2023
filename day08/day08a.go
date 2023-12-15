package day08

import (
	"strings"

	"github.com/timconinx/AoC2023/util"
)

type node struct {
	Left  string
	Right string
}

var nodes = make(map[string]node)
var directions string

func ProcessLine(line string) {
	if line == "" {
		return
	}
	if strings.Contains(line, "=") {
		parts := strings.Split(line, " = ")
		nodename := parts[0]
		if string(nodename[len(nodename)-1]) == "A" {
			startnodes = append(startnodes, nodename)
		}
		dirs := strings.Split(parts[1], ", ")
		left := util.CutPrefix(dirs[0], "(")
		right := util.CutSuffix(dirs[1], ")")
		nodes[nodename] = node{Left: left, Right: right}
		return
	}
	directions = line
}

func stepsToReachZ() int {
	var steps int
	node := nodes["AAA"]
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
		if target == "ZZZ" {
			break
		}
		node = nodes[target]
	}
	return steps
}

func Day08a(name string, dorun bool) {
	if dorun {
		util.ProcessFile("../day08/"+name+".txt", ProcessLine)
		println(stepsToReachZ())
	}
}
