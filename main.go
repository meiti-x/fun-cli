package main

import (
	"funCLI/commands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "funCLI"
	app.Usage = "this projects contain fun cli tools"

	app.Commands = []*cli.Command{
		{
			Name:        "bmi",
			HelpName:    "",
			Usage:       "funCLI bmi -w 85 --ht 1.88",
			Description: "yet another bmi calculator",
			Action:      commands.BMI,
			Flags: []cli.Flag{
				&cli.UintFlag{
					Name:    "weight",
					Usage:   "enter your weight in Kilogram",
					Aliases: []string{"w"},
				},
				&cli.Float64Flag{
					Name:    "height",
					Usage:   "enter your height in float format. for ex 1.90",
					Aliases: []string{"ht"},
				},
			},
		},
		{
			Name:        "pwdg",
			HelpName:    "",
			Usage:       "funCLI pwdg -l 14",
			Description: "yet another password generator",
			Action:      commands.Pwdg,
			Flags: []cli.Flag{
				&cli.UintFlag{
					Name:    "length",
					Usage:   "enter your password length",
					Aliases: []string{"l"},
				},
			},
		},
	}
	err := app.Run(os.Args)

	if err != nil {
		log.Fatalln(err)
	}
}
