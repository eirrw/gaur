package cmd

import (
	"github.com/go-git/go-git/v5"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"virunus.com/gaur/cfg"
	"virunus.com/gaur/util"
)

var CmdClean = &cli.Command{
	Name:   "clean",
	Usage:  "clean built packages and leftover source files",
	Action: clean,
}

func clean(c *cli.Context) error {
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

		err = gitClean(path)
		if err != nil {
			return err
		}

		err = os.Chdir(path)
		if err != nil {
			return err
		}

		err = os.RemoveAll("src")
		if err != nil {
			return err
		}
	}

	return nil
}

func gitClean(path string) error {
	r, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	err = w.Clean(&git.CleanOptions{Dir: false})
	if err != nil {
		return err
	}

	return nil
}
