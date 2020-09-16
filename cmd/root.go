/*
Copyright © 2020

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
	"github.com/spf13/cobra"
	"github.com/theNorstroem/uibuildertools/internal/cmd/runner"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string

// needed for the documentation generator
var RootCmd = rootCmd

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "uibuildertools",
	Short: "The furo ui builder toolkit",
	Long: `Furo uibuildertools contains helpful generators and templates to produce 
pattern based ui elements based on a furo spec.
Read more about the single commands in the see also section below.

Calling uibuildertools without any arguments and flags will run the flow runner 
with the default flow. 
Modify your default flow in the .uibuildertools config file to your needs. You can set any of the sub commands as default.

> Note: Environment variables are prefixed with **FUT**. 
>
> To set the specformat with the environment variable use **FUT_SPECFORMAT=value**
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		runner.Run(cmd, args)
	},
	Version: "0.0.1",
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is CWD/.uibuildertools.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Search config in home directory with name ".uibuildertools" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
		viper.SetConfigName(".uibuildertools")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix("FUT")
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}