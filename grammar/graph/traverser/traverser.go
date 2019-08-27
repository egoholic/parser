package traverser

import (
	"go/ast"
	"sync"

	"github.com/egoholic/parser/grammar/graph/vertex"
	"github.com/egoholic/parser/source"
)

type Traverser struct {
	vertex *vertex.Vertex
	source *source.Source
	ast    *ast.Node
}

func New(v *vertex.Vertex, s *source.Source, ast *ast.Node) *Traverser {
	return &Traverser{
		vertex: v,
		source: s,
		ast:    ast,
	}
}
func (t *Traverser) Step() {
}
func (t *Traverser) makeContext() {
	var wg sync.WaitGroup
}
