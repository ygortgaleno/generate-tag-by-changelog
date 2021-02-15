package regexUtils

import "regexp"

func MatchString(pattern *string, targetString *string) (matched bool) {
	matched, err := regexp.MatchString(*pattern, *targetString)
	if err != nil {
		panic(err)
	}

	return
}
