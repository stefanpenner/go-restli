package codegen

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/PapaCharlie/go-restli/protocol"

	. "github.com/dave/jennifer/jen"
	"github.com/pkg/errors"
)

const (
	EncodingJson  = "encoding/json"
	Unmarshal     = "Unmarshal"
	UnmarshalJSON = "UnmarshalJSON"
	Marshal       = "Marshal"
	MarshalJSON   = "MarshalJSON"

	Codec                = "codec"
	RestLiEncode         = "RestLiEncode"
	RestLiDecode         = "RestLiDecode"
	RestLiCodec          = "RestLiCodec"
	RestLiUrlEncoder     = "RestLiUrlEncoder"
	RestLiReducedEncoder = "RestLiReducedEncoder"

	PopulateDefaultValues = "populateDefaultValues"
	ValidateUnionFields   = "validateUnionFields"

	NetHttp = "net/http"

	ProtocolPackage = "github.com/PapaCharlie/go-restli/protocol"
)

var (
	packagePrefix *string

	CommentWrapWidth = 120

	HeaderTemplate = template.Must(template.New("header").Parse(`DO NOT EDIT

Code automatically generated by go-restli
Source files:
{{range .SourceFilenames}}- {{.}}
{{end}}`))
)

type CodeFile struct {
	SourceFilenames []string
	PackagePath     string
	Filename        string
	Code            *Statement
}

func NewCodeFile(filename string, packageSegments ...string) *CodeFile {
	return &CodeFile{
		PackagePath: filepath.Join(packageSegments...),
		Filename:    filename,
		Code:        Empty(),
	}
}

func (f *CodeFile) Write(outputDir string) (filename string, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = errors.Errorf("Could not generate model: %+v", e)
		}
	}()
	file := NewFilePath(f.PackagePath)

	header := bytes.NewBuffer(nil)
	err = HeaderTemplate.Execute(header, f)
	if err != nil {
		return "", err
	}
	file.HeaderComment(header.String())

	file.Add(f.Code)
	filename = filepath.Join(outputDir, f.PackagePath, f.Filename+".go")
	err = Write(filename, file)
	return filename, err
}

func (f *CodeFile) Identifier() string {
	return f.PackagePath + "." + f.Filename
}

func Write(filename string, file *File) error {
	b := bytes.NewBuffer(nil)
	if err := file.Render(b); err != nil {
		return errors.WithStack(err)
	}

	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return errors.WithStack(err)
	}

	_ = os.Remove(filename)

	if _, err := os.Stat(filename); err == nil {
		if removeErr := os.Remove(filename); removeErr != nil {
			return errors.WithMessagef(removeErr, "Could not delete %s", filename)
		}
	} else {
		if !os.IsNotExist(err) {
			return errors.WithStack(err)
		}
	}

	if err := ioutil.WriteFile(filename, b.Bytes(), os.FileMode(0555)); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func AddWordWrappedComment(code *Statement, comment string) *Statement {
	if comment != "" {
		code.Comment(comment)
		return code
	} else {
		return code
	}

	// WIP: RestLi comments are not behaving quite as expected, so comments get added as is, without being wrapped
	for len(comment) > CommentWrapWidth {
		if newline := strings.Index(comment[:CommentWrapWidth], "\n"); newline != -1 {
			code.Comment(comment[:newline]).Line()
			comment = comment[newline+1:]
			continue
		}

		if index := strings.LastIndexFunc(comment[:CommentWrapWidth], unicode.IsSpace); index > 0 {
			code.Comment(comment[:index]).Line()
			comment = comment[index+1:]
		} else {
			break
		}
	}

	code.Comment(comment)

	return code
}

func ExportedIdentifier(identifier string) string {
	return strings.ToUpper(identifier[:1]) + identifier[1:]
}

func PrivateIdentifier(identifier string) string {
	return strings.ToLower(identifier[:1]) + identifier[1:]
}

func ReceiverName(typeName string) string {
	return PrivateIdentifier(typeName[:1])
}

func AddFuncOnReceiver(def *Statement, receiver, typeName, funcName string) *Statement {
	return def.Func().
		Params(Id(receiver).Op("*").Id(typeName)).
		Id(funcName)
}

func AddMarshalJSON(def *Statement, receiver, typeName string, f func(def *Group)) *Statement {
	return AddFuncOnReceiver(def, receiver, typeName, MarshalJSON).
		Params().
		Params(Id("data").Index().Byte(), Err().Error()).
		BlockFunc(f)
}

func AddUnmarshalJSON(def *Statement, receiver, typeName string, f func(def *Group)) *Statement {
	return AddFuncOnReceiver(def, receiver, typeName, UnmarshalJSON).
		Params(Id("data").Index().Byte()).
		Params(Err().Error()).
		BlockFunc(f)
}

func AddRestLiEncode(def *Statement, receiver, typeName string, f func(def *Group)) *Statement {
	return AddFuncOnReceiver(def, receiver, typeName, RestLiEncode).
		Params(Id(Codec).Qual(ProtocolPackage, RestLiCodec)).
		Params(Id("data").String(), Err().Error()).
		BlockFunc(f)
}

func AddRestLiDecode(def *Statement, receiver, typeName string, f func(def *Group)) *Statement {
	return AddFuncOnReceiver(def, receiver, typeName, RestLiDecode).
		Params(Id(Codec).Qual(ProtocolPackage, RestLiCodec), Id("data").String()).
		Params(Err().Error()).
		BlockFunc(f)
}

func AddStringer(def *Statement, receiver, typeName string, f func(def *Group)) *Statement {
	return AddFuncOnReceiver(def, receiver, typeName, "String").
		Params().
		String().
		BlockFunc(f)
}

func IfErrReturn(def *Group, results ...Code) *Group {
	def.If(Err().Op("!=").Nil()).Block(Return(results...))
	return def
}

func SetPackagePrefix(prefix string) {
	if packagePrefix == nil {
		packagePrefix = &prefix
	} else {
		log.Panicln("packagePrefix can only be set once!")
	}
}

func GetPackagePrefix() string {
	if packagePrefix == nil {
		log.Panicln("packagePrefix not set!")
	}
	return *packagePrefix
}

func Bytes() *Statement {
	return Qual(ProtocolPackage, "Bytes")
}

type FieldTag struct {
	Json struct {
		Name     string
		Optional bool
	}
}

func (f *FieldTag) ToMap() map[string]string {
	tags := map[string]string{}
	if f.Json.Name != "" {
		tags["json"] = f.Json.Name
		if f.Json.Optional {
			tags["json"] += ",omitempty"
		}
	}

	return tags
}

func RestLiMethod(method protocol.RestLiMethod) *Statement {
	if method == protocol.NoMethod {
		return Qual(ProtocolPackage, "NoMethod")
	} else {
		return Qual(ProtocolPackage, "Method"+string(method))
	}
}
