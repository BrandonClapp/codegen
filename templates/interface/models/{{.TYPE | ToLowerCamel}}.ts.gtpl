export interface {{.TYPE | ToCamel }} {
    {{ range $val := .PROPS -}}
	{{$val.name | ToLowerCamel }}: {{$val.type}};
	{{ end }}
}