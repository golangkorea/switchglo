// Copyright © 2017 Jhonghee Park <jhonghee@gmail.com>
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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var glossaryFile = "glossary.md"
var outFile = "glossary_out.md"

// RootCmd is root command
var RootCmd = &cobra.Command{
	Use:   "switchglo",
	Short: "Switching and sorting terms in glossary",
	Long: `switchglo helps translators to manage glossary markdown. It automates followings:
	
	1) Translators do not need to maintain the alphabetical order of terms. switchglo does it.
	2) switchglo allows terms to be switched with their translated terms, vice versa.
	3) switchglo will produce a merged glossary from existing glossary and its switched form.
	
In order to execute these automations reliably, translators should maintain following structures.
	
	1) Without explanation
	## Term
	translated term

	2) With explanation
	## Term
	translated term. The explanation follows.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.switchglo.yaml)")
	RootCmd.PersistentFlags().StringVar(&glossaryFile, "file", "glossary.md", "glossary file (default is glossary.md)")
	RootCmd.PersistentFlags().StringVar(&outFile, "out", "glossary_out.md", "output file (default is glossary_out.md)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".switchglo") // name of config file (without extension)
	viper.AddConfigPath("$HOME")      // adding home directory as first search path
	viper.AutomaticEnv()              // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
