package crdts

import (
	"testing"
)

func TestAddingToOrSets(t *testing.T) {
	orset := OrSet{}
	sum := orset.Add(1)

	if sum.Value() != 1 {
		t.Fatalf("sum value is not 1.")
	}
}

func TestRemovingingToOrSets(t *testing.T) {
	orset := OrSet{}
	rem := orset.Remove(1)

	if rem.Value() != -1 {
		t.Fatalf("remove value is not -1")
	}
}

func TestMergingOrSets(t *testing.T) {
	orset1 := (OrSet{}).Add(1)
	orset2 := (OrSet{}).Remove(2)

	merged12 := orset1.Merge(orset2)
	merged21 := orset2.Merge(orset1)

	if merged12.Value() != -1 {
		t.Fatalf("merged12.Value is not -1, but %v", merged12.Value())
	}

	if merged21.Value() != -1 {
		t.Fatalf("merged21.Value is not -1, but %v", merged21.Value())
	}
}
