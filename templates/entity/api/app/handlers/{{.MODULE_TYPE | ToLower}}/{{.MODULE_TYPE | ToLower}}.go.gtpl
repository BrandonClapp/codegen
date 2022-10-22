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

{{ $propListLen := len .MODULE_STRUCT_PROPS }}
var {{.MODULE_TYPE | Pluralize}} = data.Queryable[{{ .MODULE_TYPE }}]{
	Table:   "public.{{ .MODULE_TYPE | ToLower }}",
	Columns: `id, 
	{{- range $i, $val := .MODULE_STRUCT_PROPS -}}
		{{ if eq $i 0 }} {{ end }}"{{ $val.name | ToSnake }}"{{ if not (IsLast $i $propListLen) }}, {{ end }}
	{{- end }}`,
	ScanFn: func(rows *sql.Rows) (*{{ .MODULE_TYPE }}, error) {
		var e = &{{ .MODULE_TYPE }}{}
		err := rows.Scan(&e.ID, 
		{{- range $i, $val := .MODULE_STRUCT_PROPS -}}
		{{ if eq $i 0 }} {{ end }}&e.{{ $val.name }}{{ if not (IsLast $i $propListLen) }}, {{ end }}
		{{- end }})
		return e, err
	},
	InsertFn: func(e *{{.MODULE_TYPE}}) []interface{} {
		values := []interface{}{sq.Expr("DEFAULT"), 
		{{- range $i, $val := .MODULE_STRUCT_PROPS -}}
		{{ if eq $i 0 }} {{ end }}&e.{{ $val.name }}{{ if not (IsLast $i $propListLen) }}, {{ end }}
		{{- end }}}
		return values
	},
	Seed: `
	create table if not exists {{ .MODULE_TYPE | ToLower  }}
	(
		id serial
			constraint {{ .MODULE_TYPE | ToLower  }}_pk
				primary key,
		{{ range $i, $val := .MODULE_STRUCT_PROPS -}}
		{{ $val.name | ToSnake }} {{ $val.type | ToPostgresType }}{{ if not (IsLast $i $propListLen) }}, {{ end }}
		{{- end }}
	);
	
	create unique index if not exists {{ .MODULE_TYPE | ToLower  }}_id_uindex
		on {{ .MODULE_TYPE | ToLower  }} (id);
	`,
}


