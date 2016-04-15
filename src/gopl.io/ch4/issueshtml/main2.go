package main

import (
	"gopl.io/ch4/github"
	"html/template"
	"log"
	"os"
)

const templ = `<h1>{{.TotalCount}} issues:</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
{{end}}
</tr>
</table>`

var report = template.Must(template.New("issuelist").Parse(templ))

func main() {
	res, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, res); err != nil {
		log.Fatal(err)
	}
}
