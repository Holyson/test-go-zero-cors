package gen

import (
	"strings"

	"github.com/Holyson/test-go-zero-cors/tools/goctl/model/sql/parser"
	"github.com/Holyson/test-go-zero-cors/tools/goctl/model/sql/template"
	"github.com/Holyson/test-go-zero-cors/tools/goctl/util"
	"github.com/Holyson/test-go-zero-cors/tools/goctl/util/pathx"
)

func genFields(table Table, fields []*parser.Field) (string, error) {
	var list []string

	for _, field := range fields {
		result, err := genField(table, field)
		if err != nil {
			return "", err
		}

		list = append(list, result)
	}

	return strings.Join(list, "\n"), nil
}

func genField(table Table, field *parser.Field) (string, error) {
	tag, err := genTag(table, field.NameOriginal)
	if err != nil {
		return "", err
	}

	text, err := pathx.LoadTemplate(category, fieldTemplateFile, template.Field)
	if err != nil {
		return "", err
	}

	output, err := util.With("types").
		Parse(text).
		Execute(map[string]interface{}{
			"name":       util.SafeString(field.Name.ToCamel()),
			"type":       field.DataType,
			"tag":        tag,
			"hasComment": field.Comment != "",
			"comment":    field.Comment,
			"data":       table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}