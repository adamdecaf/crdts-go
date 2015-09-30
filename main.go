package main

import (
	"fmt"

	"github.com/adamdecaf/crdts-go/crdts"
)

func main() {
	orset1 := crdts.OrSet { Value: 1 }
	orset2 := crdts.OrSet { Value: 2 }

	fmt.Println("crdt examples starting")

	merged, err := orset1.Add(&orset2)
	if err != nil {
		fmt.Printf("orset1 = %v, orset2 = %v, add = %v\n", orset1, orset2, merged)
	} else {
		fmt.Printf("Error while adding orset1 = %v, orset2 = %v\n", orset1, orset2)
	}

	fmt.Println("crdt examples complete")
}
