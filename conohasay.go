package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	runewidth "github.com/mattn/go-runewidth"
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

func conohasay(cow *Cow, msg Message, wrapcolumn int) (output string, err error) {
	// calculate wrapsize
	lw := 0 // line width
	for _, line := range msg {
		l := runewidth.StringWidth(line)
		if l > lw {
			lw = l
		}
	}

	if wrapcolumn <= lw {
		tmp := make([]string, 0, len(msg)*2)
		for _, line := range msg {
			wrapped := strings.Split(runewidth.Wrap(line, wrapcolumn), "\n")
			tmp = append(tmp, wrapped...)
		}
		msg = tmp

	} else {
		wrapcolumn = lw
	}

	return wrap2(msg, cow, wrapcolumn), nil
}

func wrap2(msg Message, cow *Cow, wrapcolumn int) string {
	// wrap message
	balloon := balloonText(msg, wrapcolumn, "left")

	// height
	overlap := len(balloon) - cow.Height()/2
	height := len(balloon) + cow.Height()/2

	buf := bytes.NewBuffer(make([]byte, 0, len(strings.Join(msg, ""))+cow.ArtSize))

	mi := 0
	ci := 0
	for i := 0; i < height; i++ {
		if mi >= overlap {
			if ci < len(cow.Art[ci]) {
				buf.WriteString(cow.Art[ci])
				ci++
			}
		} else {
			buf.WriteString(strings.Repeat(" ", cow.Width()))
		}

		if mi < len(balloon) {
			buf.WriteString(balloon[mi])
			mi++
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func balloonText(msg Message, wrapcolumn int, way string) []string {
	l := len(msg)
	buf := make([]string, l+5)
	p := 0

	buf[p] = " " + strings.Repeat("_", wrapcolumn+2) + " "
	p++

	for i, line := range msg {
		line = strings.Trim(line, "\r\n\t")
		line = strings.Replace(line, "\t", "", -1)

		if i == 0 && l == 1 {
			line = "< " + runewidth.FillRight(line, wrapcolumn) + ` >`
		} else if i == 0 {
			line = "/ " + runewidth.FillRight(line, wrapcolumn) + ` \`
		} else if i == l-1 {
			line = `\ ` + runewidth.FillRight(line, wrapcolumn) + " /"
		} else {
			line = "| " + runewidth.FillRight(line, wrapcolumn) + " |"
		}
		buf[p] = line
		p++
	}

	buf[p] = " " + strings.Repeat("-", wrapcolumn+2)
	if way == "left" {
		buf[p+1] = `   /`
		buf[p+2] = `  / `
	} else {
		buf[p+1] = `     \`
		buf[p+2] = `      \`
	}
	buf[p+3] = ``
	return buf
}

func wrap1(msg Message, cow *Cow, wrapcolumn int) string {
	buf := bytes.NewBuffer(make([]byte, 0, len(strings.Join(msg, ""))+cow.ArtSize))

	balloon := balloonText(msg, wrapcolumn, "right")
	buf.WriteString(strings.Join(balloon, "\n"))
	buf.WriteString("\n")
	buf.WriteString(strings.Join(cow.Art, "\n"))

	return buf.String()
}

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

func loadCow(name string, size string) (cow *Cow, err error) {
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
