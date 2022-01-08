package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/urfave/cli/v2"
	"path/filepath"
	"virunus.com/gaur/cfg"
	"virunus.com/gaur/util"
)

var CmdList = &cli.Command{
	Name:   "list",
	Usage:  "list packages managed by gaur",
	Action: list,
}

func list(c *cli.Context) error {
	config, err := cfg.GetConfig()
	if err != nil {
		return err
	}

	dirs, err := util.GetAurDirs(true)
	if err != nil {
		return err
	}

	for _, d := range dirs {
		path := filepath.Join(*config.CacheDir, d)

		err = gitList(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func gitList(path string) error {
	config, err := cfg.GetConfig()
	if err != nil {
		return err
	}

	r, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	ref, err := r.Head()
	if err != nil {
		return err
	}

	headCount, err := util.GitCountCommits(r, ref)
	if err != nil {
		return err
	}

	ref, err = r.Reference(plumbing.NewRemoteHEADReferenceName(*config.Remote), true)
	if err != nil {
		return err
	}

	remoteCount, err := util.GitCountCommits(r, ref)
	if err != nil {
		return err
	}

	if headCount != remoteCount {
		fmt.Printf("%-25s %s\n", filepath.Base(path), color.RedString("needs update"))
	}

	return nil
}
