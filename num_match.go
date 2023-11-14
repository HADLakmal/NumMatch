package nummatch

import (
	"sort"
)

type NumberMatch interface {
	RoundDown(target int64) (out int64)
	RoundUp(target int64) (out int64)
}

type NumMatch struct {
	begin   int64
	offset  int64
	exclude []int64
}

func NewNumMatchSeries(begin, offset int64, exclude []int64) NumberMatch {
	sort.Slice(exclude, func(i, j int) bool { return exclude[i] < exclude[j] })

	var excludeIndex int
	for index, excludeValue := range exclude {
		if begin > excludeValue {
			excludeIndex = index + 1
		} else {
			break
		}
	}
	return &NumMatch{
		begin:   begin,
		offset:  offset,
		exclude: exclude[excludeIndex:],
	}
}

// RoundDown given value round down near to the begining
// return the same value as if couldn't find value in series
func (n *NumMatch) RoundDown(target int64) (out int64) {
	// target is not above the start
	if target <= n.begin {
		return target
	}
	// when target is not allign with the series
	if target%n.offset != 0 {
		target = (target / n.offset) * n.offset
	}
	// serch the traget value in exclude list
	i := sort.Search(len(n.exclude), func(i int) bool {
		return n.exclude[i] >= target
	})
	out = target

	// target value found in exclude list
	if len(n.exclude) > i && n.exclude[i] == target {
	loop:
		// target value found in exclude list
		if len(n.exclude) > i && i >= 0 {
			out -= n.offset
			// check the output with in the range
			if out <= n.begin {
				return target
			}
			if i > 0 {
				// check output fell into exclude list
				if n.exclude[i-1] == out {
					i -= 1
					goto loop
				}
			}
		}
	}
	return out
}

func (n *NumMatch) RoundUp(target int64) (out int64) {
	// target is not above the start
	if target <= n.begin {
		target = n.begin
	}
	// when target is not allign with the series
	if target%n.offset != 0 {
		target = (target / n.offset) * n.offset
	}
	// serch the traget value in exclude list
	i := sort.Search(len(n.exclude), func(i int) bool {
		return n.exclude[i] >= target
	})
	out = target

	// target value found in exclude list
	if len(n.exclude) > i && n.exclude[i] == target {
	loop:
		if len(n.exclude) > i && i >= 0 {
			out += n.offset
			// check the output with in the range
			if out <= n.begin {
				return n.begin
			}
			if len(n.exclude) > i {
				// check output fell into exclude list
				if n.exclude[i+1] == out {
					i += 1
					goto loop
				}
			}
		}
	}

	return out
}
