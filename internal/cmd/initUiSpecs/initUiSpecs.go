package initUiSpecs

import (
	"fmt"
	"github.com/maltenorstroem/uibuildertools/pkg/clientspec"
	"github.com/maltenorstroem/uibuildertools/pkg/u33eBuilder"
	"github.com/maltenorstroem/uibuildertools/samples/ui_hooks"
	"github.com/spf13/cobra"
)

type ClientTypeList map[string]*clientspec.Type

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("running initUiSpecs ...")

	// just for testing purposes
	// we have to inject the hooks dynamically
	model := ui_hooks.Hook()
	u33eBuilder.ExportAbstractU33e(model)

	fmt.Println("\nstruct out")
	fmt.Println(model)

}
