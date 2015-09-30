package crdts

type State struct {
	Identifier int
	Value int
}

type OrSet struct {
	// Private fields, please don't call directly
	Additions []State
	Removals []State
}

// Note:
// This only works because integers are commutative.

func (o OrSet) Add(inc int) OrSet {
	st := State {
		Identifier: len(o.Additions),
		Value: inc,
	}

	res := OrSet {
		Additions: append(o.Additions, st),
		Removals: o.Removals,
	}

	return res
}

func (o OrSet) Remove(inc int) OrSet {
	st := State {
		Identifier: len(o.Removals),
		Value: inc,
	}

	res := OrSet {
		Additions: o.Additions,
		Removals: append(o.Removals, st),
	}

	return res
}

func (o OrSet) Merge(other OrSet) OrSet {
	st := OrSet {
		Additions: append(o.Additions, other.Additions...),
		Removals: append(o.Removals, other.Removals...),
	}

	return st
}

func (o OrSet) Value() int {
	var sum = 0

	for i := range o.Additions {
		sum += o.Additions[i].Value
	}

	for i := range o.Removals {
		sum -= o.Removals[i].Value
	}

	return sum
}
