package {{MODULE_NAME}}

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/brandonclapp/budget/core/data"
)

type {{MODULE_TYPE}} struct {
	ID    string  `json:"id"`
	{{MODULE_STRUCT_PROPS}}
}

// Accessed by {{MODULE_TYPE}}s.Get(&Filters{})
var {{MODULE_TYPE_PLURAL}} = data.Queryable[{{MODULE_TYPE}}]{
	Table:   "public.{{MODULE_NAME}}",
	Columns: `id, {{MODULE_COLUMNS}}`,
	ScanFn: func(rows *sql.Rows) (*{{MODULE_TYPE}}, error) {
		var e = &{{MODULE_TYPE}}{}
		err := rows.Scan(&e.ID, {{MODULE_SQL_PROP_MAPPINGS}})
		return e, err
	},
	InsertFn: func(e *{{MODULE_TYPE}}) []interface{} {
		values := []interface{}{sq.Expr("DEFAULT"), {{MODULE_SQL_PROP_MAPPINGS}}}
		return values
	},
	Seed: `
	create table if not exists {{MODULE_NAME}}
	(
		id serial
			constraint {{MODULE_NAME}}_pk
				primary key,
		{{MODULE_COLUMNS_SQL}}
	);
	
	create unique index if not exists {{MODULE_NAME}}_id_uindex
		on {{MODULE_NAME}} (id);
	`,
}
