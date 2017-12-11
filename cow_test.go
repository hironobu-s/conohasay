package main

import (
	"sort"
	"testing"
)

func TestListCows(t *testing.T) {
	names := []string{"conoha", "anzu", "umemiya", "logo"}
	cows := ListCows()

	sort.Strings(names)
	sort.Strings(cows)

	if len(names) != len(cows) {
		t.Errorf("Number of cows is not match the defined size.")
	}

	for i := 0; i < len(names); i++ {
		if names[i] != cows[i] {
			t.Errorf("[%s] is undefined or Extra cow?", cows[i])
		}
	}
}

func TestNewCow(t *testing.T) {
	names := []string{"conoha", "anzu", "umemiya", "logo"}
	sizes := []string{"s", "m", "l"}

	for i := 0; i < len(names); i++ {
		for j := 0; j < len(sizes); j++ {
			cow, err := NewCow(names[i], sizes[j])
			if err != nil {
				t.Error(err)

			} else if cow.Name != names[i] {
				t.Errorf("cow has the wrong name.")
			} else if cow.Size != sizes[j] {
				t.Errorf("cow has the wrong size.")
			}
		}
	}
}
