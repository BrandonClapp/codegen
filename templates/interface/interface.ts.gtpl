export interface {{.TYPE}} {
    {{ range $val := .PROPS -}}
	{{$val.name | ToLowerCamel }}: {{$val.type}};
	{{ end }}
}