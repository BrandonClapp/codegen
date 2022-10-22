# Codegen

`codegen` is a command line utility that transforms [Go templates](<(https://pkg.go.dev/text/template)>) from a provided input directory, using variables defined in a configuration file (`gen.config.json`), and then writes the output to a provided output directory.

Go is not required to be installed, just that the templates are written using Go template engine syntax and use the `.gtpl` extension.

## gen.config.json

`gen.config.json` is the configuration file used by the `codegen` CLI and **must** be present in the directory in which `codegen` is executed.

### Schema

- `templates` (array) \*required - Defines all templates that can be transformed
  - `name` (string) \*required - The unique name of the template
  - `inDir` (string) \*required - The directory where the template files reside
  - `ourDir` (string) \*required - The directory where the transformed template will be output
  - `variables` (object) - Variables that will be injected into each template file

Below is an example of what `gen.config.json` may look like.

```json
{
  "templates": [
    {
      "name": "entity",
      "inDir": "../templates/entity",
      "outDir": ".",
      "variables": {
        "MODULE_TYPE": "Payment",
        "MODULE_STRUCT_PROPS": [
          {
            "name": "Title",
            "type": "string"
          },
          {
            "name": "Amount",
            "type": "int"
          },
          {
            "name": "Date",
            "type": "string"
          }
        ]
      }
    }
  ]
}
```

Variables can then be used in templates like so:

```gotpl
type {{ .MODULE_TYPE }} struct {
	ID    string  `json:"id"`
	{{ range $val := .MODULE_STRUCT_PROPS -}}
	{{.name}} {{$val.type}} `json:"{{$val.name | ToLowerCamel }}"`
	{{ end }}
}
```

## Template Helper Functions

Some helper functions are provided to make transforming variables easier.

### String helpers

| Function             | Input             | Output             | Usage                                 |
| -------------------- | ----------------- | ------------------ | ------------------------------------- |
| **ToLower**          | AnyKind_of-STRING | anykindofstring    | `{{ .MyString \| ToLower }}`          |
| **ToKebab**          | AnyKind_of-STRING | any-kind-of-string | `{{ .MyString \| ToKebab }}`          |
| **ToScreamingKebab** | AnyKind_of-STRING | ANY-KIND-OF-STRING | `{{ .MyString \| ToScreamingKebab }}` |
| **ToSnake**          | AnyKind_of-STRING | any_kind_of_string | `{{ .MyString \| ToSnake }}`          |
| **ToScreamingSnake** | AnyKind_of-STRING | ANY_KIND_OF_STRING | `{{ .MyString \| ToScreamingSnake }}` |
| **ToCamel**          | AnyKind_of-STRING | AnyKindOfString    | `{{ .MyString \| ToCamel }}`          |
| **ToLowerCamel**     | AnyKind_of-STRING | anyKindOfString    | `{{ .MyString \| ToLowerCamel }}`     |
| **Pluralize**        | Person            | People             | `{{ .MyString \| Pluralize }}`        |

### Array helpers

**IsLast**

```
{{ $propListLen := len .MODULE_STRUCT_PROPS }}
{{- range $i, $val := .MODULE_STRUCT_PROPS -}}
		{{ if eq $i 0 }} {{ end }}"{{ $val.name | ToSnake }}"{{ if not (IsLast $i $propListLen) }}, {{ end }}
{{- end }}`,
```
