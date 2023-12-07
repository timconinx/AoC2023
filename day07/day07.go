package day07

import (
	"fmt"
	"sort"
	"strings"

	"github.com/samber/lo"
	"github.com/timconinx/AoC2023/util"
)

const (
	order     string = "J23456789TQKA"
	highcard  int    = 0
	onepair   int    = 1
	twopair   int    = 2
	threeook  int    = 3
	fullhouse int    = 4
	fourook   int    = 5
	fiveook   int    = 6
)

var games map[string]int = make(map[string]int)

func ProcessLine(line string) {
	parts := strings.Split(line, " ")
	games[parts[0]] = util.Atoi(parts[1])
}

func getTotalWinnings() int {
	values := make(map[string]int)
	lo.ForEach(strings.Split(order, ""), func(item string, index int) {
		values[item] = index
	})
	hands := lo.Keys(games)
	sort.Slice(hands, func(i, j int) bool {
		cati := cat(hands[i])
		catj := cat(hands[j])
		if cati == catj {
			handi := strings.Split(hands[i], "")
			handj := strings.Split(hands[j], "")
			for k := 0; k < 5; k++ {
				vletteri := values[handi[k]]
				vletterj := values[handj[k]]
				if vletteri != vletterj {
					return vletteri < vletterj
				}
			}
			panic(fmt.Sprintf("we can't reach here! %v == %v", hands[i], hands[j]))
		} else {
			return cati < catj
		}
	})
	var total int
	lo.ForEach(hands, func(item string, index int) {
		total += games[item] * (index + 1)
	})
	return total
}

func cat(hand string) int {
	cards := strings.Split(hand, "")
	occ := lo.CountValues(cards)
	noj := occ["J"]
	switch len(occ) {
	case 5:
		if noj == 0 {
			return highcard
		} else {
			return onepair
		}
	case 4:
		switch noj {
		case 0:
			return onepair
		case 1:
			return threeook
		case 2:
			return threeook
		default:
			panic("")
		}
	case 3:
		if occ[cards[0]] == 2 || occ[cards[1]] == 2 || occ[cards[2]] == 2 {
			switch noj {
			case 0:
				return twopair
			case 1:
				return fullhouse
			case 2:
				return fourook
			default:
				panic("")
			}
		} else {
			switch noj {
			case 0:
				return threeook
			case 1:
				return fourook
			case 3:
				return fourook
			default:
				panic("")
			}
		}
	case 2:
		if occ[cards[0]] == 2 || occ[cards[0]] == 3 {
			switch noj {
			case 0:
				return fullhouse
			case 2:
				return fiveook
			case 3:
				return fiveook
			default:
				panic("")
			}
		} else {
			switch noj {
			case 0:
				return fourook
			case 1:
				return fiveook
			case 4:
				return fiveook
			default:
				panic("")
			}
		}
	case 1:
		return fiveook
	default:
		panic("this can't happen!")
	}
}

func Day07(name string, dorun bool) {
	if dorun {
		util.ProcessFile("../day07/"+name+".txt", ProcessLine)
		println(getTotalWinnings())
	}
}
