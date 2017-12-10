package main

import (
	"sort"
	"strings"
)

func listCows() []string {
	list := make([]string, 0, 5)
	for _, cow := range Assets.Files {
		p := strings.Split(cow.Name(), "-")
		m := false
		for _, l := range list {
			if l == p[0] {
				m = true
			}
		}
		if !m {
			list = append(list, p[0])
		}
	}

	sort.Strings(list)
	return list
}
