package utils

import (
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
	Alpha            ReleaseType = "alpha"
	Beta             ReleaseType = "beta"
	ReleaseCandidate ReleaseType = "rc"
	Production       ReleaseType = "production"
)

func makeSemanticVersion(currentTag *string) (version semver.Version) {
	version, err := semver.Make(*currentTag)
	if err != nil {
		panic(err)
	}

	return
}

func incrementVersion(version *semver.Version, semanticVersionType *SemanticVersionType) {
	switch *semanticVersionType {
	case Major:
		err := version.IncrementMajor()
		if err != nil {
			panic(err)
		}
	case Minor:
		err := version.IncrementMinor()
		if err != nil {
			panic(err)
		}
	case Patch:
		err := version.IncrementPatch()
		if err != nil {
			panic(err)
		}
	}
}

func setSemanticVersionReleaseType(version *semver.Version, releaseType *ReleaseType) {
	switch *releaseType {
	case Production:
		version.Pre = make([]semver.PRVersion, 0)
	case Alpha, Beta, ReleaseCandidate:
		version.Pre = make([]semver.PRVersion, 1)

		prVersion, err := semver.NewPRVersion(string(*releaseType))
		if err != nil {
			panic(err)
		}

		version.Pre[0] = prVersion
	}
}

func GenerateTag(currentTag *string, releaseType *ReleaseType, semanticVersionType *SemanticVersionType) string {
	version := makeSemanticVersion(currentTag)
	incrementVersion(&version, semanticVersionType)
	setSemanticVersionReleaseType(&version, releaseType)

	return version.String()
}
