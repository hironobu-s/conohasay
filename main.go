package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

const (
	COWS_DIR = "cows"
	MAXWIDTH = 50
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
	message := strings.Join(ctx.Args(), " ")
	aa, width, err := loadCow("conoha")
	if err != nil {
		return err
	}
	_ = width
	message = formatMessage(message)

	fmt.Fprintf(os.Stdout, "%s%s", message, aa)
	return nil
}

func formatMessage(message string) string {
	rb := bytes.NewBufferString(message)
	ob := bytes.NewBufferString("")

	width := MAXWIDTH
	if rb.Len() < width {
		width = rb.Len()
	}

	// header
	ob.WriteString(" ")
	ob.WriteString(strings.Repeat("_", width+2))
	ob.WriteString("  ")
	ob.WriteString(newline)

	// body
	first := true
	for {
		buf := rb.Next(width)

		format := "%-" + strconv.Itoa(width) + "s"
		line := fmt.Sprintf(format, string(buf))

		if first && rb.Len() == 0 {
			// There is only one row in the message.
			ob.WriteString("< ")
			ob.WriteString(line)
			ob.WriteString(` >`)
			ob.WriteString(newline)
			break

		} else if first {
			ob.WriteString("/ ")
			ob.WriteString(line)
			ob.WriteString(` \`)
			ob.WriteString(newline)
			first = false

		} else if rb.Len() == 0 {
			ob.WriteString(`\ `)
			ob.WriteString(line)
			ob.WriteString(" /")
			ob.WriteString(newline)
			break

		} else {
			ob.WriteString("| ")
			ob.WriteString(line)
			ob.WriteString(" |")
			ob.WriteString(newline)
		}
	}

	// footer
	ob.WriteString(" ")
	ob.WriteString(strings.Repeat("-", width+2))
	ob.WriteString(" ")
	ob.WriteString(newline)

	ob.WriteString(`    \`)
	ob.WriteString(newline)
	ob.WriteString(`     \`)
	ob.WriteString(newline)

	return ob.String()
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
