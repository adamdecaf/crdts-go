package main

import (
	"fmt"

	"github.com/adamdecaf/crdts-go/crdts"
)

func main() {
	orset := crdts.OrSet{}

	// Just add and remove ints
	sum := orset.Add(1).Remove(3).Add(1).Add(2)
	fmt.Printf("orset = %v, sum = %v\n", orset.Value(), sum.Value())
	fmt.Printf("sum after = %v\n", sum)

	// Merge two OrSets
	orset2 := crdts.OrSet{}
	sum2 := orset2.Add(10).Remove(5).Add(-1).Add(6).Remove(2)
	merged := sum.Merge(sum2)
	fmt.Printf("merged = %v, total = %v\n", merged.Value(), merged)
}
