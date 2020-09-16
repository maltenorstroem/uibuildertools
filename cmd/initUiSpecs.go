package cmd

import (
	"github.com/maltenorstroem/uibuildertools/internal/cmd/initUiSpecs"
	"github.com/spf13/cobra"
)

// initUiSpecsCmd represents the initUiSpecs command
var initUiSpecsCmd = &cobra.Command{
	Use:   "initUiSpecs",
	Short: "Creates u33e UI spec files",
	Long: `Creates the u33e UI specification files according to the hooks entered in the .uibuildertools config (hooks.types, hook.services)

To configure the hooks add a "hook" in the hooks section of your .uibuildertools config.

Example Config:

	[.uibuildertools]
	commands:
	  publish_npm: "./scripts/test.sh"
	hooks:
      service:
		- node_modules/@furo/ui-builder/_scripts/hook-init-reference-search.js
	  type:
		- node_modules/@furo/ui-builder/_scripts/hook-init-form.js
`,
	Run: initUiSpecs.Run,
}

// needed for the documentation generator
var InitUiSpecsCmd = initUiSpecsCmd

func init() {
	rootCmd.AddCommand(initUiSpecsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
