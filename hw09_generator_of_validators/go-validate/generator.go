package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type TemplateField struct {
	Name   string
	Type   string
	Regexp string
	Len    string
	Max    string
	Min    string
	In     string
}

type TemplateStruct struct {
	Name   string
	Fields []TemplateField
}

var validate = `func (t {{ .Name }}) Validate() ([]ValidationError, error) {
  var validErrs []ValidationError
  {{ range .Fields}}

  {{- if (eq .Type "int")}}
  {{ template "intValidate" . }}
  {{- end }}

  {{- if (eq .Type "string")}}
  {{ template "stringValidate" . }}
  {{- end }}

  {{- if (eq .Type "[]int")}}
  for i, _ := range t.{{ .Name }} {
    {{ $new := arrayElement . }}
    {{ template "stringValidate" $new }}
  }
  {{- end }}

  {{- if (eq .Type "[]string")}}
  for i, _ := range t.{{ .Name }} {
    {{ $new := arrayElement . }}
    {{ template "stringValidate" $new }}
  }
  {{- end }}

  {{ end }}
  return validErrs, nil
}
`

var intValidate = `{{- if (ne .Min "")}}

if t.{{.Name}} < {{.Min}} {
  validErrs = append(validErrs, ValidationError{
    Field: "{{.Name}}",
    Err: fmt.Errorf("%d should be more than %d", t.{{.Name}}, {{.Min}}),
  })
}

{{- end }}

{{- if (ne .Max "")}}

if t.{{.Name}} > {{.Max}} {
  validErrs = append(validErrs, ValidationError{
    Field: "{{.Name}}",
    Err: fmt.Errorf("%d should be less than %d", t.{{.Name}}, {{.Max}}),
  })
}

{{- end }}

{{- if (ne .In "")}}

intVariants := []int{ {{.In}} }
isVariant := false
for _, i := range intVariants {
  if t.{{.Name}} == i {
    isVariant = true
    break
  }
}
if !isVariant {
  validErrs = append(validErrs, ValidationError{
    Field: "{{.Name}}",
    Err: fmt.Errorf("not allowed value %d", t.{{.Name}}),
  })
}

{{- end }}
`

var stringValidate = `{{- if (ne .In "")}}

strVariants := []string{ {{ .In }} }
isStrVariant := false
for _, i := range strVariants {
  if string(t.{{.Name}}) == i {
    isStrVariant = true
    break
  }
}
if !isStrVariant {
  validErrs = append(validErrs, ValidationError{
    Field: "{{.Name}}",
    Err: fmt.Errorf("not allowed value %s", t.{{.Name}}),
  })
}

{{- end }}

{{- if (ne .Len "")}}

if len(t.{{.Name}}) != {{.Len}} {
  validErrs = append(validErrs, ValidationError{
    Field: "{{.Name}}",
    Err: fmt.Errorf("wrong length: expected %d, got %d", {{.Len}}, len(t.{{.Name}})),
  })
}

{{- end }}

{{- if (ne .Regexp "")}}

matched, err := regexp.MatchString({{ .Regexp }}, t.{{.Name}})
if err != nil {
  return validErrs, fmt.Errorf("wrong regexp")
}
if !matched {
  validErrs = append(validErrs, ValidationError{
    Field: "{{.Name}}",
    Err: fmt.Errorf("%s does not match regexp", t.{{.Name}}),
  })
}

{{- end }}
`

func generate(path string, validateMap map[string][]item) error {
	t, err := template.New("validate").Funcs(template.FuncMap{
		"arrayElement": func(field TemplateField) TemplateField {
			field.Type = strings.Split(field.Type, "[]")[1]
			field.Name += "[i]"
			return field
		},
	}).Parse(validate)
	if err != nil {
		return err
	}

	_, err = t.New("intValidate").Parse(intValidate)
	if err != nil {
		return err
	}

	_, err = t.New("stringValidate").Parse(stringValidate)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal("can't close file", err)
		}
	}()

	beginning := `// Code generated by go-validate. DO NOT EDIT.
package models

import (
  "fmt"
  "regexp"
)

type ValidationError struct{
  Field string
  Err   error
}

`
	_, err = file.Write([]byte(beginning))
	if err != nil {
		return err
	}
	tempStruct := prepareStruct(validateMap)
	for _, v := range tempStruct {
		err = t.ExecuteTemplate(file, "validate", v)
		if err != nil {
			return err
		}
	}
	return nil
}

func prepareStruct(tmpMap map[string][]item) []TemplateStruct {
	tempStruct := []TemplateStruct{}
	for _, s := range tmpMap {
		tagRegexp := ""
		tagLen := ""
		tagMax := ""
		tagMin := ""
		tagIn := ""
		tempFields := []TemplateField{}

		for _, f := range s {
			tagRegexp = ""
			tagLen = ""
			tagMax = ""
			tagMin = ""
			tagIn = ""
			tags := strings.Split(f.fieldTags, " ")
			for k, v := range tags {
				quote := ""
				if f.fieldType == "string" {
					quote = "\""
				}
				switch v {
				case "regexp":
					tagRegexp = "\"" + tags[k+1] + "\""
				case "len":
					tagLen = tags[k+1]
				case "max":
					tagMax = tags[k+1]
				case "min":
					tagMin = tags[k+1]
				case "in":
					tags = tags[k+1:]
					for i := range tags {
						tags[i] = quote + tags[i] + quote
					}
					tagIn = strings.Join(tags, ",")
				default:
					continue
				}
			}
			tempFields = append(tempFields, TemplateField{
				Name:   f.fieldName,
				Type:   f.fieldType,
				Regexp: tagRegexp,
				Len:    tagLen,
				Max:    tagMax,
				Min:    tagMin,
				In:     tagIn,
			})
		}
		tempStruct = append(tempStruct, TemplateStruct{
			Name:   s[0].structName,
			Fields: tempFields,
		})
	}
	return tempStruct
}
