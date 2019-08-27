package vertex

import "fmt"

var prefixes = [3]string{}

const (
	DUMB   = 0
	TOKEN  = 1
	STRING = 2
)

func init() {
	prefixes[DUMB] = "DUMB"
	prefixes[TOKEN] = "<TOKEN>"
	prefixes[STRING] = "<STRING>"
}

type (
	Vertex struct {
		kind        int
		name        string
		token       []rune
		transitions map[string]*Vertex
	}
)

func Dumb() *Vertex {
	return &Vertex{
		kind:        DUMB,
		token:       nil,
		transitions: map[string]*Vertex{},
	}
}
func Token(tk string) *Vertex {
	return &Vertex{
		kind:        TOKEN,
		token:       []rune(tk),
		transitions: map[string]*Vertex{},
	}
}
func String() *Vertex {
	return &Vertex{
		kind:        STRING,
		token:       nil,
		transitions: map[string]*Vertex{},
	}
}
func (v *Vertex) Kind() int {
	return v.kind
}
func (v *Vertex) Name() string {
	var name string
	switch v.kind {
	case DUMB:
		name = v.prefix()
	case TOKEN:
		name = fmt.Sprintf("%s:`%s`", v.prefix(), string(v.token))
	case STRING:
		name = fmt.Sprintf("%s:\"...\"", v.prefix())
	}
	return name
}
func (v *Vertex) prefix() string {
	return prefixes[v.kind]
}
func (v *Vertex) TransitsTo(vertices ...*Vertex) {
	for _, tv := range vertices {
		if v.hasTransitionTo(tv) {
			panic(fmt.Errorf("vertex `%s` already has transition to `%s`", v.Name(), tv.Name()))
		}

		v.transitions[tv.Name()] = tv
	}
}
func (v *Vertex) Transitions() []*Vertex {
	ts := []*Vertex{}
	for _, v := range v.transitions {
		ts = append(ts, v)
	}
	return ts
}
func (v *Vertex) hasTransitionTo(anotherV *Vertex) bool {
	_, has := v.transitions[anotherV.Name()]
	return has
}
