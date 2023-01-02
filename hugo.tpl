{{block "album" . }}

Album {{.Title}}
{{ $owner := .Owner }}
{{ $albumid := .Id }}
{{range .Photos }}
{{ "{{<" }} figure src="https://live.staticflickr.com/{{ .Server }}/{{ .Id }}_{{ .Secret}}_b.jpg" link="https://www.flickr.com/photos/{{$owner}}/in/album-{{$albumid}}/" target="_new" caption="{{.Title}}" {{ ">}}" }}
{{end}}

{{end}}