package main

import (
	"unicode/utf8"

	"github.com/rp1s/chzcake/internal/parser/token/position"
	"github.com/rp1s/chzcake/internal/parser/token/token"
	"github.com/rp1s/chzcake/internal/parser/token/types"
)

type Lexer struct {
	Input    []byte
	Filename string

	line uint64
	col  uint64

	pos    uint64
	curPos uint64
	rn     rune
	rnSize uint64

	tokPos  uint64
	tokLine uint64
	tokCol  uint64
}

func New(input []byte, filename string) *Lexer {
	l := &Lexer{
		Input:    input,
		Filename: filename,

		line: 1,
		col:  0,

		tokLine: 1,
	}
	return l
}

func (l *Lexer) advance() *Lexer {
	if l.curPos >= uint64(len(l.Input)) {
		l.pos = l.curPos
		l.rn = 0
		l.rnSize = 0
		return l
	}

	l.pos = l.curPos

	if b := l.Input[l.curPos]; b < utf8.RuneSelf {
		l.rn = rune(b)
		l.rnSize = 1
	} else {
		r, size := utf8.DecodeRune(l.Input[l.curPos:])
		l.rn = r
		l.rnSize = uint64(size)
	}

	l.curPos += l.rnSize

	if l.rn == '\n' {
		l.line++
		l.col = 0
	} else {
		l.col++
	}

	return l
}

func (self *Lexer) peek() rune {
	if self.curPos >= uint64(len(self.Input)) {
		return 0
	}
	if b := self.Input[self.curPos]; b < utf8.RuneSelf {
		return rune(b)
	}
	r, _ := utf8.DecodeRune(self.Input[self.curPos:])
	return r
}

func (self *Lexer) tok(kind types.TokenKind) token.Token {
	return token.Token{
		Kind: kind,
		Pos: position.Position{
			FileName: self.Filename,
			Line:     uint64(self.tokLine),
			Column:   uint64(self.tokCol),
			Offset:   uint64(self.tokPos),
		},
	}
}
