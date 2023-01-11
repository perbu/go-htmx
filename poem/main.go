package poem

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//go:embed poem.txt
var poem string

type Poem struct {
	content  []string
	current  int
	previous int
}

func NewPoem() *Poem {

	// split the poem into lines:
	lines := strings.Split(poem, "\n")

	return &Poem{
		content: lines,
		current: 1,
	}
}

// Next returns the next line of the poem, and the previous line of the poem.
func (p *Poem) Next() (string, string) {
	p.current++
	if p.current >= len(p.content) {
		p.current = 1
	}

	next := p.content[p.current]
	previous := p.content[p.current-1]
	p.previous = p.current
	return next, previous
}

func (p *Poem) NextHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("poemHandler")
	// get the next line of the poem:
	next, previous := p.Next()
	// write the next line of the poem to the response:
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(fmt.Sprintf(`<div style="color:grey">%s</div><br>
<div>%s</div>`, previous, next)))
	log.Println("next:", next)
	log.Println("previous:", previous)

}
