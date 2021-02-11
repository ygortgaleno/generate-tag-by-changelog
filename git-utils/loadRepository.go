package gitUtils

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

//	Function that loads repository of current directory
func LoadRepository() *git.Repository {
	repository, err := git.PlainOpen(os.Getenv("PWD"))
	if err != nil {
		fmt.Println("Error on load git repository: ", err)
		os.Exit(1)
	}

	return repository
}
