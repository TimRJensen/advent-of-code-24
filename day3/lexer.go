package main

const (
	FUNC = iota
	BAD_FUNC
	EOF
)

const (
	Do   string = "do"
	Dont string = "don't"
)

func isValidNum(b byte) bool {
	// codepoints < '0' will underflow and be greater than '9'.
	return b-'0' < '9'
}

func isValidMul(b []byte) bool {
	return len(b) > 3 && b[len(b)-1] == '(' && b[len(b)-2] == 'l' && b[len(b)-3] == 'u' && b[len(b)-4] == 'm'
}

func isValidDo(b []byte) bool {
	return len(b) > 2 && b[len(b)-1] == '(' && b[len(b)-2] == 'o' && b[len(b)-3] == 'd'
}

func isValidDont(b []byte) bool {
	return len(b) > 5 && b[len(b)-1] == '(' && b[len(b)-2] == 't' && b[len(b)-3] == '\'' && b[len(b)-4] == 'n' && b[len(b)-5] == 'o' && b[len(b)-6] == 'd'
}

type token struct {
	typ byte
	val []byte
}

type lexer struct {
	input []byte
	pos   int
}

func (l *lexer) consumeMul() (int, *token) {
	flag := true
	i := l.pos + 1
	for ; flag; i++ {
		switch {
		case isValidNum(l.input[i]), l.input[i] == ',':
			flag = i+1 < len(l.input)
		case l.input[i] == ')':
			return i, &token{FUNC, l.input[l.pos+1 : i]}
		default:
			return i, &token{BAD_FUNC, l.input[l.pos+1 : i]}
		}
	}
	return i, &token{BAD_FUNC, l.input[l.pos:i]}
}

func (l *lexer) consumeDo() (int, *token) {
	return l.pos + 2, &token{FUNC, l.input[l.pos-2 : l.pos]}
}

func (l *lexer) consumeDont() (int, *token) {
	return l.pos + 2, &token{FUNC, l.input[l.pos-5 : l.pos]}
}

func (l *lexer) consumeFunc() (int, *token) {
	switch {
	case isValidMul(l.input[:l.pos+1]):
		return l.consumeMul()
	case isValidDo(l.input[:l.pos+1]):
		return l.consumeDo()
	case isValidDont(l.input[:l.pos+1]):
		return l.consumeDont()
	default:
		return l.pos, &token{BAD_FUNC, l.input[:l.pos]}
	}
}

func (l *lexer) next() *token {
	flag := l.pos < len(l.input)
	for flag {
		switch l.input[l.pos] {
		case '(':
			pos, tok := l.consumeFunc()
			l.pos = pos + 1
			return tok
		default:
			l.pos++
			flag = l.pos < len(l.input)
		}
	}

	return &token{EOF, nil}
}

func newLexer(input []byte) *lexer {
	return &lexer{input: input}
}
