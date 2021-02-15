package utils

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func loadRepositoryHead(repository *git.Repository) (headRepository *plumbing.Reference) {
	headRepository, err := repository.Head()
	if err != nil {
		panic(err)
	}

	return
}

func TagLastCommit(repository *git.Repository, tag *string) (tagReference *plumbing.Reference) {
	headRepository := loadRepositoryHead(repository)

	tagReference, err := repository.CreateTag(*tag, headRepository.Hash(), &git.CreateTagOptions{
		Message: *tag,
	})
	if err != nil {
		panic(err)
	}

	return
}
