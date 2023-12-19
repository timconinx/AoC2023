package day19

import (
	"regexp"
	"strings"

	"github.com/samber/lo"
	. "github.com/timconinx/AoC2023/util"
)

type (
	part struct {
		r map[string]int
	}
	test struct {
		r    string
		t    string
		v    int
		dest string
	}
)

var workflows = make(map[string][]test)
var accepted []part
var processingSecondPart bool

func (p part) sumUp() int {
	return lo.Reduce(lo.Values(p.r), func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
}

func (t test) run(p part) bool {
	switch t.t {
	case "<":
		return p.r[t.r] < t.v
	case ">":
		return p.r[t.r] > t.v
	case "true":
		return true
	default:
		panic("unexpected operand " + t.t)
	}
}

func ProcessLine(line string) {
	if line == "" {
		processingSecondPart = true
		return
	}
	if !processingSecondPart {
		wf := regexp.MustCompile(`^(\w+)\{(.+?)\}$`).FindStringSubmatch(line)
		wfparts := strings.Split(wf[2], ",")
		var tests []test
		for i := 0; i < len(wfparts)-1; i++ {
			t := regexp.MustCompile(`^(\w)(.)(\d+):(\w+)$`).FindStringSubmatch(wfparts[i])
			tests = append(tests, test{t[1], t[2], Atoi(t[3]), t[4]})
		}
		tests = append(tests, test{t: "true", dest: wfparts[len(wfparts)-1]})
		workflows[wf[1]] = tests
		return
	}
	line = CutPrefix(line, "{")
	line = CutSuffix(line, "}")
	pparts := strings.Split(line, ",")
	partmap := make(map[string]int)
	for _, pp := range pparts {
		xmas := strings.Split(pp, "=")
		partmap[xmas[0]] = Atoi(xmas[1])
	}
	thepart := part{r: partmap}
	dest := "in"
	for {
		if dest == "A" {
			accepted = append(accepted, thepart)
			break
		}
		if dest == "R" {
			break
		}
		workflow := workflows[dest]
		for i := 0; i < len(workflow); i++ {
			if workflow[i].run(thepart) {
				dest = workflow[i].dest
				break
			}
		}
	}
}

func Day19(name string, dorun bool) {
	if dorun {
		ProcessFile("../day19/"+name+".txt", ProcessLine)
		println(lo.Reduce(accepted, func(agg int, item part, _ int) int {
			return agg + item.sumUp()
		}, 0))
	}
}
