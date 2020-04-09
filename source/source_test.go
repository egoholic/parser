package source_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/stoa-bd/parser/source"
)

var _ = Describe("source", func() {
	Describe("New()", func() {
		It("returns *Source", func() {
			Expect(New("source")).To(BeAssignableToTypeOf(&Source{}))
		})
	})

	Describe("*Source", func() {
		var source

		BeforeEach(func(){
			source = New("source")
		})

		Describe(".Pop()", func() {
			Context("when it has more runes", func() {
				It("returns next rune", func(){
					cur, r, ok := source.Pop()
          Expect(cur).To(Equal(0, 's', true))
				})
			})
		})
		Describe(".Commit()", func() {

		})
		Describe(".Rollback()", func() {

		})
	})
})
