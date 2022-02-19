package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/urfave/cli/v2"
	"virunus.com/gaur/cfg"
	"virunus.com/gaur/util"
)

var CmdUpgrade = &cli.Command{
	Name:   "upgrade",
	Usage:  "upgrade a pacakage",
	Action: update,
}

func upgrade(c *cli.Context) error {
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

		err = gitUpdate(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func gitUpdate(path string) error {
	config, err := cfg.GetConfig()
	if err != nil {
		return err
	}

	r, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	result := r.Fetch(&git.FetchOptions{})
	if result != nil && result != git.NoErrAlreadyUpToDate {
		return result
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

	udp := color.GreenString("up to date")
	if headCount != remoteCount {
		udp = color.RedString("needs update")
	}

	fmt.Printf("%-25s %s\n", filepath.Base(path), udp)

	return nil
}
