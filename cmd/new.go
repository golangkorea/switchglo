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
	"log"
	"os"
	"strings"

	"github.com/golangkorea/switchglo/lib"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add new glossary item",
	Long: `new command add a new item to existing glossary. It requires 3 pieces of information:
	
1) Term
2) Translation
3) Explanation (Optional)`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		var term string
		for len(strings.TrimSpace(term)) == 0 {
			fmt.Print("Enter Term (*): ")
			term, _ = reader.ReadString('\n')
		}

		var translation string
		for len(strings.TrimSpace(translation)) == 0 {
			fmt.Print("Enter Translation (*): ")
			translation, _ = reader.ReadString('\n')
		}

		fmt.Print("Enter Explanation (^Z to end): ")
		explanation, _ := reader.ReadString('\032') // Enter ctrl-z to end the input

		newInfoBlock := lib.InfoBlock{
			Term:        strings.TrimSpace(term),
			Translation: strings.TrimSpace(translation),
			Explanation: explanation,
		}

		file, err := os.OpenFile(cmd.Flag("file").Value.String(), os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
		defer file.Close()

		_, err = file.WriteString(newInfoBlock.String())
		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}

	},
}

func init() {
	RootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
