# Abstract U33E UI Builder

## How to use

``` Go
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
```

## Output
``` Yaml
description: My super dooper mega component
summary: Please provide a nice and meaningful summary
import_members:
- member: i18n
  path: '@furo/framework/src/i18n.js'
  annotation: ""
- member: fbp
  path: '@furo/framework/src/fbp.js'
  annotation: eslint-disable-next-line no-unused-vars
imports:
- '@furo/data/src/furo-data-object.js'
component_name: furo-ui5-super-comp
extends: FBP(LitElement)
path: ""
styles:
  :host:
    display: block
    height: 100%
  :host([hidden]):
    display: none
template:
- component: furo-form-layouter
  description: Please provide a description
  flags: []
  attributes: {}
  methods: {}
  events: {}
  inner_text: ""
  children:
  - component: furo-ui5-data-text-input
    description: Please provide a description
    flags: []
    attributes: {}
    methods:
      ƒ-bind-data: --entity(*.name)
    events: {}
    inner_text: ""
    children: []
  - component: furo-ui5-data-text-input
    description: Please provide a description
    flags: []
    attributes: {}
    methods:
      ƒ-bind-data: --entity(*.first_name)
    events: {}
    inner_text: ""
    children: []
- component: furo-data-object
  description: Please provide a description
  flags: []
  attributes: {}
  methods: {}
  events: {}
  inner_text: ""
  children: []
- component: furo-collection-agent
  description: Please provide a description
  flags: []
  attributes: {}
  methods: {}
  events: {}
  inner_text: ""
  children: []
property:
  design:
    description: Defines the design. Available options are "Default", "Emphasized", "Positive", "Negative", and "Transparent".
    property_type: string
    reflect: true
    attribute_name: design
  type:
    description: Put your description here
    property_type: string
    reflect: false
    attribute_name: type
exposed_wire:
  focus:
    name: focus
    wire: --focus
    description: Focus main element
methods:
  bindData:
    args: data
    description: |-
      Bind your furo-data-object event @-object-ready
       @public
       @param data
    code: CiAgICB0aGlzLl9GQlBUcmlnZ2VyV2lyZSgnLS1kYXRhJywgZGF0YSk7CiAgICB0aGlzLmZpZWxkID0gZGF0YTs=
keyboard_shortcuts:
- description: Shortcut SAVE
  key: S
  ctrl: true
  global: false
  alt: false
  wire: --savePressed
theme: FormUI5BaseTheme

```