package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mattn/go-runewidth"
	"github.com/urfave/cli"
)

const (
	COWS_DIR = "cows"
)

var newline = "\n"

func main() {
	app := cli.NewApp()

	app.Name = "conohasay"
	app.Description = app.Name
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "testflag,t",
			Usage: "test flags",
		},
	}
	app.Action = action
	app.RunAndExitOnError()
}

func action(ctx *cli.Context) error {

	aa, width, err := loadCow("conoha")
	if err != nil {
		return err
	}

	message := readMessage(ctx)
	message = formatMessage(message, width/2)

	fmt.Fprintf(os.Stdout, message+aa)
	return nil
}

func readMessage(ctx *cli.Context) string {
	var input string

	// body
	if ctx.NArg() > 0 {
		input = strings.Join(ctx.Args(), " ")

	} else {
		sc := bufio.NewScanner(os.Stdin)
		sc.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if atEOF && len(data) == 0 {
				return 0, nil, nil
			}
			return len(data), data, nil
		})

		sc.Scan()
		input = sc.Text()
	}
	return strings.Trim(input, "\r\n \t")
}

func formatMessage(input string, maxWrapsize int) string {
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

func loadCow(name string) (aa string, width int, err error) {
	p := filepath.Join(COWS_DIR, name+".cow")
	data, err := ioutil.ReadFile(p)
	if err != nil {
		goto RET
	}

	aa = string(data)
	width = strings.Index(aa, "\n")

RET:
	return aa, width, err
}
