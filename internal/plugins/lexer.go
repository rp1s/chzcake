package plugins

import 	"github.com/rp1s/chzcake/internal/parser/token/token"


type RulePlugin struct

func NextToken(rn rune, lexer *Lexer) (token.Token, err error)
