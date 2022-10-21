package {{ .MODULE_TYPE | ToLower }}

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/brandonclapp/budget/core/data"
)

type {{.MODULE_TYPE }} struct {
	ID    string  `json:"id"`
	{{ range $val := .MODULE_STRUCT_PROPS -}}
	{{.name}} {{$val.type}} `json:"{{$val.name | ToLowerCamel }}"`
	{{ end }}
}

