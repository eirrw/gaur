package util

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"sort"
	"virunus.com/gaur/cfg"
)

func GitCountCommits(r *git.Repository, ref *plumbing.Reference) (int, error) {
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return 0, err
	}

	var cCount int
	cIter.ForEach(func(commit *object.Commit) error {
		cCount++

		return nil
	})
	if err != nil {
		return 0, err
	}

	return cCount, nil
}

func GetAurDirs(alpha bool) ([]string, error) {
	config, err := cfg.GetConfig()
	if err != nil {
		return nil, err
	}

	// open the CWD
	d, err := os.Open(*config.CacheDir)
	if err != nil {
		return nil, err
	}

	// Read all entries in the CWD
	dirs, err := d.ReadDir(0)
	if err != nil {
		return nil, err
	}

	var dirNames []string
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		dirNames = append(dirNames, dir.Name())
	}

	if alpha {
		sort.Strings(dirNames)
	}

	return dirNames, nil
}
