package initUiSpecs

import (
	"fmt"
	"github.com/maltenorstroem/uibuildertools/pkg/clientspec"
	"github.com/maltenorstroem/uibuildertools/pkg/u33eBuilder"
	"github.com/spf13/cobra"
)

type ClientTypeList map[string]*clientspec.Type

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("running initUiSpecs ...")

	// Testing the u33e abstract model
	root := u33eBuilder.CreateU33e("furo-ui5-super-comp")
	root.AddDescription("My super dooper mega component")
	root.AddSummary("Please provide a nice and meaningful summary")

	fmt.Println(root)
	// get spec directory from config
	// get all hook files
	// go through hooks and create u33e
}
