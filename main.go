package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
	"github.com/urfave/cli"
)

var newline = "\n"

func main() {
	app := cli.NewApp()

	app.Name = "conohasay"
	app.Description = app.Name
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "cowfile,f",
			Usage: "Load a charaster picture from a file.",
		},
		cli.StringFlag{
			Name:  "mikumo,m",
			Usage: "Specifies a particular character to use.",
			Value: "conoha",
		},

		cli.StringFlag{
			Name:  "size,s",
			Usage: "Specifies a size of picture.",
			Value: "s",
		},
	}
	app.Action = action
	app.RunAndExitOnError()
}

func action(ctx *cli.Context) error {
	aa, width, err := loadCow(ctx)
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

func loadCow(ctx *cli.Context) (aa string, width int, err error) {
	name := ctx.String("mikumo")
	if name != "conoha" && name != "anzu" && name != "logo" {
		return aa, width, fmt.Errorf("Undefined character name. [%s]", name)
	}

	size := ctx.String("size")
	if size != "s" && size != "m" && size != "l" {
		return aa, width, fmt.Errorf("Undefined image size. [%s]", name)
	}

	//pp.Printf("%v\n", Assets.Files)
	file := name + "-" + size + ".cow"
	f, err := Assets.Open(file)
	if err != nil {
		return aa, width, fmt.Errorf("Could not load the character. [%s]", file)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return aa, width, err
	}

	aa = string(data)
	width = strings.Index(aa, "\n")

	return aa, width, nil
}
