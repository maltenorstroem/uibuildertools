package initUiSpecs

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/theNorstroem/uibuildertools/pkg/clientspec"
)

type ClientTypeList map[string]*clientspec.Type

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("running initUiSpecs ...")

	// get spec directory from config
	// get all hook files
	// go through hooks and create u33e
}
