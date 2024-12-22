package main

type grid struct {
	buff []byte
	n    int
}

func (g grid) left(i int) (bool, int) {
	if i-1 < 0 || (i-1)/g.n < i/g.n {
		return false, i
	}
	return true, i - 1
}

func (g grid) right(i int) (bool, int) {
	if i+1 >= len(g.buff) || (i+1)/g.n > i/g.n {
		return false, i
	}
	return true, i + 1
}

func (g grid) up(i int) (bool, int) {
	if i/g.n == 0 {
		return false, i
	}
	return true, i - g.n
}

func (g grid) down(i int) (bool, int) {
	if i/g.n == g.n-1 {
		return false, i
	}
	return true, i + g.n
}

func (g grid) upLeft(i int) (bool, int) {
	if ok, i := g.up(i); ok {
		return g.left(i)
	}
	return false, i
}

func (g grid) upRight(i int) (bool, int) {
	if ok, i := g.up(i); ok {
		return g.right(i)
	}
	return false, i
}

func (g grid) downLeft(i int) (bool, int) {
	if ok, i := g.down(i); ok {
		return g.left(i)
	}
	return false, i
}

func (g grid) downRight(i int) (bool, int) {
	if ok, i := g.down(i); ok {
		return g.right(i)
	}
	return false, i
}

func (g grid) isValidSequence(i int, op func(int) (bool, int), txt string, pos int) bool {
	if g.buff[i] != txt[pos] {
		return false
	}

	ok, i := op(i)
	if !ok || g.buff[i] != txt[pos+1] {
		return false
	}
	if g.buff[i] == txt[len(txt)-1] {
		return true
	}
	return g.isValidSequence(i, op, txt, pos+1)
}
