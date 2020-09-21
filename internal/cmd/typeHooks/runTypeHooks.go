package typeHooks

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"plugin"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("running type hooks ...")

	// load plugins from .uibuilderhooks hooks section
	plugins := []string{"./samples/uiHooks/testTypeHook.so"}
	LoadPlugins(plugins)
}

func LoadPlugins(plugins []string) {

	for _, p := range plugins {

		p, err := plugin.Open(p)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// 2. look up a symbol (an exported function or variable)
		// in this case, variable TestTypeHook
		symTestTypeHook, err := p.Lookup("Hook")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		symTestTypeHook.(func())()

		u33e, err := p.Lookup("U33e")
		if err != nil {
			panic(err)
		}
		fmt.Println(u33e)
	}

}
