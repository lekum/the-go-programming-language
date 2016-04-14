package main

import (
	"gopl.io/ch4/github"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
URL:
{{.HTMLURL}}
User:
{{.User.Login}}
Title: {{.Title | printf "%.64s" | toUpper}}
Age:
{{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo, "toUpper": strings.ToUpper}).Parse(templ))

func main() {
	res, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, res); err != nil {
		log.Fatal(err)
	}
}
