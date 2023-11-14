package nummatch

type Exclude func(int64) bool

type NumSeries struct {
	begin   int64
	offset  int64
	exclude Exclude
}

func NewNumSeries(begin, offset int64, exclude Exclude) NumberMatch {
	return &NumSeries{
		begin:   begin,
		offset:  offset,
		exclude: exclude,
	}
}

func (n *NumSeries) RoundDown(target int64) (out int64) {
	// target is not above the start
	if target <= n.begin {
		return target
	}
	// when target is not allign with the series
	if target%n.offset != 0 {
		target = (target / n.offset) * n.offset
	}
	out = target
loop:
	if n.exclude(out) {
		out -= n.offset
		// check the output with in the range
		if out <= n.begin {
			return target
		}
		goto loop
	}
	return out
}

func (n *NumSeries) RoundUp(target int64) (out int64) {
	// target is not above the start
	if target <= n.begin {
		target = n.begin
	}
	// when target is not allign with the series
	if target%n.offset != 0 {
		target = (target / n.offset) * n.offset
	}
	out = target
loop:
	if n.exclude(out) {
		out += n.offset
		goto loop
	}
	return out
}
