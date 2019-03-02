// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	log  "github.com/sirupsen/logrus"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// default in initConfig, unless passed as flag
var cfgFile string
var verbose bool
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "blync-studio-light",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {	
		if verbose {
			log.SetLevel(log.InfoLevel)
			log.Info("Verbose logging enabled")
		} 
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
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
	log.SetLevel(log.WarnLevel)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.blync-studio-light.yaml)")
	rootCmd.PersistentFlags().IntP("device", "d", 0, "Device index for light to interface with")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose","v", false, "Include info level logs")
	//nolint:errcheck
	viper.BindPFlag("device", rootCmd.PersistentFlags().Lookup("device"))
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")


	cobra.OnInitialize(initConfig)


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

		// Search config in home directory with name ".blync-studio-light" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".blync-studio-light")
		cfgFile = home + "/.blync-studio-light.json"
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Infoln("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Warnln("No configuration found, functionality will be limited until you run `config init` ")
	}
}
