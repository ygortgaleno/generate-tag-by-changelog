package utils

import (
	"os"

	"github.com/go-git/go-git/v5"
)

//	Function that loads repository of current directory
func LoadRepository() *git.Repository {
	repository, err := git.PlainOpen(os.Getenv("PWD"))
	if err != nil {
		panic(err)
	}

	return repository
}
