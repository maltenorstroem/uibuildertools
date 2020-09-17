package initUiSpecs

import (
	"encoding/json"
	"fmt"
	"github.com/maltenorstroem/uibuildertools/pkg/clientspec"
	"github.com/maltenorstroem/uibuildertools/pkg/u33eBuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ClientTypeList map[string]*clientspec.Type

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("running initUiSpecs ...")

	// Testing the u33e abstract model
	root := u33eBuilder.CreateU33e("furo-ui5-super-comp")
	root.Extends = "FBP(LitElement)"
	root.AddDescription("My super dooper mega component").AddSummary("Please provide a nice and meaningful summary").AddImportWithMember(u33eBuilder.MemberImport{
		Member: "i18n",
		Path:   "@furo/framework/src/i18n.js",
	}).AddImportWithMember(u33eBuilder.MemberImport{
		Member:     "fbp",
		Path:       "@furo/framework/src/fbp.js",
		Annotation: "eslint-disable-next-line no-unused-vars",
	}).AddImport("@furo/data/src/furo-data-object.js")

	// Properties
	root.AddProperty("type", u33eBuilder.Property{
		Description:   "Put your description here",
		PropertyType:  "string",
		Reflect:       false,
		AttributeName: "type",
	}).AddProperty("design", u33eBuilder.Property{
		Description:   "Defines the design. Available options are \"Default\", \"Emphasized\", \"Positive\", \"Negative\", and \"Transparent\".",
		PropertyType:  "string",
		Reflect:       true,
		AttributeName: "design",
	})

	// Methods
	root.AddMethod("bindData", u33eBuilder.Method{
		Args:        "data",
		Description: "Bind your furo-data-object event @-object-ready\n @public\n @param data",
		Code:        "CiAgICB0aGlzLl9GQlBUcmlnZ2VyV2lyZSgnLS1kYXRhJywgZGF0YSk7CiAgICB0aGlzLmZpZWxkID0gZGF0YTs=",
	})

	// Keyboard Shortcuts
	root.AddKeyboardShortcut(u33eBuilder.KeyboardShortcut{
		Description: "Shortcut SAVE",
		Key:         "S",
		Ctrl:        true,
		Global:      false,
		Alt:         false,
		Wire:        "--savePressed",
	})

	// Exposed Wire Function
	root.AddExposedWire("focus", u33eBuilder.ExposedWire{
		Name:        "focus",
		Wire:        "--focus",
		Description: "Focus main element",
	})

	// Add theme
	root.Theme = "FormUI5BaseTheme"

	// Add CSS Stuff

	//importedU33e := u33eBuilder.ImportU33e("./samples/u33e/test-form-ui5.u33e")

	var file []byte
	//var importedFile []byte
	outputFormat := viper.GetString("u33eFormat")
	if outputFormat == "yaml" {
		file, _ = yaml.Marshal(root)
		//	importedFile, _ = yaml.Marshal(importedU33e)
	} else {
		file, _ = json.Marshal(root)
		//	importedFile, _ = json.Marshal(importedU33e)
	}
	_ = ioutil.WriteFile(viper.GetString("uiSpecsOutputDir")+"/test."+outputFormat, file, 0644)
	//_ = ioutil.WriteFile(viper.GetString("uiSpecsOutputDir")+"/test-import."+outputFormat, importedFile, 0644)

	fmt.Println("\nstruct out")
	fmt.Println(root)

}
