package gitUtils

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

// Function that return the current tag name
func GetCurrentTag(repository *git.Repository) (currentTag string) {
	tagsObjects, errTagsObjects := repository.TagObjects()
	if errTagsObjects != nil {
		fmt.Println("Error on load tags objects: ", errTagsObjects)
		os.Exit(1)
	}

	tagObject, errTagObject := tagsObjects.Next()
	if errTagObject != nil {
		fmt.Println("Error on load tag object: ", errTagObject)
		os.Exit(1)
	}

	currentTag = tagObject.Name

	return
}
