package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type Composer struct {
	Name, ImgURL string
}

type Composers []Composer

func (c Composers) Composer() Composer {
	return c[rand.Intn(len(c))]
}

var (
	host = flag.String(`addr`, ``, `tcp listen address`)
	port = flag.Int(`port`, 8080, `tcp listen port`)

	composers = Composers{
		{`Johann Sebastian Bach`, `http://historylists.org/images/johann-sebastian-bach.jpg`},
		{`Wolfgang Amadeus Mozart`, `http://historylists.org/images/wolfgang-amadeus-mozart.jpg`},
		{`Ludwig van Beethoven`, `http://historylists.org/images/ludwig-van-beethoven.jpg`},
		{`Giuseppe Verdi`, `http://historylists.org/images/giuseppe-verdi.jpg`},
		{`Pyotr Ilyich Tchaikovsky`, `http://historylists.org/images/pyotr-ilyich-tchaikovsky.jpg`},
		{`Frederic Chopin`, `http://historylists.org/images/frederic-chopin.jpg`},
		{`Antonio Vivaldi`, `http://historylists.org/images/antonio-vivaldi.jpg`},
		{`Giacomo Puccini`, `http://historylists.org/images/giacomo-puccini.jpg`},
		{`George Frideric Handel`, `http://historylists.org/images/george-frideric-handel.jpg`},
		{`Igor Stravinsky`, `http://historylists.org/images/igor-stravinsky.jpg`},
	}

	composerHTML = `<!DOCTYPE html>
<html>
  <body>
    <h1>{{.Name}}</h1>
    <img src='{{.ImgURL}}'/>
  </body>
</html>`
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	if err := run(composers.Composer(), *host, *port); err != nil {
		panic(err)
	}
}

func run(c Composer, host string, port int) error {
	composerTempl, err := template.New(`composer`).Parse(composerHTML)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		if err := composerTempl.Execute(&buf, c); err != nil {
			panic(err)
		}

		fmt.Fprintf(w, buf.String())
	})

	return http.ListenAndServe(strings.Join([]string{host, strconv.Itoa(port)}, ":"), nil)
}
