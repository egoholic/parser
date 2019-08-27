package source_test

import (
	"sync"

	. "github.com/egoholic/parser/source"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("source", func() {
	Describe("*Source", func() {
		Describe(".Stream()", func() {
			It("streams runes", func() {
				var (
					wg        sync.WaitGroup
					sig       = make(chan bool)
					runes     = []rune("abcdef")
					s         = New(runes, &wg, sig)
					ch        = s.Stream()
					collected = []rune{}
					finished  = false
				)
				wg.Add(1)
				go func(wg *sync.WaitGroup, ch <-chan rune, sig <-chan bool) {
					defer wg.Done()
					for {
						select {
						case r := <-ch:
							collected = append(collected, r)
						case <-sig:
							finished = true
							return
						}
					}
				}(&wg, ch, sig)
				wg.Wait()
				Expect(collected).To(Equal([]rune("abcdef")))
				Expect(finished).To(BeTrue())
			})
		})
	})
})
