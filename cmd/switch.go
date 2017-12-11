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
	"io/ioutil"
	"log"
	"os"

	"github.com/golangkorea/switchglo/lib"

	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch terms with their translations",
	Long: `switchglo allows terms to be switched with their translated terms, vice versa.

In order to execute these automations reliably, translators should maintain following structures.

	1) Without explanation
	## Term
	translated term

	2) With explanation
	## Term
	translated term. The explanation follows.`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := ioutil.ReadFile(cmd.Flag("file").Value.String())
		if err != nil {
			log.Fatalf("Failed to read glossary file: %s", err.Error())
			os.Exit(1)
		}
		glossary, err := lib.NewGlossary(string(input))
		if err != nil {
			log.Fatalf("Failed to create InfoBlock slice from markdown: %s", err.Error())
			os.Exit(1)
		}
		lib.Switch(glossary)
		lib.Sort(glossary)

		outputFile := cmd.Flag("out").Value.String()
		err = lib.WriteToFile(outputFile, glossary)
		if err != nil {
			log.Fatalf("Failed to output new glossary file, %s: %s", outputFile, err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(switchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// switchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// switchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
