package ast_test

import (
	. "github.com/egoholic/parser/ast"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ast", func() {
	Describe("*Node", func() {
		Describe(".AddChild()", func() {
			It("adds child node", func() {
				n1 := New("test", nil)
				n2 := New("test-child", nil)
				Expect(n1.Children).To(BeEmpty())
				n1.AddChild(n2)
				Expect(n1.Children).To(Equal([]*Node{n2}))
			})
		})
	})
})
