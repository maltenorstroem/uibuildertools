package u33eBuilder

import (
	"encoding/json"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

// Creates a u33e abstract file model
func CreateU33e(component string) (u33e U33eModel) {
	var u33eFile U33eModel
	u33eFile.ComponentName = component

	return u33eFile
}

// Imports u33e ui spec file
func ImportU33e(fpath string) (u33e U33eModel) {
	dataBytes, readError := ioutil.ReadFile(fpath)
	if readError != nil {
		log.Fatal(readError)
	}
	parseError := yaml.Unmarshal(dataBytes, &u33e)

	if parseError != nil {
		log.Fatal(parseError)
	}
	return u33e
}

// Exports abstract u33e model to file
func ExportAbstractU33e(u33e U33eModel) (error error) {
	var file []byte
	outputFormat := viper.GetString("u33eFormat")
	if outputFormat == "yaml" {
		file, _ = yaml.Marshal(u33e)
	} else {
		file, _ = json.Marshal(u33e)
	}
	err := ioutil.WriteFile(viper.GetString("uiSpecsOutputDir")+"/test."+outputFormat, file, 0644)

	return err
}

// Adds a description to the model of the u33e abstract file
func (u33e *U33eModel) AddDescription(desc string) (u33 *U33eModel) {
	u33e.Description = desc
	return u33e
}

// Adds a summary to the model of the u33e abstract file
func (u33e *U33eModel) AddSummary(summary string) (u33 *U33eModel) {
	u33e.Summary = summary
	return u33e
}

// Adds an import with single or multiple exports from a module
func (u33e *U33eModel) AddImportWithMember(memberImport MemberImport) (u33 *U33eModel) {
	u33e.ImportMembers = append(u33e.ImportMembers, memberImport)
	return u33e
}

// Adds an import of a module for its side effects only
func (u33e *U33eModel) AddImport(importPath string) (u33 *U33eModel) {
	u33e.Imports = append(u33e.Imports, importPath)
	return u33e
}

// Adds a property
func (u33e *U33eModel) AddProperty(name string, prop Property) (u33 *U33eModel) {
	if u33e.Properties == nil {
		u33e.Properties = make(map[string]Property)
	}
	u33e.Properties[name] = prop
	return u33e
}

// Adds a method
func (u33e *U33eModel) AddMethod(name string, method Method) (u33 *U33eModel) {
	if u33e.Methods == nil {
		u33e.Methods = make(map[string]Method)
	}
	u33e.Methods[name] = method
	return u33e
}

// Adds an exposed wire function
func (u33e *U33eModel) AddExposedWire(name string, expWire ExposedWire) (u33 *U33eModel) {
	if u33e.ExposedWires == nil {
		u33e.ExposedWires = make(map[string]ExposedWire)
	}
	u33e.ExposedWires[name] = expWire
	return u33e
}

// Adds a keyboard shortcut
func (u33e *U33eModel) AddKeyboardShortcut(key KeyboardShortcut) (u33 *U33eModel) {
	u33e.KeyboardShortcuts = append(u33e.KeyboardShortcuts, key)
	return u33e
}

// Adds a css style block
func (u33e *U33eModel) AddCssStyleBlock(selector string, cssBlock map[string]string) (u33 *U33eModel) {
	if u33e.Styles == nil {
		u33e.Styles = make(map[string]map[string]string)
	}
	u33e.Styles[selector] = cssBlock
	return u33e
}

// Adds a dom node to the template root
// Do not forget to add a import for the used componentName
func (u33e *U33eModel) AddRootNode(componentname string) (node *DomNode) {

	comp := DomNode{
		Component:   componentname,
		Description: "Please provide a description",
		Flags:       nil,
		Attributes:  nil,
		Methods:     nil,
		Events:      nil,
		InnerText:   "",
		Children:    nil,
	}
	u33e.Template = append(u33e.Template, &comp)
	return &comp
}

// Adds a dom node to the template root
// Do not forget to add a import for the used componentName
func (parent *DomNode) AddDomNode(componentname string) (node *DomNode) {

	comp := DomNode{
		Component:   componentname,
		Description: "Please provide a description",
		Flags:       nil,
		Attributes:  nil,
		Methods:     nil,
		Events:      nil,
		InnerText:   "",
		Children:    nil,
	}
	parent.Children = append(parent.Children, &comp)
	return &comp
}

// Adds attributes to a DomNode
func (domNode *DomNode) AddAttribute(key string, value string) (node *DomNode) {
	if domNode.Attributes == nil {
		domNode.Attributes = make(map[string]string)
	}
	domNode.Attributes[key] = value
	return domNode
}

// Adds flow based method to a DomNode
func (domNode *DomNode) AddMethod(name string, wire string) (node *DomNode) {
	if domNode.Methods == nil {
		domNode.Methods = make(map[string]string)
	}
	domNode.Methods["ƒ-"+name] = wire
	return domNode
}

// Adds flow based event listener to a DomNode
func (domNode *DomNode) AddEventlistener(event string, wire string) (node *DomNode) {
	if domNode.Events == nil {
		domNode.Events = make(map[string]string)
	}
	domNode.Events["@-"+event] = wire
	return domNode
}
