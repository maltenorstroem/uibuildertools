package initUiSpecs

import (
	"fmt"
	"github.com/maltenorstroem/uibuildertools/pkg/clientspec"
	"github.com/spf13/cobra"
)

type ClientTypeList map[string]*clientspec.Type

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("running initUiSpecs ...")

}
