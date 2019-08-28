package context

import (
	"sync"

	"github.com/egoholic/parser/grammar/graph/vertex"
)

type Context struct {
	wg       *sync.WaitGroup
	vertices []*vertex.Vertex
}

func New(wg *sync.WaitGroup, vertices []*vertex.Vertex) *Context {
	return &Context{
		wg:       wg,
		vertices: vertices,
	}
}

func (c *Context) Exec() error {
	c.wg.Add(1)
	defer c.wg.Done()
	for _, v := range c.vertices {
		c.wg.Add(1)
		go func(wg *sync.WaitGroup, v *vertex.Vertex) {
			defer wg.Done()
		}(c.wg, v)
	}
}

func (c *Context) parseWith(wg *sync.WaitGroup, v *vertex.Vertex) {
	defer wg.Done()

}

func (c *Context) switchToParseEndingMode(endingVertices []*vertex.Vertex) {
	for _, v := range endingVertices {
		c.wg.Add(1)
		go func(wg *sync.WaitGroup, v *vertex.Vertex) {
			defer wg.Done()
		}(c.wg, v)
	}
}
