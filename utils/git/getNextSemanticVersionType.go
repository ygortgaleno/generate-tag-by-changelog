package utils

import (
	"fmt"
	"strings"

	cmdUtils "github.com/ygortgaleno/generate-tag-by-changelog/v1/utils/cmd"
	regexUtils "github.com/ygortgaleno/generate-tag-by-changelog/v1/utils/regex"
)

func getTypesOfLastCommitsSinceLastTag(currentTag *string) string {
	commandString := fmt.Sprintf("git log %s..@ --oneline", *currentTag)
	output := cmdUtils.RunCommand(&commandString)

	cleanedOutput := strings.ReplaceAll(output, "\n", " ")
	cleanedOutput = strings.ToLower(cleanedOutput)

	return cleanedOutput
}

func classifyNextVersionByCommits(commitsSinceLastTag *string) string {
	Breaking, Feat, Fix := "breaking", "feat", "fix|refactor|chore"
	if regexUtils.MatchString(&Breaking, commitsSinceLastTag) {
		return "major"
	} else if regexUtils.MatchString(&Feat, commitsSinceLastTag) {
		return "minor"
	} else if regexUtils.MatchString(&Fix, commitsSinceLastTag) {
		return "patch"
	}

	return ""
}

func GetNextSemanticVersionType(currentTag *string) string {
	commitsSinceLastTag := getTypesOfLastCommitsSinceLastTag(currentTag)
	return classifyNextVersionByCommits(&commitsSinceLastTag)
}
