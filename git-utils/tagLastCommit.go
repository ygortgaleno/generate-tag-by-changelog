package gitUtils

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func TagLastCommit(repository *git.Repository, tag *string) *plumbing.Reference {
	headRepository, err := repository.Head()
	if err != nil {
		fmt.Println("Error in get repository's head:", err)
	}

	ref, err := repository.CreateTag(*tag, headRepository.Hash(), &git.CreateTagOptions{
		Message: *tag,
	})
	if err != nil {
		fmt.Println("Error in tag repository:", err)
		os.Exit(1)
	}

	return ref
}
