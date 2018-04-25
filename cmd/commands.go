package cmd

import "github.com/urfave/cli"

// Commands ...
func Commands() []cli.Command {
	return []cli.Command{
		{
			Name:    "translate",
			Aliases: []string{"t"},
			Usage:   "bai du translate",
			Action:  translator,
		},
	}
}

func Command() cli.Command {
	return cli.Command{
		Name:    "translate",
		Aliases: []string{"t"},
		Usage:   "bai du translate",
		Action:  translator,
	}
}

func Action(c *cli.Context) {
	translator(c)
}
