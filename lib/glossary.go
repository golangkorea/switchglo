package lib

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

// InfoBlock is information block consisting of term, translation and explanation
type InfoBlock struct {
	Term        string
	Translation string
	Explanation string
}

// Switch swaps term and translation
func (ib *InfoBlock) Switch() {
	term := ib.Term
	ib.Term = ib.Translation
	ib.Translation = term
}

func (ib *InfoBlock) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("## %s\n", ib.Term))
	buffer.WriteString(fmt.Sprintf("%s.", ib.Translation))
	buffer.WriteString(ib.Explanation)

	return buffer.String()
}

// ByTerm sort type for []InfoBlock
type ByTerm []InfoBlock

func (g ByTerm) Len() int {
	return len(g)
}

func (g ByTerm) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g ByTerm) Less(i, j int) bool {
	return strings.Compare(g[i].Term, g[j].Term) == -1
}

// Switch term with translation in all InfoBlock
func Switch(glossary []InfoBlock) {
	for index := 0; index < len(glossary); index++ {
		glossary[index].Switch()
	}
}

// Sort all InfoBlock
func Sort(glossary []InfoBlock) {
	sort.Sort(ByTerm(glossary))
}

// WriteToFile writes markdown to given file
func WriteToFile(out string, glossaries ...[]InfoBlock) error {
	var buffer bytes.Buffer
	for _, glossary := range glossaries {
		for _, ib := range glossary {
			buffer.WriteString(fmt.Sprintf("%s\n\n", ib.String()))
		}
	}
	err := ioutil.WriteFile(out, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

// NewGlossary creates new glossary
func NewGlossary(mdinput string) ([]InfoBlock, error) {

	glossary := []InfoBlock{}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(blackfriday.Run([]byte(mdinput))))

	if err != nil {
		return nil, errors.New("Failed to create goquery.Document object from input: " + mdinput)
	}
	var expidx int
	doc.Find("body").Children().Each(func(i int, s *goquery.Selection) {
		n := s.Get(0)
		switch n.Data {
		case "h2":
			html, err := s.Html()
			if err == nil {
				glossary = append(glossary, InfoBlock{Term: html, Translation: "", Explanation: ""})
			}
			expidx = 0
		case "pre":
			html, err := s.Find("code").Html()
			if err == nil {
				glossary[len(glossary)-1].Explanation += fmt.Sprintf("\n```\n%s\n```\n", html)
				expidx++
			}
		default:
			html, err := s.Html()
			if err == nil {
				switch expidx {
				case 0:
					tokens := strings.Split(html, ".")
					glossary[len(glossary)-1].Translation = tokens[0]
					glossary[len(glossary)-1].Explanation = fmt.Sprintf("%s\n", strings.Join(tokens[1:], " "))
				default:
					glossary[len(glossary)-1].Explanation += fmt.Sprintf("\n%s\n", html)
				}
				expidx++
			}
		}

	})

	return glossary, nil
}
