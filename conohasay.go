package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	runewidth "github.com/mattn/go-runewidth"
	"github.com/urfave/cli"
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

func conohasay(input string, maxWrapsize int) string {
	rows := strings.Split(input, "\n")

	// calculate wrapsize
	wrapsize := 0
	for _, line := range rows {
		s := runewidth.StringWidth(line)
		if s > maxWrapsize {
			wrapsize = maxWrapsize
			break
		} else if s > wrapsize {
			wrapsize = s
		}
	}

	if wrapsize == maxWrapsize {
		tmp := make([]string, 0, len(rows)*2)
		for _, line := range rows {
			wrapped := strings.Split(runewidth.Wrap(line, maxWrapsize), "\n")
			tmp = append(tmp, wrapped...)
		}
		rows = tmp
	}

	buf := bytes.NewBuffer(make([]byte, 0, len(input)*2))

	// header
	buf.WriteString(" " + strings.Repeat("_", wrapsize+2) + " " + newline)

	// message
	l := len(rows)
	for i, line := range rows {
		line = strings.Trim(line, "\r\n\t")
		line = strings.Replace(line, "\t", "", -1)

		if i == 0 && l == 1 {
			line = "< " + runewidth.FillRight(line, wrapsize) + ` >`
		} else if i == 0 {
			line = "/ " + runewidth.FillRight(line, wrapsize) + ` \`
		} else if i == l-1 {
			line = `\ ` + runewidth.FillRight(line, wrapsize) + " /"
		} else {
			line = "| " + runewidth.FillRight(line, wrapsize) + " |"
		}
		buf.WriteString(line + "\n")
	}

	// footer
	buf.WriteString(" " + strings.Repeat("-", wrapsize+2) + " " + newline)
	buf.WriteString(`    \` + newline)
	buf.WriteString(`     \` + newline)

	return buf.String()
}

func loadCow(ctx *cli.Context) (aa string, width int, err error) {
	name := ctx.String("mikumo")
	if name != "conoha" && name != "anzu" && name != "umemiya" && name != "logo" {
		return aa, width, fmt.Errorf("Undefined character name. [%s]", name)
	}

	size := ctx.String("size")
	if size != "s" && size != "m" && size != "l" {
		return aa, width, fmt.Errorf("Undefined image size. [%s]", name)
	}

	file := name + "-" + size + ".cow"
	f, err := Assets.Open(file)
	if err != nil {
		return aa, width, fmt.Errorf("Could not load the character. [name=%s, size=%s]", name, size)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return aa, width, err
	}

	aa = string(data)
	width = strings.Index(aa, "\n")

	return aa, width, nil
}
