/*
Package cmd defines commands
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"go-dynamic-fetch/fetcher"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-dynamic-fetch",
	Short: "Provides utility to load dynamic web page content",
	Long:  `Allows you to request data from dynamic web pages and interact with it`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// put everything in a config
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-dynamic-fetch.yaml)")

	// specify all configurable values as options instead
	rootCmd.PersistentFlags().Bool("headless", false, "Use headless shell")
	rootCmd.PersistentFlags().StringP("url", "u", "", "URL that you are fetching HTML content for")
	rootCmd.PersistentFlags().StringP("agent", "a", fetcher.DefaultUserAgent, "User agent to request as - if not specified the default is used")
	rootCmd.PersistentFlags().StringP("selector", "s", "", "Selector for element to wait for - if not specified we do not wait and just dump static elements")
	rootCmd.PersistentFlags().IntP("timeout", "t", -1, "Timeout for context - if none is specified a default background context will be used")
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

		// Search config in home directory with name ".go-dynamic-fetch" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-dynamic-fetch")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
