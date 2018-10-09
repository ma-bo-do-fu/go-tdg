package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"errors"
	"strconv"
	"path/filepath"
	"strings"
)

var errCommandHelp = fmt.Errorf("command help shown")

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		commandTxt,
	}

	err := app.Run(os.Args)
	if err != nil {
		if err != errCommandHelp {
			fmt.Errorf("error", "%s", err)
		}
		fmt.Println(err)
		os.Exit(1)
	}
}

func validateValue(input string) (float64, error) {
	unit := input[len(input)-1:]

	bytes := 0.0
	inputString := ""
	if _, err := strconv.ParseFloat(unit, 32); err == nil {
		bytes = 1.0
		inputString = input
	} else {
		inputString = input[:len(input)-1]
		switch unit {
		case "k", "K":
			bytes = 1024
		case "m", "M":
			bytes = 1024 * 1024
		case "g", "G":
			bytes = 1024 * 1024 * 1024
		default:
			return 0, errors.New("invalid unit")
		}
	}

	num, err := strconv.ParseFloat(inputString, 32)
	if err != nil {
		return 0, errors.New("invalid value")
	}
	value := num * bytes
	return value, nil
}

var commandTxt = cli.Command{
	Name:  "txt",
	Usage: "make .txt file",
	Action: func(c *cli.Context) error {

		if c.NArg() < 2 {
			return errors.New("invalid args")
		}

		arg := c.Args().First()

		value, err := validateValue(arg)
		if err != nil {
			return err
		}

		data := strings.Repeat("a", int(value))

		path, err := filepath.Abs(c.Args().Get(1) + ".txt")
		if err != nil {
			return err
		}

		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()

		file.Write(([]byte)(data))

		fmt.Println("save file to " + path)

		return nil
	},
}
