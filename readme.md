# Codegen

**[Demo](https://www.youtube.com/watch?v=QxpDWy37Z2Y)**

`codegen` is a command line utility that transforms [Go templates](<(https://pkg.go.dev/text/template)>) from a provided input directory, using variables defined in a configuration file (`gen.config.json`), and then writes the output to a provided output directory.

Go is not required to be installed, just that the templates are written using Go template engine syntax and use the `.gtpl` extension.

---

### gen.config.json

`gen.config.json` is the configuration file used by the `codegen` CLI and **must** be present in the directory in which `codegen` is executed and defines information about all template directories, output, and injected variables.

**Properties**

- `templates` (array) \*required - Defines all templates that can be transformed
  - `name` (string) \*required - The unique name of the template
  - `inDir` (string) \*required - The directory where the template files reside
  - `ourDir` (string) \*required - The directory where the transformed template will be output
  - `variables` (object) - Variables that will be injected into each template file

---

## Example

**gen.config.json**

```json
{
  "templates": [
    {
      "name": "interface",
      "inDir": "../../templates/interface",
      "outDir": "../../sandbox",
      "variables": {
        "TYPE": "UpcomingMovie",
        "PROPS": [
          {
            "name": "Name",
            "type": "string"
          },
          {
            "name": "Rating",
            "type": "string"
          },
          {
            "name": "LengthInMinutes",
            "type": "number"
          },
          {
            "name": "ReleaseDate",
            "type": "Date"
          }
        ]
      }
    }
  ]
}
```

Command (ran from directory containing config file):

```
codegen interface
```

Reads: `../templates/interface/models/{{.TYPE | ToLowerCamel}}.ts.gtpl`

```
export interface {{.TYPE | ToCamel }} {
    {{ range $val := .PROPS -}}
	{{$val.name | ToLowerCamel }}: {{$val.type}};
	{{ end }}
}
```

Outputs: `./models/upcomingMovie.ts`

```ts
export interface UpcomingMovie {
  name: string;
  rating: string;
  lengthInMinutes: number;
  releaseDate: Date;
}
```

---

## Template Helper Functions

Some helper functions are provided to make transforming variables easier.

### String helpers

| Function             | Input             | Output             | Usage                               |
| -------------------- | ----------------- | ------------------ | ----------------------------------- |
| **ToLower**          | AnyKind_of-STRING | anykindofstring    | `{{.MyString \| ToLower}}`          |
| **ToKebab**          | AnyKind_of-STRING | any-kind-of-string | `{{.MyString \| ToKebab}}`          |
| **ToScreamingKebab** | AnyKind_of-STRING | ANY-KIND-OF-STRING | `{{.MyString \| ToScreamingKebab}}` |
| **ToSnake**          | AnyKind_of-STRING | any_kind_of_string | `{{.MyString \| ToSnake}}`          |
| **ToScreamingSnake** | AnyKind_of-STRING | ANY_KIND_OF_STRING | `{{.MyString \| ToScreamingSnake}}` |
| **ToCamel**          | AnyKind_of-STRING | AnyKindOfString    | `{{.MyString \| ToCamel}}`          |
| **ToLowerCamel**     | AnyKind_of-STRING | anyKindOfString    | `{{.MyString \| ToLowerCamel}}`     |
| **Pluralize**        | Person            | People             | `{{.MyString \| Pluralize}}`        |

### Array helpers

**IsLast** - determine if an index is the last index in an array

Example:

```
{{ $arrayLength := len .SomeArrayVariable }}
{{- range $i, $val := .SomeArrayVariable -}}
  {{- if (IsLast $i $arrayLength) -}}
    {{$val}} is the last element
  {{ end }}
{{- end }}
```

**Misc Notes**

- Templates may contain nested folders. The output will retain this folder structure relative to the `outDir`
- Template file and folder names may contain template syntax, i.e `{{ .SomeVariable | ToSnake }}.py`
