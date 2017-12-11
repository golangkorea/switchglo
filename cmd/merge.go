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

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge glossary with its translated terms",
	Long:  `Merges existing glossary with it translated terms`,
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
		lib.Sort(glossary)

		input, err = ioutil.ReadFile(cmd.Flag("file").Value.String())
		if err != nil {
			log.Fatalf("Failed to read glossary file: %s", err.Error())
			os.Exit(1)
		}
		switched, err := lib.NewGlossary(string(input))
		if err != nil {
			log.Fatalf("Failed to create InfoBlock slice from markdown: %s", err.Error())
			os.Exit(1)
		}
		lib.Switch(switched)
		lib.Sort(switched)

		outputFile := cmd.Flag("out").Value.String()
		err = lib.WriteToFile(outputFile, glossary, switched)
		if err != nil {
			log.Fatalf("Failed to output new glossary file, %s: %s", outputFile, err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(mergeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mergeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mergeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
