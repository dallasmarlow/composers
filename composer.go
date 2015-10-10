package composers

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

var (
	composerHTML = `<!DOCTYPE html>
<html>
  <body>
    <h1>{{.Name}}</h1>
    <img src='{{.ImgURL}}'/>
  </body>
</html>`
)

type Composer struct {
	Name, ImgURL string
}

func (c Composer) Listen(addr string) error {
	templ, err := template.New(c.Name).Parse(composerHTML)
	if err != nil {
		return err
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		if err := templ.Execute(&buf, c); err != nil {
			panic(err)
		}

		fmt.Fprintf(w, buf.String())
	})

	return http.ListenAndServe(addr, nil)
}