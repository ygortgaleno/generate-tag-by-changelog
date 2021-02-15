package utils

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

func PushTagRepository(repository *git.Repository, tag *string) {
	err := repository.Push(&git.PushOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
		RefSpecs:   []config.RefSpec{config.RefSpec(fmt.Sprintf("refs/tags/%s:refs/tags/%s", *tag, *tag))},
	})

	if err != nil {
		panic(err)
	}
}
