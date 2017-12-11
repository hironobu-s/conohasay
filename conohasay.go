package main

import (
	"bytes"
	"strings"

	runewidth "github.com/mattn/go-runewidth"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	horizontalOverlap = 3
)

// Conohasay generates and outputs the ballooned message with ASCII picture.
func Conohasay(cow *Cow, msg Message, wrapcolumn int) (output string, err error) {
	// Calculate wrapsize
	if wrapcolumn <= 0 {
		// Default wrapcolumn is detected from the terminal width.
		// if it can't detect, we use 40 as default size.
		wrapcolumn, _, err = terminal.GetSize(1)
		if err != nil {
			wrapcolumn = 40
		}
	}

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

	w, _, _ := terminal.GetSize(1)
	if w > wrapcolumn+cow.Width() {
		return formatH(msg, cow, wrapcolumn), nil
	}
	return formatV(msg, cow, wrapcolumn), nil
}

// formatV generates the output with the vertical layout.
func formatV(msg Message, cow *Cow, wrapcolumn int) string {
	balloon := balloonText(msg, wrapcolumn, "right")

	buf := bytes.NewBuffer(make([]byte, 0, len(strings.Join(msg, ""))+cow.ArtSize))
	buf.WriteString(strings.Join(balloon, "\n"))
	buf.WriteString("\n")
	buf.WriteString(strings.Join(cow.Art, "\n"))

	return buf.String()
}

// formatH generates the output with the horizontal layout.
func formatH(msg Message, cow *Cow, wrapcolumn int) string {
	balloon := balloonText(msg, wrapcolumn, "left")

	overlap := len(balloon) - cow.Height() + horizontalOverlap
	buf := bytes.NewBuffer(make([]byte, 0, len(strings.Join(msg, ""))+cow.ArtSize))
	mi := 0
	ci := 0
	for ci < len(cow.Art)-1 {
		if mi >= overlap {
			buf.WriteString(cow.Art[ci])
			ci++
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
