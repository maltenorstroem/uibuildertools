package microtypes

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/theNorstroem/spectools/pkg/ast/typeAst"
	"github.com/theNorstroem/spectools/pkg/orderedmap"
	"github.com/theNorstroem/spectools/pkg/specSpec"
	"github.com/theNorstroem/spectools/pkg/specSpec/furo"
	"github.com/theNorstroem/spectools/pkg/util"
	"path"
	"regexp"
	"strconv"
	"strings"
)

type MicroTypelist struct {
	MicroTypesByName    map[string]*MicroType
	MicroTypesASTByName map[string]MicroTypeAst
	MicroTypes          []*MicroType `yaml:"types"`
}

func (l *MicroTypelist) UpateTypelist(typelist *typeAst.Typelist) {
	// build list to delete specs which are not types.yaml
	deleteList := map[string]bool{}
	for typeName, _ := range typelist.TypesByName {
		// mark every item as deletable
		deleteList[typeName] = true
	}

	for typename, mType := range l.MicroTypesASTByName {
		deleteList[typename] = false
		// create type on Typelist if it does not exist
		if typelist.TypesByName == nil {
			typelist.TypesByName = map[string]*typeAst.TypeAst{}
		}

		AstType, ok := typelist.TypesByName[typename]
		if !ok {
			typelist.TypesByName[typename] = &typeAst.TypeAst{
				Path:     mType.TargetPath,
				FileName: mType.Type + ".type.spec",
				TypeSpec: specSpec.Type{},
			}
			AstType = typelist.TypesByName[typename]
		}

		AstType.TypeSpec.Name = mType.Type
		AstType.TypeSpec.Type = mType.Type
		AstType.TypeSpec.Description = mType.Description
		if AstType.TypeSpec.XProto == nil {
			AstType.TypeSpec.XProto = &specSpec.Typeproto{
				Imports:    []string{},
				Options:    map[string]string{},
				Package:    "",
				Targetfile: "",
			}
		}

		AstType.TypeSpec.XProto.Package = mType.Package
		AstType.TypeSpec.XProto.Targetfile = mType.Target
		AstType.TypeSpec.XProto.Options["go_package"] = util.GetGoPackageName(mType.TargetPath)
		AstType.TypeSpec.XProto.Options["java_package"] = "com." + mType.Package
		AstType.TypeSpec.XProto.Options["java_outer_classname"] = strings.Title(strings.Replace(path.Base(mType.Target), ".proto", "Proto", 1))
		AstType.TypeSpec.XProto.Options["java_multiple_files"] = "true"

		fieldDeleteList := map[string]bool{}
		if AstType.TypeSpec.Fields != nil {
			AstType.TypeSpec.Fields.Map(func(iKey interface{}, iValue interface{}) {
				fieldDeleteList[iKey.(string)] = true
			})
		}
		for pair := mType.Fields.Oldest(); pair != nil; pair = pair.Next() {
			mFieldname := pair.Key.(string)
			mField := pair.Value.(FieldMap)
			// check for fields create if they do not exist and update later
			if AstType.TypeSpec.Fields == nil {
				AstType.TypeSpec.Fields = orderedmap.New()
			}
			var AstField *specSpec.Field
			afInterface, ok := AstType.TypeSpec.Fields.Get(mFieldname)

			// remove field from deletelist
			fieldDeleteList[mFieldname] = false

			if !ok {
				// check for moved name create
				ids := AstType.TypeSpec.Fields.Filter(func(iKey interface{}, iValue interface{}) bool {
					t := iValue.(*specSpec.Field)
					return t.XProto.Number == mField.FieldId
				})
				if len(ids) > 0 {
					// renamed field detected => copy from old fieldname
					AstField = ids[0].Value.(*specSpec.Field)
				} else {
					AstField = &specSpec.Field{
						XProto: &specSpec.Fieldproto{
							Number: mField.FieldId,
						},
						XUi:         nil,
						Constraints: map[string]*furo.FieldConstraint{},
						Description: "",
						Meta: &furo.FieldMeta{
							Default:  "",
							Hint:     "",
							Label:    strings.Join([]string{"label", mType.Package, mType.Type, mFieldname}, "."),
							Options:  &furo.Fieldoption{},
							Readonly: false,
							Repeated: false,
						},
						Type: "",
					}
				}

			} else {
				AstField = afInterface.(*specSpec.Field)
			}

			AstField.Type = mField.Type
			AstField.Description = mField.Description
			AstField.XProto.Number = mField.FieldId

			// check for __proto

			if mField.Required {
				AstField.Constraints["required"] = &furo.FieldConstraint{
					Is:      "true",
					Message: mFieldname + " is required",
				}
			} else {
				// remove if setted
				delete(AstField.Constraints, "required")
			}

			AstField.Meta.Default = mField.DefaultValue
			AstField.Meta.Readonly = mField.Readonly
			AstField.Meta.Repeated = mField.Repeated

			// Assign to Node again
			AstType.TypeSpec.Fields.Set(mFieldname, AstField)
		}

		for fieldname, del := range fieldDeleteList {
			if del {
				AstType.TypeSpec.Fields.Delete(fieldname)
			}
		}
	}
	// delete the item
	for typename, del := range deleteList {
		if del {
			typelist.DeleteType(typename)
		}
	}
}

