package main

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func Test_combineAndCount(t *testing.T) {

	total, _ := combineAndCount(0, 0, testRanges())
	require.Equal(t, 14, total)
}

func Test_combine(t *testing.T) {
	combined := combine(testRanges())

	total := 0
	for _, r := range combined {
		total += r.count()
	}
	require.Equal(t, 14, total)
}

func testRanges() []*idRange {
	ranges := []*idRange{
		{from: 3, to: 5},
		{from: 10, to: 14},
		{from: 16, to: 20},
		{from: 12, to: 18},
	}

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].from == ranges[j].from {
			return ranges[i].to < ranges[j].to
		}
		return ranges[i].from < ranges[j].from
	})

	return ranges
}
