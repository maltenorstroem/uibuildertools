package ui_hooks

import "github.com/maltenorstroem/uibuildertools/pkg/u33eBuilder"

func Hook() (u33e u33eBuilder.U33eModel) {
	// Sample how to build a hook - Testing the u33e abstract model
	root := u33eBuilder.CreateU33e("furo-ui5-super-comp")
	// adds extension point and mixin
	root.Extends = "FBP(LitElement)"
	// adds theme
	root.Theme = "FormUI5BaseTheme"
	// adds description (chained)
	root.AddDescription("My super dooper mega component").AddSummary("Please provide a nice and meaningful summary")
	// declares imports and annotiation
	root.AddImportWithMember(u33eBuilder.MemberImport{
		Member: "i18n",
		Path:   "@furo/framework/src/i18n.js",
	}).AddImportWithMember(u33eBuilder.MemberImport{
		Member:     "fbp",
		Path:       "@furo/framework/src/fbp.js",
		Annotation: "eslint-disable-next-line no-unused-vars",
	}).AddImport("@furo/data/src/furo-data-object.js")
	// declares properties
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
	// Add CSS Stuff
	cssHost := make(map[string]string)
	cssHost["display"] = "block"
	cssHost["height"] = "100%"

	cssHostHidden := make(map[string]string)
	cssHostHidden["display"] = "none"

	root.AddCssStyleBlock(":host", cssHost).AddCssStyleBlock(":host([hidden])", cssHostHidden)

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

	// Create component in template root
	node := root.AddRootNode("furo-form-layouter")

	txtName := node.AddDomNode("furo-ui5-data-text-input")
	txtName.AddMethod("bind-data", "--entity(*.name)")

	txtFirstName := node.AddDomNode("furo-ui5-data-text-input")
	txtFirstName.AddMethod("bind-data", "--entity(*.first_name)")

	root.AddRootNode("furo-data-object")
	root.AddRootNode("furo-collection-agent")

	return root
}
