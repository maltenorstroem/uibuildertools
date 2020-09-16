package u33eBuilder

// Representation of a U33E UI Builder File
type U33e struct {
	model U33eModel `json:"model" yaml:"model"`
}

type U33eModel struct {
	Description   string   `json:"description" yaml:"description"`
	Summary       string   `json:"summary" yaml:"summary"`
	ImportMembers []string `json:"import_members" yaml:"import_members"`
	Imports       []string `json:"imports" yaml:"imports"`
	ComponentName string   `json:"component_name" yaml:"component_name"`
	Path          string   `json:"path" yaml:"path"`
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
	Children    []DomNode         `json:"children" yaml:"children"`
}

// Creates a u33e abstract file model
func CreateU33e(component string) (u33e U33e) {
	var u33eFile U33e
	u33eFile.model.ComponentName = component

	return u33eFile

}

// Adds a description to the model of the u33e abstract file
func (u33e *U33e) AddDescription(desc string) {
	u33e.model.Description = desc
}

// Adds a summary to the model of the u33e abstract file
func (u33e *U33e) AddSummary(summary string) {
	u33e.model.Summary = summary
}
