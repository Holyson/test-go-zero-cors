package template

import (
	"fmt"

	"github.com/Holyson/test-go-zero-cors/tools/goctl/util"
)

// ModelCustom defines a template for extension
const ModelCustom = `package {{.pkg}}
{{if .withCache}}
import (
	"github.com/Holyson/test-go-zero-cors/core/stores/cache"
	"github.com/Holyson/test-go-zero-cors/core/stores/sqlx"
)
{{else}}
import "github.com/Holyson/test-go-zero-cors/core/stores/sqlx"
{{end}}
var _ {{.upperStartCamelObject}}Model = (*custom{{.upperStartCamelObject}}Model)(nil)

type (
	// {{.upperStartCamelObject}}Model is an interface to be customized, add more methods here,
	// and implement the added methods in custom{{.upperStartCamelObject}}Model.
	{{.upperStartCamelObject}}Model interface {
		{{.lowerStartCamelObject}}Model
	}

	custom{{.upperStartCamelObject}}Model struct {
		*default{{.upperStartCamelObject}}Model
	}
)

// New{{.upperStartCamelObject}}Model returns a model for the database table.
func New{{.upperStartCamelObject}}Model(conn sqlx.SqlConn{{if .withCache}}, c cache.CacheConf{{end}}) {{.upperStartCamelObject}}Model {
	return &custom{{.upperStartCamelObject}}Model{
		default{{.upperStartCamelObject}}Model: new{{.upperStartCamelObject}}Model(conn{{if .withCache}}, c{{end}}),
	}
}
`

// ModelGen defines a template for model
var ModelGen = fmt.Sprintf(`%s

package {{.pkg}}
{{.imports}}
{{.vars}}
{{.types}}
{{.new}}
{{.delete}}
{{.find}}
{{.insert}}
{{.update}}
{{.extraMethod}}
{{.tableName}}
`, util.DoNotEditHead)
