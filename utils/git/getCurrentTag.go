package utils

import (
	"errors"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

func loadAllTags(repository *git.Repository) (allTags storer.ReferenceIter) {
	allTags, err := repository.Tags()
	if err != nil {
		panic(err)
	}

	return
}

func getAnotatedTags(repository *git.Repository, tag *plumbing.Reference) (isAnotated bool, tagName string) {
	isAnotated = false

	tagObject, err := repository.TagObject(tag.Hash())
	if err == nil {
		isAnotated = true
		tagName = tagObject.Name
	}

	return
}

// Function that return the current tag name or throws a panic error if not found.
func GetCurrentTag(repository *git.Repository) (currentTag string) {
	allTags := loadAllTags(repository)

	err := allTags.ForEach(func(tag *plumbing.Reference) error {
		isAnotatedTag, tagName := getAnotatedTags(repository, tag)
		if isAnotatedTag {
			currentTag = tagName
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	if currentTag == "" {
		panic(errors.New("Cannot find any anotated taged in this repository."))
	}

	return
}
