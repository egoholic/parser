package source

import "sync"

type Source struct {
	runes []rune
	wg    *sync.WaitGroup
	sig   chan<- bool
}

func New(runes []rune, wg *sync.WaitGroup, sig chan<- bool) *Source {
	return &Source{
		runes: runes,
		wg:    wg,
		sig:   sig,
	}
}

func (s *Source) Stream() <-chan rune {
	ch := make(chan rune)
	s.wg.Add(1)
	go s.stream(ch)
	return ch
}

func (s *Source) stream(ch chan rune) {
	defer s.wg.Done()
	defer s.close()
	for _, r := range s.runes {
		ch <- r
	}
}

func (s *Source) close() {
	s.sig <- true
}
