package u33eBuilder

// Internal model of the abstract u33e file
type U33eModel struct {
	Description       string                       `json:"description" yaml:"description"`
	Summary           string                       `json:"summary" yaml:"summary"`
	ImportMembers     []MemberImport               `json:"import_members" yaml:"import_members"`
	Imports           []string                     `json:"imports" yaml:"imports"`
	ComponentName     string                       `json:"component_name" yaml:"component_name"`
	Extends           string                       `json:"extends" yaml:"extends"`
	Path              string                       `json:"path" yaml:"path"`
	Styles            map[string]map[string]string `json:"styles" yaml:"styles"`
	Template          []*DomNode                   `json:"template" yaml:"template"`
	Properties        map[string]Property          `json:"property" yaml:"property"`
	ExposedWires      map[string]ExposedWire       `json:"exposed_wire" yaml:"exposed_wire"`
	Methods           map[string]Method            `json:"methods" yaml:"methods"`
	KeyboardShortcuts []KeyboardShortcut           `json:"keyboard_shortcuts" yaml:"keyboard_shortcuts"`
	Theme             string                       `json:"theme" yaml:"theme"`
}

// Declares an import with single or multiple exports from a module
type MemberImport struct {
	Member     string `json:"member" yaml:"member"`
	Path       string `json:"path" yaml:"path"`
	Annotation string `json:"annotation" yaml:"annotation"`
}

// Declares a method
// The attribute `Code` is understood as a base64 encoded string
type Method struct {
	Args        string `json:"args" yaml:"args"`
	Description string `json:"description" yaml:"description"`
	Code        string `json:"code" yaml:"code"`
}

// Declares an abstract web component property
type Property struct {
	Description   string `json:"description" yaml:"description"`
	PropertyType  string `json:"property_type" yaml:"property_type"`
	Reflect       bool   `json:"reflect" yaml:"reflect"`
	AttributeName string `json:"attribute_name" yaml:"attribute_name"`
}

// Css style block
type CssBlock struct {
	CssProps []CssProp `json:"css_props" yaml:"css_props"`
}

// Css single property
type CssProp struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
}

// Declares a keyboard shortcut
type KeyboardShortcut struct {
	Description string `json:"description" yaml:"description"`
	Key         string `json:"key" yaml:"key"`
	Ctrl        bool   `json:"ctrl" yaml:"ctrl"`
	Global      bool   `json:"global" yaml:"global"`
	Alt         bool   `json:"alt" yaml:"alt"`
	Wire        string `json:"wire" yaml:"wire"`
}

// Declares an exposed wire function
type ExposedWire struct {
	Name        string `json:"name" yaml:"name"`
	Wire        string `json:"wire" yaml:"wire"`
	Description string `json:"description" yaml:"description"`
}

// Simplified representation of a dome node
type DomNode struct {
	Component   string            `json:"component" yaml:"component"`
	Description string            `json:"description" yaml:"description"`
	Flags       []string          `json:"flags" yaml:"flags"`
	Attributes  map[string]string `json:"attributes" yaml:"attributes"`
	Methods     map[string]string `json:"methods" yaml:"methods"`
	Events      map[string]string `json:"events" yaml:"events"`
	InnerText   string            `json:"inner_text" yaml:"inner_text"`
	Children    []*DomNode        `json:"children" yaml:"children"`
}
