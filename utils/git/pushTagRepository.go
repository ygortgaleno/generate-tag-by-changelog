package utils

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

func PushTagRepository(repository *git.Repository, tag *string) {
	err := repository.Push(&git.PushOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
		RefSpecs:   []config.RefSpec{config.RefSpec("refs/tags/*:refs/tags/*")},
	})

	if err != nil {
		panic(err)
	}
}
