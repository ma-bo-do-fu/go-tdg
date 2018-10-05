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

var commandTxt = cli.Command{
	Name:  "txt",
	Usage: "make .txt file",
	Action: func(c *cli.Context) error {

		if c.NArg() < 2 {
			return errors.New("invalid args")
		}

		arg := c.Args().First()
		unit := arg[len(arg)-1:]

		bytes := 0.0
		num := ""
		if _, err := strconv.ParseFloat(unit,32); err == nil {
			bytes = 1.0
			num = arg
		} else {
			num = arg[:len(arg)-1]
			switch unit {
			case "k", "K":
				bytes = 1024
			case "m", "M":
				bytes = 1024 * 1024
			case "g", "G":
				bytes = 1024 * 1024 * 1024
			default:
				return errors.New("invalid unit")
			}
		}

		s, err := strconv.ParseFloat(num,32)
		if err != nil {
			return errors.New("invalid value")
		}

		count := s * bytes

		fmt.Println("count")
		fmt.Println(count)

		fmt.Println("int count")
		fmt.Println(int(count))

		data := strings.Repeat("a", int(count))

		path,err := filepath.Abs(c.Args().Get(1) + ".txt")
		if err != nil {
			return err
		}

		fmt.Println("path")
		fmt.Println(filepath.Abs(path))

		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()

		file.Write(([]byte)(data))

		return nil
	},
}
