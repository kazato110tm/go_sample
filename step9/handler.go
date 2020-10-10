package main

import (
	"html/template"
	"net/http"
)

type Handlers struct {
	ab *AccountBook
}

func NewHandlers(ab *AccountBook) *Handlers {
	return &Handlers{ab: ab}
}

var listTmpl = template.Must(template.New("list").Parse(`<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8"/>
	<title>家計簿</title>
</head>
<body>
	<h1>家計簿</h1>
	<h2>最新{{len .}}件</h2>
	{{- if . -}}
	<table border="1">
	<tr><th>品目</th><th>値段<th></tr>
		{{- range .}}
		<tr><td>{{.Category}}</td><td>{{.Price}}円</td></tr>
		{{- end}}
	</table>
	{{- else}}
	データがありません
	{{- end}}
</body>
</html>
`))

func (hs *Handlers) ListHandler(w http.ResponseWriter, r *http.Request) {
	items, err := hs.ab.GetItems(10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := listTmpl.Execute(w, items); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
