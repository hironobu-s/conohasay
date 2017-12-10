package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var newline = "\n"

func main() {
	app := cli.NewApp()

	app.Name = "conohasay"
	app.Description = app.Name
	app.Flags = []cli.Flag{
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
		cli.BoolFlag{
			Name:  "list,l",
			Usage: "List all charaster names.",
		},
	}
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

	aa, width, err := loadCow(ctx)
	if err != nil {
		return err
	}

	message := scanMessage(ctx)
	message = conohasay(message, width/2)

	fmt.Fprintf(os.Stdout, message+aa)
	return nil
}

func scanMessage(ctx *cli.Context) string {
	var input string

	// body
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
			input = strings.Trim(sc.Text(), newline)
			if input != "" {
				break
			}
		}
	}
	return strings.Trim(input, "\r\n \t")
}
