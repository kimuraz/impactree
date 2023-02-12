package cli

import "strings"

type Feat struct {
	Name string
}

func FeatParser(featString string) []Feat {
	strings.ReplaceAll(featString, " ", "")
	strings.ReplaceAll(featString, "@feat:", "")
	var featList []Feat
	for _, feat := range strings.Split(featString, ",") {
		featList = append(featList, Feat{Name: feat})
	}

	return featList
}
