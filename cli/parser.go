package cli

import (
	"encoding/json"
	"strings"
)

type Rel struct {
	File string   `json:"file"`
	Feat []string `json:"features"`
	Rel  []string `json:"relations"`
}

func FeatParser(featString string) []string {
	var list string = ""
	split := strings.Split(featString, "@feat")
	if len(split) > 1 {
		list = split[1]
	} else {
		split = strings.Split(featString, "@rel")
		list = split[1]
	}
	var featList []string
	for _, feat := range strings.Split(list, ",") {
		if feat != " " && feat != "" {
			feat = strings.TrimSpace(feat)
			feat = strings.Split(feat, " ")[0]
			featList = append(featList, feat)
		}
	}

	return featList
}

func (r *Rel) ToJSON() string {
	relJson, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(relJson)
}
