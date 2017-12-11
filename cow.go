package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Cow struct {
	Name    string
	Size    string
	Art     []string // TerminalArt
	ArtSize int

	// cache
	width int
}

// Width return the length of the terminal art
func (c *Cow) Width() int {
	if c.width == 0 {
		for _, line := range c.Art {
			// "[38" is ansi escape command(set forground color).
			cc := strings.Count(line, "[38") + strings.Count(line, " ")
			if c.width < cc {
				c.width = cc
			}
		}
	}
	return c.width
}

// Height return the count of lines
func (c *Cow) Height() int {
	return len(c.Art)
}

// NewCow creates the struct of Cow
func NewCow(name string, size string) (cow *Cow, err error) {
	file := name + "-" + size + ".cow"
	f, err := Assets.Open(file)
	if err != nil {
		return nil, fmt.Errorf("Could not load the character. [name=%s, size=%s]", name, size)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	cow = &Cow{
		Name:    name,
		Size:    size,
		Art:     strings.Split(string(data), "\n"),
		ArtSize: len(data),
	}

	return cow, nil
}

func ListCows() []string {
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
