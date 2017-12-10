package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

var newline = "\n"

const (
	DEFAULT_WRAPCOLUMN = 40
)

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
		cli.IntFlag{
			Name:  "wrapcolumn, w",
			Usage: "Specifies roughly where the message should be wrapped. Default is " + strconv.Itoa(DEFAULT_WRAPCOLUMN),
			Value: DEFAULT_WRAPCOLUMN,
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

	name := ctx.String("mikumo")
	size := ctx.String("size")
	if size != "s" && size != "m" && size != "l" {
		return fmt.Errorf(`Parameter size should be "l", "m" or "s".\n`)
	}

	message := scanMessage(ctx)

	cow, err := loadCow(name, size)
	if err != nil {
		return err
	}

	wrapcolumn := ctx.Int("wrapcolumn")
	output, err := conohasay(cow, message, wrapcolumn)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, output)
	return nil
}

func scanMessage(ctx *cli.Context) string {
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
			input = strings.Trim(sc.Text(), newline)
			if input != "" {
				break
			}
		}
	}
	return strings.Trim(input, "\r\n \t")
}
