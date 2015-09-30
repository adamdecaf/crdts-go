package crdts

import (
	"fmt"
)

type OrSet struct {
	Value int

	// Private fields
	Additions []int
	Deletions []int
}

func (o OrSet) Add(other *OrSet) (OrSet, error) {
	if other != nil {
		return o, fmt.Errorf("Empty OrSet passed in during Add")
	}

	// res := OrSet {
	// 	Value:
	// }

	return o, nil
}

// Remove func
