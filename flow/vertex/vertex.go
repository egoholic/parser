package vertex

import "fmt"

const STRV_NAME = "<STRING>"

type (
	Vertex struct {
		Name        string
		token       []rune
		transitions map[string]*Vertex
	}
)

func Token(tk string) *Vertex {
	return &Vertex{
		Name:        tk,
		token:       []rune(tk),
		transitions: map[string]*Vertex{},
	}
}
func String() *Vertex {
	return &Vertex{
		Name:        STRV_NAME,
		token:       nil,
		transitions: map[string]*Vertex{},
	}
}
func (v *Vertex) TransitsTo(vertices ...*Vertex) {
	for _, tv := range vertices {
		if v.hasTransitionTo(tv) {
			panic(fmt.Errorf("vertex `%s` already has transition to `%s`", v.Name, tv.Name))
		}

		v.transitions[tv.Name] = tv
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
	_, has := v.transitions[anotherV.Name]
	return has
}
