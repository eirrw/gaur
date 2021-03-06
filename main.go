package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
	"virunus.com/gaur/cmd"
)

const AppName = "gaur"

func main() {
	app := cli.App{
		Name:                   AppName,
		HelpName:               "",
		Usage:                  "",
		UsageText:              "",
		ArgsUsage:              "",
		Version:                "",
		Description:            "",
		Commands:               []*cli.Command{
			cmd.CmdUpdate,
			cmd.CmdList,
			cmd.CmdClean,
			cmd.CmdInit,
		},
		Flags:                  nil,
		EnableBashCompletion:   false,
		HideHelp:               false,
		HideHelpCommand:        false,
		HideVersion:            false,
		BashComplete:           nil,
		Before:                 nil,
		After:                  nil,
		Action:                 nil,
		CommandNotFound:        nil,
		OnUsageError:           nil,
		Compiled:               time.Time{},
		Authors:                nil,
		Copyright:              "",
		Reader:                 nil,
		Writer:                 nil,
		ErrWriter:              nil,
		ExitErrHandler:         nil,
		Metadata:               nil,
		ExtraInfo:              nil,
		CustomAppHelpTemplate:  "",
		UseShortOptionHandling: false,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
