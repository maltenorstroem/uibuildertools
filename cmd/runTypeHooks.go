package cmd

import (
	"github.com/maltenorstroem/uibuildertools/internal/cmd/typeHooks"
	"github.com/spf13/cobra"
)

// initUiSpecsCmd represents the initUiSpecs command
var runTypeHooksCmd = &cobra.Command{
	Use:   "runTypeHooks",
	Short: "Runs the type hooks",
	Long: `Runs the type hooks according to the hooks entered in the .uibuildertools config (hooks.types)

To configure the hooks add a "hook" in the hooks section of your .uibuildertools config.

Example Config:

	[.uibuildertools]
	commands:
	  publish_npm: "./scripts/test.sh"
	hooks:
      service:
		- ./samples/uiHooks/testTypeHook.go
	  type:
		- ./samples/uiHooks/testServiceHook.go
`,
	Run: typeHooks.Run,
}

// needed for the documentation generator
var RunTypeHooksCmd = runTypeHooksCmd

func init() {
	rootCmd.AddCommand(RunTypeHooksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
