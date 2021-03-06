package typeAst

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/theNorstroem/spectools/pkg/specSpec"
	"github.com/theNorstroem/spectools/pkg/util"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

type Typelist struct {
	TypesByName          map[string]*TypeAst
	InstalledTypesByName map[string]*TypeAst
	SpecDir              string
}

type TypeAst struct {
	Path     string // relative path of spec file to SpecDir
	FileName string
	TypeSpec specSpec.Type
}

var Format = "json"

// set the storage format
func (l *Typelist) setStorageFormat(format string) {
	Format = format
}

// loads a spec directory and installed specs to the typelist
func (l *Typelist) LoadTypeSpecsFromDir(specDir string) {
	l.TypesByName = loadTypeSpecsFromDir(specDir)
}

// loads a spec directory and installed specs to the typelist
func (l *Typelist) LoadInstalledTypeSpecsFromDir(specDir ...string) {
	// create map if it does not exist
	if l.InstalledTypesByName == nil {
		l.InstalledTypesByName = map[string]*TypeAst{}
	}
	for _, dir := range specDir {
		tlist := loadTypeSpecsFromDir(dir)
		for tname, v := range tlist {
			l.InstalledTypesByName[tname] = v
		}
	}
}

func loadTypeSpecsFromDir(specDir string) (typesMap map[string]*TypeAst) {
	typesMap = map[string]*TypeAst{}
	err := filepath.Walk(specDir,
		func(fpath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasSuffix(fpath, "type.spec") {
				filename := path.Base(fpath)
				sdlen := len(strings.Split(specDir, "/"))
				if strings.HasPrefix(specDir, "./") {
					sdlen--
				}

				relativePath := path.Dir(strings.Join(strings.Split(fpath, "/")[sdlen:], "/"))
				AstType := &TypeAst{
					Path:     relativePath, // store Path without specDir
					FileName: filename,
					TypeSpec: readAndUnmarshalSpec(fpath),
				}

				typesMap[strings.Join([]string{AstType.TypeSpec.XProto.Package, AstType.TypeSpec.Type}, ".")] = AstType
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	return typesMap
}

func readAndUnmarshalSpec(fpath string) (s specSpec.Type) {
	dataBytes, readError := ioutil.ReadFile(fpath)
	if readError != nil {
		log.Fatal(readError)
	}
	parseError := yaml.Unmarshal(dataBytes, &s) //reads yaml and json because json is just a subset of yaml
	if parseError != nil {
		fmt.Println(fpath + ":1:1")
		log.Fatal(parseError)
	}

	// convert fields from yaml.Node to Field type
	for pair := s.Fields.Oldest(); pair != nil; pair = pair.Next() {
		fieldYamlNode := pair.Value.(*yaml.Node)
		var AstField *specSpec.Field
		fieldYamlNode.Decode(&AstField)
		pair.Value = AstField
	}

	return s
}

// Stores the spec to disc
func (a *TypeAst) Save(specDir string) {
	filepath := path.Dir(path.Join(specDir, a.Path, a.FileName))
	filename := path.Join(filepath, a.FileName)

	var d []byte
	var err error
	switch Format {
	case "json":
		d, err = a.ToJson()
		break
	case "yaml":
		d, err = a.ToYaml()
		break
	}
	if err != nil {
		panic(err)
	}
	util.MkdirRelative(filepath)
	err = ioutil.WriteFile(filename, d, 0644)
	if err != nil {
		panic(err)
	}
}

func (a *TypeAst) ToJson() (d []byte, err error) {
	d, err = json.MarshalIndent(a.TypeSpec, "", " ")
	return d, err
}

// returns unindented json
func (a *TypeAst) ToJsonFlat() (d []byte, err error) {
	d, err = json.Marshal(a.TypeSpec)
	return d, err
}

func (a *TypeAst) ToYaml() (d []byte, err error) {
	d, err = yaml.Marshal(&a.TypeSpec)
	return d, err
}

// stores the typelist to the spec directory
func (l *Typelist) SaveAllTypeSpecsToDir(specDir string) {
	for _, typeAst := range l.TypesByName {
		typeAst.Save(specDir)
	}
}

//
func (l *Typelist) ResolveProtoImportForType(fqTypeName string) (imp string, typeFound bool) {
	//check on installed and spec tpelist
	imp = ""
	if l.TypesByName[fqTypeName] == nil && l.InstalledTypesByName[fqTypeName] == nil {
		return imp, false
	}
	if l.TypesByName[fqTypeName] != nil {
		imp = l.TypesByName[fqTypeName].GetProtoTarget()
	}
	if l.InstalledTypesByName[fqTypeName] != nil {
		imp = l.InstalledTypesByName[fqTypeName].GetProtoTarget()
	}

	return imp, true
}

func (a *TypeAst) GetProtoTarget() (proto string) {
	protoFile := a.TypeSpec.XProto.Targetfile
	return path.Join(a.Path, protoFile)
}

// updates the imports on each type
func (l *Typelist) UpdateImports() {
	for t, v := range l.TypesByName {
		self, _ := l.ResolveProtoImportForType(t)
		imports := []string{}
		v.TypeSpec.Fields.Map(func(iFieldname interface{}, iField interface{}) {
			field := iField.(*specSpec.Field)
			fileToImport := field.Type
			// map imports
			if strings.HasPrefix(fileToImport, "map") {
				regex := regexp.MustCompile(`,([^>]*)`)
				matches := regex.FindStringSubmatch(fileToImport)
				fileToImport = strings.TrimSpace(matches[1])
			}

			// string, uint,... does not need to be imported
			f := strings.Split(fileToImport, ".")
			if len(f) > 1 {
				i, found := l.ResolveProtoImportForType(fileToImport)
				if found && i != self {
					imports = append(imports, i)
				} else {
					if i != self {
						fmt.Println(util.ScanForStringPosition(fileToImport, path.Join(viper.GetString("typeSpecDir"), l.TypesByName[t].Path, l.TypesByName[t].FileName))+":Import", fileToImport, "not found in type", t, "on field", iFieldname)
						fmt.Println(util.ScanForStringPosition(fileToImport, viper.GetString("muSpec.types"))+":Import not found. Check your muSpec types if you came from there. Field:", iFieldname)
					}
				}
			}

		})
		v.TypeSpec.XProto.Imports = imports
	}
}

// Deletes the spec from disk and removes the element from List
func (l *Typelist) DeleteType(typename string) {
	// delete the file
	filepath := path.Join(viper.GetString("typeSpecDir"), l.TypesByName[typename].Path, l.TypesByName[typename].FileName)
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DELETED", filepath)
	}

	delete(l.TypesByName, typename)

}
