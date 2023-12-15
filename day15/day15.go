package day15

import (
	"regexp"
	"strings"

	"github.com/samber/lo"
	"github.com/timconinx/AoC2023/util"
)

type lens struct {
	Label string
	Focus int
}

var sum int
var boxes = make(map[int][]lens)

func ProcessLine(line string) {
	parts := strings.Split(line, ",")
	for _, part := range parts {
		sum += hashcode(part)
		pieces := regexp.MustCompile(`^(\w+)([=-])(\d*)$`).FindStringSubmatch(part)
		label := pieces[1]
		op := pieces[2]
		hc := hashcode(label)
		box := boxes[hc]
		switch op {
		case "-":
			boxes[hc] = removeFromBox(box, label)
		case "=":
			focus := util.Atoi(pieces[3])
			boxes[hc] = replaceOrAppendBox(box, label, focus)
		default:
			panic("wrong operand in part " + part)
		}
	}
}

func replaceOrAppendBox(box []lens, label string, focus int) []lens {
	for i := 0; i < len(box); i++ {
		if box[i].Label == label {
			box[i].Focus = focus
			return box
		}
	}
	return append(box, lens{Label: label, Focus: focus})
}

func removeFromBox(box []lens, label string) []lens {
	return lo.Filter(box, func(item lens, _ int) bool {
		return item.Label != label
	})
}

func hashcode(s string) int {
	return lo.Reduce([]byte(s), func(agg int, b byte, _ int) int {
		agg += int(b)
		agg *= 17
		agg %= 256
		return agg
	}, 0)
}

func totalValue() int {
	var value int
	for i, b := range boxes {
		for j, l := range b {
			value += (i + 1) * (j + 1) * l.Focus
		}
	}
	return value
}

func Day15(name string, dorun bool) {
	if dorun {
		for i := 0; i < 256; i++ {
			boxes[i] = []lens{}
		}
		util.ProcessFile("../day15/"+name+".txt", ProcessLine)
		println(sum)
		println(totalValue())
	}
}
