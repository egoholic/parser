package vertex_test

import (
	"fmt"

	. "github.com/egoholic/parser/grammar/graph/vertex"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("vertex", func() {
	Describe("String()", func() {
		It("returns Vertex", func() {
			Expect(String()).NotTo(BeNil())
		})
	})

	Describe("Token()", func() {
		It("returns Vertex", func() {
			Expect(Token("-- UP")).NotTo(BeNil())
		})
	})

	Describe("vertex", func() {
		Describe(".Name()", func() {
			It("returns name", func() {
				var vertex *Vertex
				vertex = Token("-- UP")
				Expect(vertex.Name()).To(Equal("<TOKEN>:`-- UP`"))
				vertex = String()
				Expect(vertex.Name()).To(Equal("<STRING>:\"...\""))
				vertex = Dumb()
				Expect(vertex.Name()).To(Equal("DUMB"))
			})
		})

		Describe(".Transitions()", func() {
			It("returns transitions", func() {
				var v1, v2 *Vertex
				v1 = Token("-- UP")
				v2 = String()
				Expect(v1.Transitions()).To(BeEmpty())
				v1.TransitsTo(v2)
				Expect(v1.Transitions()).To(Equal([]*Vertex{v2}))
				Expect(v2.Transitions()).To(BeEmpty())
				v2.TransitsTo(v1)
				Expect(v2.Transitions()).To(Equal([]*Vertex{v1}))
			})
		})

		Describe(".TransitsTo()", func() {
			Context("when transition to the given vertex exists", func() {
				It("panics", func() {
					v := Token("START")
					tv := String()
					defer func() {
						r := recover()
						Expect(r).NotTo(BeNil())
						Expect(r.(error).Error()).To(Equal(fmt.Sprintf("vertex `%s` already has transition to `%s`", v.Name(), tv.Name())))
					}()
					v.TransitsTo(tv)
					v.TransitsTo(tv)
				})
			})
			Context("when transition to the given vertex does not exist", func() {
				It("adds transition to the given vertexes", func() {
					v := Token("START")
					Expect(v.Transitions()).To(BeEmpty())
					tv1 := String()
					tv2 := Token("END")
					v.TransitsTo(tv1, tv2)
					Expect(v.Transitions()).To(Equal([]*Vertex{tv1, tv2}))
				})
			})
		})
	})
})
