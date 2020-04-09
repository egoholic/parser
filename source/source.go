package source

import "fmt"

type Source struct {
	origin      []rune
	currentCur  int
	commitedCur int
}

func New(origin string) *Source {
	return &Source{
		origin:      []rune(origin),
		currentCur:  0,
		commitedCur: 0,
	}
}

func (s *Source) Pop() (cur int, r rune, ok bool) {
	cur = s.currentCur
	ok, r = s.origin[cur]
	return
}
func (s *Source) Commit(n int) (err error) {
	if n < s.commitedCur {
		return fmt.Errorf("new commited cursor `%d` should be bigger than previous `%d`", n, s.commitedCur)
	}
	s.commitedCur = n
	return
}
func (s *Source) Rollback() (n int) {
	s.currentCur = s.commitedCur
}
