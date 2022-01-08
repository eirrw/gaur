package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"virunus.com/gaur/cfg"
)

var CmdInit = &cli.Command{
	Name:   "init",
	Usage:  "initialize gaur configuration",
	Action: initialize,
}

func initialize(c *cli.Context) error {
	config, err := cfg.New()
	if err != nil {
		return err
	}

	i, err := cfg.IsInit()
	if err != nil {
		return err
	}

	if i {
		fmt.Print("Config was previously created. Reinitialize? (y/N): ")
		var confirm string

		_, err := fmt.Scanln(&confirm)
		if err != nil && err.Error() != "unexpected newline" {
			return err
		}
		if confirm != "y" && confirm != "Y" {
			fmt.Println("Initialization cancelled")
			return nil
		}
	}

	fmt.Printf("Select cache directory (%s): ", *config.CacheDir)
	_, err = fmt.Scanln(config.CacheDir)
	if err != nil && err.Error() != "unexpected newline" {
		return err
	}

	fmt.Printf("Select default remote (%s): ", *config.Remote)
	_, err = fmt.Scanln(config.Remote)
	if err != nil && err.Error() != "unexpected newline" {
		return err
	}

	err = cfg.Initialize(config)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	fmt.Println("Config initialized")

	return nil
}
