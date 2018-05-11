// Copyright © 2018 Ryan Armstrong <cowboys6750@gmail.com>
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

	homedir "github.com/mitchellh/go-homedir"
	bamboo "github.com/rcarmstrong/go-bamboo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const version = "BambooCTL v0.1.0"

var (
	cli         *bamboo.Client
	cfgFile     string
	versionFlag bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "bambooctl",
	Short: "bambooctl is a commandline tool to help manage and interact with the Atlassian Bamboo CI server.",
	Long: `bambooctl [--version] [--help] <command> [args]

bambooctl is a commandline tool to help manage and interact with the Atlassian Bamboo CI server. 
Many of the commands require admin privleges, but some can be accessed with lesser user permissions.
	
	Admin Commands:
		project			Project related operations
		
		
	Non-Admin Commands:
		ToDo			ToDo`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bambooctl.yaml)")
	RootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "")

	// Blank Run func to allow the output of RootCmd.Use in error messages and help output
	RootCmd.Run = func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Println(version)
		} else {
			RootCmd.Usage()
		}
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".bambooctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".bambooctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading in config:", err)
		os.Exit(1)
	}

	cli = bamboo.NewSimpleClient(nil, viper.GetString("username"), viper.GetString("password"))

	if viper.GetString("url") != "" {
		cli.SetURL(viper.GetString("url"))
	}
}
