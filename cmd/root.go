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
	"os"

	"github.com/vishnraj/go-dynamic-fetch/fetcher"

	"github.com/spf13/cobra"

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
		fetcher.Log().Panicf("%v", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// put everything in a config
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-dynamic-fetch.yaml)")

	// specify all configurable values as options instead
	// additional flags will be specified in sub-commands
	rootCmd.PersistentFlags().Bool("headless", false, "Use headless shell")
	rootCmd.PersistentFlags().StringSliceP("agents", "a", fetcher.DefaultUserAgents, "User agent(s) to request as - if not specified the default is used")
	rootCmd.PersistentFlags().IntP("timeout", "t", -1, "Timeout for context - if none is specified a default background context will be used")
	rootCmd.PersistentFlags().Bool("wait_error_dump", false, "If an error is encountered during the wait phase, where the expected element is not loaded, dump the page contents")
	rootCmd.PersistentFlags().Bool("wait_error_location", false, "If an error is encountered during the wait phase, where the expected element is not loaded, log the current URL we are at")

	rootCmd.PersistentFlags().Bool("redis_dump", false, "Set this option for all dumps to go to the redis database that we connet to this app")
	rootCmd.PersistentFlags().String("redis_url", "", "If we want to send dumps to a redis database we must set a valid URL")
	rootCmd.PersistentFlags().String("redis_password", "", "If we need a password to login to the redis database, specify it")
	rootCmd.PersistentFlags().Int("redis_key_expiration", 0, "The duration, in secondds that keys will remain in redis for - default value of zero makes this indefinite")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find working directory.
		home, err := os.Getwd()
		if err != nil {
			fetcher.Log().Panicf("%v", err)
		}

		// Search config in home directory with name ".go-dynamic-fetch" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-dynamic-fetch")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fetcher.Log().Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}
