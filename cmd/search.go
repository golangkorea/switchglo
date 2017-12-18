// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/golangkorea/switchglo/lib"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search text in glossary",
	Long:  `Search text in glossary and returns the list of terms and translations`,
	Run: func(cmd *cobra.Command, args []string) {

		input, err := ioutil.ReadFile(cmd.Flag("file").Value.String())
		if err != nil {
			log.Fatalf("Failed to read glossary file: %s", err.Error())
		}
		glossary, err := lib.NewGlossary(string(input))
		index := lib.NewSearchIndex(glossary)

		reader := bufio.NewReader(os.Stdin)
		var searchTerm string
		for len(strings.TrimSpace(searchTerm)) == 0 {
			fmt.Println("Please refer to http://www.blevesearch.com/docs/Query-String-Query/ for query language.")
			fmt.Print("Enter Search Term (*): ")
			searchTerm, _ = reader.ReadString('\n')
		}

		query := bleve.NewQueryStringQuery(searchTerm)
		searchRequest := bleve.NewSearchRequest(query)
		searchRequest.Highlight = bleve.NewHighlight()
		searchResults, err := index.Search(searchRequest)
		if err != nil {
			log.Fatalf("Failed to search with %s: %s", searchTerm, err.Error())
		}

		srTxt := searchResults.String()
		srTxt = strings.Replace(srTxt, "<mark>", "*", -1)
		srTxt = strings.Replace(srTxt, "</mark>", "*", -1)
		fmt.Println(srTxt)
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
