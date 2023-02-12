package cli

import (
	"testing"
)

func TestFeatParser(t *testing.T) {
	var featString string = "/* @feat feat1, feat2, feat3 */"
	var expectedFeatList []string = []string{"feat1", "feat2", "feat3"}
	var featList []string = FeatParser(featString)
	if len(featList) != len(expectedFeatList) {
		t.Errorf("Expected %d features, got %d", len(expectedFeatList), len(featList))
	}
	for i, feat := range featList {
		if feat != expectedFeatList[i] {
			t.Errorf("Expected %s, got %s", expectedFeatList[i], feat)
		}
	}
}

func TestRelToJSON(t *testing.T) {
	var rel Rel = Rel{
		File: "file1",
		Feat: []string{"feat1", "feat2", "feat3"},
		Rel:  []string{"rel1", "rel2", "rel3"},
	}
	var expectedRelJSON string = `{"file":"file1","features":["feat1","feat2","feat3"],"relations":["rel1","rel2","rel3"]}`
	var relJSON string = rel.ToJSON()
	if relJSON != expectedRelJSON {
		t.Errorf("Expected %s, got %s", expectedRelJSON, relJSON)
	}
}
