package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/urfave/cli/v2"
	"virunus.com/gaur/cfg"
	"virunus.com/gaur/util"
)

// CmdClean is the main entry for the clean command
var CmdClean = &cli.Command{
	Name:   "clean",
	Usage:  "clean built packages and leftover source files",
	Action: clean,
}

func clean(_ *cli.Context) error {
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

	err = w.Clean(&git.CleanOptions{Dir: true})
	if err != nil {
		return err
	}

	s, err := w.Status()
	if err != nil {
		return err
	}

	fmt.Println(path)
	fmt.Println(s.IsUntracked(path))

	return nil
}

func cleanIgnoredFiles(path string) error {
	_, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	return nil
}
