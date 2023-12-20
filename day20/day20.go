package day20

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	. "github.com/timconinx/AoC2023/util"
)

type (
	pulse struct {
		t bool
	}
	module interface {
		Accept(from string, p pulse) bool
		Act()
		Dests() []string
	}
	standardmodule struct {
		name     string
		dests    []string
		willsend pulse
	}
	broadcaster struct {
		standardmodule
	}
	flipflop struct {
		standardmodule
		state bool
	}
	conjuction struct {
		standardmodule
		inputs map[string]bool
	}
	testmodule struct {
		standardmodule
	}
)

func (p pulse) String() string {
	return fmt.Sprintf("-%v->", p.t)
}

func (p pulse) State(s bool) pulse {
	return pulse{t: s}
}

func (s *standardmodule) Dests() []string {
	return s.dests
}

var agenda []string

func (s *standardmodule) Act() {
	for _, target := range s.dests {
		//		println(fmt.Sprintf("%v %v %v", s.name, s.willsend, target))
		p := newPulse(s.willsend.t)
		if modules[target].Accept(s.name, p) {
			agenda = append(agenda, target)
		}
	}
	for {
		if len(agenda) == 0 {
			break
		}
		item := agenda[0]
		agenda = lo.Drop(agenda, 1)
		modules[item].Act()
	}
}
func (b *broadcaster) Accept(_ string, p pulse) bool {
	b.willsend = p
	return true
}
func (f *flipflop) Accept(_ string, p pulse) bool {
	if !p.t {
		f.state = !f.state
		f.willsend = p.State(f.state)
		return true
	}
	return false
}
func (c *conjuction) Accept(from string, p pulse) bool {

	_, ok := c.inputs[from]
	if !ok {
		panic("key not in conjuction map: " + from)
	}
	c.inputs[from] = p.t
	allhigh := lo.Reduce(lo.Values(c.inputs), func(agg bool, i bool, _ int) bool {
		return agg && i
	}, true)
	c.willsend = p.State(!allhigh)
	if c.name == "vd" && c.willsend.t {
		vdhigh++
	}
	if c.name == "ns" && c.willsend.t {
		nshigh++
	}
	if c.name == "bh" && c.willsend.t {
		bhhigh++
	}
	if c.name == "dl" && c.willsend.t {
		dlhigh++
	}

	return true
}
func (c *conjuction) init() {
	c.inputs = make(map[string]bool)
}
func (c *conjuction) addSource(s string) {
	c.inputs[s] = false
}
func (t testmodule) Accept(_ string, p pulse) bool {
	if t.name == "rx" && !p.t {
		rxlowpulses++
	}
	return false
}

func newPulse(s bool) pulse {
	if s {
		highpulses++
	} else {
		lowpulses++
	}
	return pulse{t: s}
}

var modules = make(map[string]module)
var highpulses, lowpulses int
var rxlowpulses int
var vdhigh, nshigh, bhhigh, dlhigh int

func ProcessLine(line string) {
	parts := strings.Split(line, " -> ")
	src := parts[0]
	dests := strings.Split(parts[1], ", ")
	var m module
	var ok bool
	if src == "broadcaster" {
		m = &broadcaster{standardmodule: standardmodule{name: src, dests: dests}}
	} else if src, ok = strings.CutPrefix(src, "%"); ok {
		m = &flipflop{standardmodule: standardmodule{name: src, dests: dests}}
	} else if src, ok = strings.CutPrefix(src, "&"); ok {
		m = &conjuction{standardmodule: standardmodule{name: src, dests: dests}}
	} else {
		panic("unexpected src " + src)
	}
	modules[src] = m
}

func connectConjuctions() {
	for n := range modules {
		switch c := modules[n].(type) {
		case *conjuction:
			c.init()
			for m := range modules {
				if lo.Contains(modules[m].Dests(), n) {
					c.addSource(m)
				}
			}
		}
	}
}

func testInputsCase() {
	for n := range modules {
		for _, m := range modules[n].Dests() {
			if !lo.Contains(lo.Keys(modules), m) {
				modules[m] = &testmodule{standardmodule: standardmodule{name: m}}
			}
		}
	}
}

func pushTheButton() {
	modules["broadcaster"].Accept("button", newPulse(false))
	modules["broadcaster"].Act()
}

func Day20(name string, dorun bool) {
	if dorun {
		ProcessFile("../day20/"+name+".txt", ProcessLine)
		connectConjuctions()
		testInputsCase()
		/*for i := 0; i < 1000; i++ {
			pushTheButton()
		}
		println(fmt.Sprintf("%v low, %v high = %v", lowpulses, highpulses, lowpulses*highpulses))
		*/
		var nr int
		var bhdlnr, dlnsnr, nsvdnr, vdbhnr int
		for {
			rxlowpulses = 0
			bhhigh = 0
			dlhigh = 0
			nshigh = 0
			vdhigh = 0
			nr++
			pushTheButton()
			if bhdlnr == 0 && bhhigh == 1 { //&& dlhigh == 1 {
				bhdlnr = nr
				println(fmt.Sprintf("bh/dl found at %v", nr))
			}
			if dlnsnr == 0 && dlhigh == 1 { //&& nshigh == 1 {
				dlnsnr = nr
				println(fmt.Sprintf("dl/ns found at %v", nr))
			}
			if nsvdnr == 0 && nshigh == 1 { //&& vdhigh == 1 {
				nsvdnr = nr
				println(fmt.Sprintf("ns/vd found at %v", nr))
			}
			if vdbhnr == 0 && vdhigh == 1 { //&& bhhigh == 1 {
				vdbhnr = nr
				println(fmt.Sprintf("vd/bh found at %v", nr))
			}
			if bhdlnr > 0 && dlnsnr > 0 && nsvdnr > 0 && vdbhnr > 0 {
				break
			}
		}
		println(Kgv(bhdlnr, Kgv(dlnsnr, Kgv(nsvdnr, vdbhnr))))
		println(rxlowpulses)
	}
}
