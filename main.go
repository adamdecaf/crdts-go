package main

import (
	"fmt"

	"github.com/adamdecaf/crdts-go/crdts"
)

func main() {
	orset := crdts.OrSet {}

	fmt.Println("crdt examples starting")

	merged := orset.Add(1).Remove(3).Add(1).Add(2)
	fmt.Printf("orset = %v, merged = %v\n", orset.Value(), merged.Value())
	fmt.Printf("merged after = %v\n", merged)

	fmt.Println("crdt examples complete")
}
