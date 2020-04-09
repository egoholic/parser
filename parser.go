package parser

import (
	"go/ast"
	"sync"

	"github.com/egoholic/parser/grammar/graph/vertex"
)

type Parser struct {
	grammar *vertex.Vertex
}

func New(v *vertex.Vertex) *Parser {
	return &Parser{
		grammar: v,
	}
}

func (p *Parser) Parse(t string) (n *ast.Node, err error) {
	trs := p.grammar.Transitions()
	var wg sync.WaitGroup
	for _, vtx := range trs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			switch vtx.Kind() {
			case vertex.DUMB:
			case vertex.TOKEN:
			case vertex.STRING:
			}
		}
	}
}
