package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/urfave/cli"
)

// Message is something given by user
type Message []string

func main() {
	app := cli.NewApp()
	app.Name = "conohasay"
	app.Description = app.Name + " is a program that generates ASCII picture of ConoHa characters"
	app.Version = "1.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Hironobu Saito",
			Email: "hiro@hironobu.org",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "character,c",
			Usage: "Specifies a particular character to use. You can also use -l option to get the list of characters",
			Value: "conoha",
		},
		cli.StringFlag{
			Name:  "size,s",
			Usage: `Specifies the size of the picture. The value should be "s", "m" or "l".`,
			Value: "s",
		},
		cli.BoolFlag{
			Name:  "force-vertical,f",
			Usage: "Force the vertical layout",
		},

		cli.IntFlag{
			Name:  "wrapcolumn,W",
			Usage: "Specifies roughly where the message should be wrapped",
			Value: 80,
		},
		cli.BoolFlag{
			Name:  "l,list",
			Usage: "List characters",
		},
	}

	// Setup the help writer
	cli.AppHelpTemplate = `
{{.Name}} version {{.Version}}
Usage: {{.Name}} [-flv] [-h] [-c name]
[-s size(s,m or l)] [-W wrapcolumn] [message]
`
	originalHelpPrinter := cli.HelpPrinter
	cli.HelpPrinter = func(w io.Writer, templ string, d interface{}) {
		app := d.(*cli.App)
		buf := bytes.NewBuffer(make([]byte, 0, len(templ)*2))
		originalHelpPrinter(buf, templ, app)

		cow, err := NewCow("conoha", "s")
		if err != nil {
			fmt.Fprintf(os.Stdout, buf.String())
			return
		}

		msg := strings.Split(buf.String(), "\n")
		output, err := conohasay(cow, msg, 50)
		if err != nil {
			fmt.Fprintf(os.Stdout, buf.String())
			return
		}
		fmt.Fprintf(os.Stdout, output)
	}

	// Run
	app.Action = action
	app.RunAndExitOnError()
}

func action(ctx *cli.Context) error {
	if ctx.Bool("list") {
		cows := listCows()
		for _, cow := range cows {
			fmt.Fprintf(os.Stdout, "%s\n", cow)
		}
		return nil
	}

	name := ctx.String("character")
	size := ctx.String("size")
	if size != "s" && size != "m" && size != "l" {
		return fmt.Errorf(`Parameter size should be "l", "m" or "s".\n`)
	}

	cow, err := NewCow(name, size)
	if err != nil {
		return err
	}

	wrapcolumn := ctx.Int("wrapcolumn")
	message := scanMessage(ctx)
	output, err := conohasay(cow, message, wrapcolumn)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, output)
	return nil
}

func scanMessage(ctx *cli.Context) Message {
	var input string

	if ctx.NArg() > 0 {
		input = strings.Join(ctx.Args(), " ")

	} else {
		for {
			sc := bufio.NewScanner(os.Stdin)
			sc.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
				if atEOF && len(data) == 0 {
					return 0, nil, nil
				}
				return len(data), data, nil
			})

			sc.Scan()
			input = strings.Trim(sc.Text(), "\r\n")
			if input != "" {
				break
			}
		}
	}

	input = strings.Trim(input, "\r\n \t")
	return strings.Split(input, "\n")
}