// holds a single type from microspec
type MicroType struct {
	Type   string                 `yaml:"type"`
	Fields *orderedmap.OrderedMap `yaml:"fields,omitempty"`
	Target string                 `yaml:"target,omitempty"`
}

func (mt MicroType) ToMicroTypeAst() MicroTypeAst {

	// make the fieldmaps
	fields := orderedmap.New() //was map[string]FieldMap{}
	for pair := mt.Fields.Oldest(); pair != nil; pair = pair.Next() {
		fieldstring := pair.Value.(string)
		fieldName := pair.Key.(string)
		field := NewFieldMap()
		field.ParseFieldString(fieldstring)
		fields.Set(fieldName, field)
	}

	// parse title and description
	regex := regexp.MustCompile(`^([^#(]*):?([^#]*)?(#(.*))?$`)
	matches := regex.FindStringSubmatch(mt.Type)
	if len(matches) == 0 {
		fmt.Println("typeline not parseable", mt.Type)
	}

	typedescription := ""
	if matches[4] != "" {
		typedescription = strings.TrimSpace(matches[4])
	}
	typeName := strings.TrimSpace(matches[1])
	typeArr := strings.Split(typeName, ".")
	targetpath := strings.ToLower(typeArr[0])
	packagename := strings.ToLower(typeArr[0])
	targetname := strings.ToLower(typeArr[0]) + ".proto"
	if len(typeArr) > 1 {
		// last segement is typename
		typeName = strings.TrimSpace(typeArr[len(typeArr)-1])
		// the other segments are the package
		packagename = strings.Join(typeArr[:len(typeArr)-1], ".")
		targetpath = strings.Join(typeArr[:len(typeArr)-1], "/")
		// target file name
		targetname = typeArr[len(typeArr)-2] + ".proto"
		if mt.Target != "" {
			// if optional target was given
			targetname = mt.Target
		}
	}

	mAst := MicroTypeAst{
		Type:        typeName,
		Package:     packagename,
		TargetPath:  targetpath,
		Description: typedescription,
		Fields:      fields,
		Target:      targetname,
	}

	return mAst
}

func NewFieldMap() FieldMap {
	return FieldMap{Readonly: false, Required: false, Repeated: false, Type: "string", DefaultValue: "", Description: "no description", FieldId: 1}
}

type MicroTypeAst struct {
	Type           string                 `yaml:"type"`
	ProtoImports   []string               `yaml:"imports"`
	Package        string                 `yaml:"package,omitempty"`
	TargetPath     string                 // to find out the file to write
	Description    string                 `yaml:"description"`
	Fields         *orderedmap.OrderedMap `yaml:"fields,omitempty"`
	Target         string                 `yaml:"target,omitempty"`
	MakeEntity     bool
	makeCollection bool
}

// updates a type ast
func (mAst MicroTypeAst) UpdateTypeAst(ast typeAst.TypeAst) {

}

// field string will be converted to this type
// this type will be converted to fieldmap
type FieldMap struct {
	Readonly     bool
	Required     bool
	Repeated     bool
	Type         string
	DefaultValue string
	Description  string
	FieldId      int32
}

func (m *FieldMap) ParseFieldString(s string) {
	regex := regexp.MustCompile(`^(-*)?(\**)?(\[.?])?([^#=:]*):?([^=#]*)(=([^#]*))?(#(.*))?$`)
	matches := regex.FindStringSubmatch(s)
	if len(matches) == 0 {
		fmt.Println("field not parsed", s)
		return
	}
	if matches[1] != "" {
		m.Readonly = true
	}
	if matches[2] != "" {
		m.Required = true
	}
	if matches[3] != "" {
		m.Repeated = true
	}
	if matches[4] != "" {
		m.Type = strings.TrimSpace(matches[4])
	}
	if matches[5] != "" {
		n, _ := strconv.Atoi(strings.TrimSpace(matches[5]))
		m.FieldId = int32(n)
	} else {
		fmt.Println(util.ScanForStringPosition(s, viper.GetString("muSpec.types"))+":Field numbers must be positive integers", s)
	}
	if matches[7] != "" {
		m.DefaultValue = strings.TrimSpace(matches[7])
	}
	if matches[9] != "" {
		m.Description = strings.TrimSpace(matches[9])
	}
}
