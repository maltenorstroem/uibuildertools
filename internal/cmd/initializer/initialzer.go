package initializer

import (
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

//	module: "github.com/sample/repo" #the name of the git repo where your spec project resides
//	version: "v0.0.1"
//	dependencies:
//	- "git@git.devres.internal.adcubum.com:7999/syrsl/syrius-claimsmgmt-servicesmgmt-bff-spec.git master"
//	- "github.com/theNorstroem/furoBaseTypes v1.0.0" # The importer looks for all **/*.type.spec files recursive The importer looks for all **/*.service.spec files recursive
//	dependenciesDir: installedSpecs
//	u33eFormat: "yaml" #set to yaml or json
//	installedSpecDir: "./specs"
//	uiSpecsOutputDir: "./ui_specs/"
//
//	commands:
//	dummyCmd: "scripts/dummy.sh"
//
//	flows:
//default:
//	- initUiSpecs
//
//	hooks:
//	types:
//	- "./node_modules/@furo/ui-builder/_scripts/hook-initializer-form-ui5.js"
//	services:
//	- "./node_modules/@furo/ui-builder/_scripts/hook-initializer-reference-search-ui5.js"
//	build:
//	components:
//	targetDir: "./dist/components/generated_components"
type UiBuilderConfig struct {
	Module           string              `json:"module" yaml:"module"`
	Version          string              `json:"version" yaml:"version"`
	Dependencies     []string            `json:"dependencies" yaml:"dependencies"`
	DependencyDir    string              `json:"dependency_dir" yaml:"dependency_dir"`
	U33eFormat       string              `json:"u33e_format" yaml:"u33e_format"`
	InstalledSpecDir string              `json:"installed_spec_dir" yaml:"installed_spec_dir"`
	UiSpecsOutputDir string              `json:"ui_specs_output_dir" yaml:"ui_specs_output_dir"`
	Commands         map[string]string   `json:"commands" yaml:"commands"`
	Flows            map[string][]string `json:"flow" yaml:"flows"`
	Build            map[string]string   `json:"build" yaml:"build"`
}

func Run(cmd *cobra.Command, args []string) {
	// check if .uibuildertools file exist
	_, err := checkConfig()

	if err != nil {
		createDefaultConfig()
		log.Println("Default .uibuildertools config created.")
	} else {
		log.Println("Existing .uibuildertools config used.")
	}

}

// Checks if uibuildertools config file is present
func checkConfig() (conf *UiBuilderConfig, err error) {
	dataBytes, fileErr := ioutil.ReadFile(".uibuildertools")
	if fileErr != nil {
		err = fileErr
	}
	parseErr := yaml.Unmarshal(dataBytes, &conf)
	if parseErr != nil {
		err = fileErr
	}
	return conf, err
}

// Creates a default uibuildertools config file
func createDefaultConfig() (conf *UiBuilderConfig) {
	defaultConfig := UiBuilderConfig{
		Module:           "github.com/sample/repo",
		Version:          "v0.0.1",
		Dependencies:     nil,
		DependencyDir:    "",
		U33eFormat:       "yaml",
		InstalledSpecDir: "./specs",
		UiSpecsOutputDir: "./ui_specs",
		Commands:         nil,
		Flows:            nil,
	}
	defaultConfig.Flows = make(map[string][]string)
	defaultFlow := []string{
		"initialize",
		"initUiSpecs",
	}
	defaultConfig.Flows["default"] = defaultFlow
	defaultConfig.Build = make(map[string]string)
	defaultConfig.Build["componentsTargetDir"] = "./dist/components/generated_components"

	var confOutput []byte
	confOutput, _ = yaml.Marshal(defaultConfig)

	_ = ioutil.WriteFile(".uibuildertools", confOutput, 0644)
	return conf
}
