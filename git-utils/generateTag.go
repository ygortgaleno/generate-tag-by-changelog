package gitUtils

import (
	"fmt"
	"os"

	semver "github.com/blang/semver/v4"
)

type SemanticVersionType string

type ReleaseType string

const (
	Major SemanticVersionType = "major"
	Minor SemanticVersionType = "minor"
	Patch SemanticVersionType = "patch"
)

const (
	Beta             ReleaseType = "beta"
	Alpha            ReleaseType = "alpha"
	ReleaseCandidate ReleaseType = "rc"
	Production       ReleaseType = "production"
)

func GenerateTag(currentTag *string, releaseType *ReleaseType, semanticVersionType *SemanticVersionType) string {
	version, err := semver.Make(*currentTag)
	if err != nil {
		fmt.Println("Error in load tag into semantic version:", err)
		os.Exit(1)
	}

	if *semanticVersionType == Major {
		version.IncrementMajor()
	} else if *semanticVersionType == Minor {
		version.IncrementMinor()
	} else if *semanticVersionType == Patch {
		version.IncrementPatch()
	}

	if *releaseType == Production {
		version.Pre = make([]semver.PRVersion, 0)
	} else {
		version.Pre = make([]semver.PRVersion, 1)
		prVersion, err := semver.NewPRVersion(string(*releaseType))
		if err != nil {
			fmt.Println("Error on create pre release version:", err)
			os.Exit(1)
		}

		version.Pre[0] = prVersion
	}

	return version.String()
}
