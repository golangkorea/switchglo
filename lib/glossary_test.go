package lib

import (
	"fmt"
	"testing"
)

var testcases = []struct {
	md    string
	valid []InfoBlock
}{
	{
		`## term4
번역4. 추가설명. https://www.google.com

## term1
번역1

## term3
번역3. explanation in multiple line

additional line 1

` +
			"```\n" +
			"additional line 2\n" +
			"```" +
			`

## term2
번역2. explanation`,
		[]InfoBlock{
			InfoBlock{Term: "term4", Translation: "번역4", Explanation: " 추가설명. <a href=\"https://www.google.com\">https://www.google.com</a>"},
			InfoBlock{Term: "term1", Translation: "번역1", Explanation: ""},
			InfoBlock{Term: "term3", Translation: "번역3", Explanation: " explanation in multiple lineadditional line 1\n\n```\nadditional line 2\n\n```\n\n"},
			InfoBlock{Term: "term2", Translation: "번역2", Explanation: " explanation"},
		},
	},
}

func TestNewGlossary(t *testing.T) {

	glossary, _ := NewGlossary(testcases[0].md)

	for index := 0; index < len(glossary); index++ {
		if glossary[index].Term != testcases[0].valid[index].Term ||
			glossary[index].Translation != testcases[0].valid[index].Translation ||
			glossary[index].Explanation != testcases[0].valid[index].Explanation {
			t.Fatalf("Expected %#v but got %#v", testcases[0].valid[index], glossary[index])
		}
	}
}

func TestSwitch(t *testing.T) {

	glossary, _ := NewGlossary(testcases[0].md)

	Switch(glossary)

	for index := 0; index < len(glossary); index++ {
		if glossary[index].Term != testcases[0].valid[index].Translation ||
			glossary[index].Translation != testcases[0].valid[index].Term ||
			glossary[index].Explanation != testcases[0].valid[index].Explanation {
			t.Fatalf("Expected %#v but got %#v", testcases[0].valid[index], glossary[index])
		}
	}
}

func TestSort(t *testing.T) {

	glossary, _ := NewGlossary(testcases[0].md)

	Sort(ByTerm(glossary))

	if glossary[0].Term != testcases[0].valid[1].Term ||
		glossary[0].Translation != testcases[0].valid[1].Translation ||
		glossary[0].Explanation != testcases[0].valid[1].Explanation {
		t.Fatalf("Expected %#v but got %#v", testcases[0].valid[1], glossary[0])
	}
	if glossary[2].Term != testcases[0].valid[2].Term ||
		glossary[2].Translation != testcases[0].valid[2].Translation ||
		glossary[2].Explanation != testcases[0].valid[2].Explanation {
		t.Fatalf("Expected %#v but got %#v", testcases[0].valid[2], glossary[2])
	}

	Switch(glossary)
	Sort(ByTerm(glossary))

	if glossary[0].Term != testcases[0].valid[1].Translation ||
		glossary[0].Translation != testcases[0].valid[1].Term ||
		glossary[0].Explanation != testcases[0].valid[1].Explanation {
		t.Fatalf("Expected %#v but got %#v", testcases[0].valid[1], glossary[0])
	}
	if glossary[2].Term != testcases[0].valid[2].Translation ||
		glossary[2].Translation != testcases[0].valid[2].Term ||
		glossary[2].Explanation != testcases[0].valid[2].Explanation {
		t.Fatalf("Expected %#v but got %#v", testcases[0].valid[2], glossary[2])
	}
}

func TestSearch(t *testing.T) {
	glossary, _ := NewGlossary(testcases[0].md)

	index, err := NewSearchIndex(glossary)
	if err != nil {
		t.Errorf("Failed to create new index: %s", err.Error())
		t.Fail()
	}

	searchResult, err := Search("term1", index)

	for _, hit := range searchResult.Hits {
		document := hit.Document
		fmt.Println(document)
	}

}
