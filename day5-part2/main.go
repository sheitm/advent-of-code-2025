package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type idRange struct {
	from, to int
}

func (i *idRange) count() int {
	return i.to - i.from + 1
}

func main() {
	var idRanges []*idRange
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}

		arr := strings.Split(line, "-")
		from, _ := strconv.Atoi(arr[0])
		to, _ := strconv.Atoi(arr[1])
		ir := &idRange{from, to}
		idRanges = append(idRanges, ir)
	}

	sort.Slice(idRanges, func(i, j int) bool {
		if idRanges[i].from == idRanges[j].from {
			return idRanges[i].to < idRanges[j].to
		}
		return idRanges[i].from < idRanges[j].from
	})

	//naiveTotal := 0
	//for _, r := range idRanges {
	//	//fmt.Printf("%d-%d-%d\n", r.from, r.to, r.count())
	//	naiveTotal += r.count()
	//}
	//fmt.Println(naiveTotal)

	c := combine(idRanges)

	total := 0
	for _, r := range c {
		total += r.count()
	}
	fmt.Println(total)

	//total, _ := combineAndCount(0, 0, idRanges)
	//fmt.Println(total)

	// 332897640654914 too low!
	// 332897640654914
	// 320341090889471
	// 417568065826953

}

func combine(ranges []*idRange) []*idRange {
	var combined []*idRange
	i := 0
	for {
		if i == len(ranges)-1 {
			combined = append(combined, ranges[i])
			break
		}
		first := ranges[i]
		next := ranges[i+1]
		if first.to < next.from {
			combined = append(combined, first)
			i++
			continue
		}
		next.from = first.from
		i++
	}
	return combined

	//i := 0
	//anyCombinations := false
	//for {
	//	if i == len(ranges)-2 {
	//		break
	//	}
	//
	//	first := ranges[i]
	//	next := ranges[i+1]
	//	if first.to < next.from {
	//		combined = append(combined, first)
	//		combined = append(combined, next)
	//		i += 2
	//		continue
	//	}
	//	anyCombinations = true
	//	c := &idRange{from: first.from, to: next.to}
	//	combined = append(combined, c)
	//	i += 2
	//}
	//
	//if anyCombinations {
	//	return combine(combined)
	//}
	//
	//return combined
}

func combineAndCount(iteration, total int, ranges []*idRange) (int, []*idRange) {
	iteration++

	if len(ranges) == 0 {
		return total, nil
	}
	if len(ranges) == 1 {
		return total + ranges[0].count(), nil
	}

	first := ranges[0]
	next := ranges[1]

	if first.to < next.from {
		fmt.Printf("SIM iteration: %d, increment: %d total: %d len: %d\n", iteration, first.count(), total, len(ranges))
		return combineAndCount(iteration, total+first.count(), ranges[1:])
	}

	combined := &idRange{from: first.from, to: next.to}
	fmt.Printf("COM iteration: %d, increment: %d total: %d len: %d\n", iteration, combined.count(), total, len(ranges))
	return combineAndCount(iteration, total+combined.count(), ranges[2:])

}
