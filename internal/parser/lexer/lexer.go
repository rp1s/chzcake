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

	CustomRuleLexer []int
}

func New(input []byte, filename string, CustomRuleLexer []int) *Lexer {
	if input == nil {
		input = input[:0]
	}
	l := &Lexer{
		Input:    input,
		Filename: filename,

		line: 1,
		col:  0,

		tokLine: 1,
	}
	l.advance()
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

func (l *Lexer) peek() rune {
	if l.curPos >= uint64(len(l.Input)) {
		return 0
	}
	if b := l.Input[l.curPos]; b < utf8.RuneSelf {
		return rune(b)
	}
	r, _ := utf8.DecodeRune(l.Input[l.curPos:])
	return r
}

func (l *Lexer) tok(kind types.TokenKind) token.Token {
	return token.Token{
		Kind: kind,
		Pos: position.Position{
			FileName: l.Filename,
			Line:     uint64(l.tokLine),
			Column:   uint64(l.tokCol),
			Offset:   uint64(l.tokPos),
		},
	}
}

func (l *Lexer) NextToken() token.Token {
	l.tokPos = l.pos
	l.tokLine = l.line
	l.tokCol = l.col

	if isSpace(l.rn) {
		for isSpace(l.rn) {
			l.advance()
		}
		return l.tok(types.TokenKindWHITESPACE)
	}

	switch l.rn {

	}

	return l.tok(types.TokenKindINVALID)
}

func isSpace(r rune) bool {
	switch r {
	case ' ', '\t', '\n', '\r', '\v', '\f':
		return true
	}
	return r > 0x7F && isSpaceUnicode(r)
}

func isSpaceUnicode(r rune) bool {
	switch r {
	case 0x00A0,
		0x1680,
		0x2000, 0x2001, 0x2002, 0x2003,
		0x2004, 0x2005, 0x2006, 0x2007,
		0x2008, 0x2009, 0x200A,
		0x2028, 0x2029,
		0x202F, 0x205F,
		0x3000,
		0xFEFF:
		return true
	}
	return false
}
