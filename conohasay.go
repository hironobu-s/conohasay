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

func conohasay(cow *Cow, input string, wrapcolumn int) (output string, err error) {
	rows := strings.Split(input, "\n")

	// calculate wrapsize
	lw := 0 // line width
	for _, line := range rows {
		l := runewidth.StringWidth(line)
		if l > lw {
			lw = l
		}
	}

	if wrapcolumn <= lw {
		tmp := make([]string, 0, len(rows)*2)
		for _, line := range rows {
			wrapped := strings.Split(runewidth.Wrap(line, wrapcolumn), "\n")
			tmp = append(tmp, wrapped...)
		}
		rows = tmp

	} else {
		wrapcolumn = lw
	}

	// Wrap the message
	buf := bytes.NewBuffer(make([]byte, 0, len(input)*2))
	buf.WriteString(" " + strings.Repeat("_", wrapcolumn+2) + " " + newline)

	l := len(rows)
	for i, line := range rows {
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
		buf.WriteString(line + "\n")
	}
	buf.WriteString(" " + strings.Repeat("-", wrapcolumn+2) + " " + newline)

	// Append the terminal art after the message
	buf.WriteString(`    \  ` + "\n")
	for i, line := range strings.Split(cow.Art, "\n") {
		if i == 0 {
			buf.WriteString(`     \ ` + line + "\n")
		} else {
			buf.WriteString(`       ` + line + "\n")
		}
	}

	return buf.String(), nil
}

type Cow struct {
	Name string
	Size string
	Art  string //TerminalArt
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
		Name: name,
		Size: size,
		Art:  string(data),
	}

	return cow, nil
}
