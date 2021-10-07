package reader

import (
	"bufio"
	"bytes"
	"io"
)

// A Reader consumes tokens from an input stream. Use the Next method to parse
// tokens sequentially from the input until a non-nil error is reported.
type Reader struct {
	input    *bufio.Reader // the unconsumed input
	text     *bytes.Buffer // the text of the current token
	err      error         // the last non-nil error reported
	pos, end int           // the span of the current token
}

// Type denotes the lexical type of a token.
type Type int

// The legal values for a token type.
const (
	Invalid Type = iota // an invalid token
	Word                // an unquoted word (no whitespace): foo
	Name                // an unquoted symbol: :bar
	Quoted              // a quoted word (whitespace OK): "foo bar"
	LSquare             // a left square bracket: [
	RSquare             // a right square bracket: ]
	LParen              // a left parenthesis: (
	RParen              // a right parenthesis: )
	LCurly              // a left bracket: {
	RCurly              // a right bracket: }

	numTypes
)

func New(r io.Reader) *Reader {
	panic("not implemented")
}

// Next advances r to the next token in the input, and returns nil if a valid
// token is available. If no further tokens are available it returns io.EOF.
// Otherwise, the error reports what went wrong.
func (r *Reader) Next() error {

}

/*
Lexical grammar:

word            -- no whitespace
"quoted word"   -- multi-line OK
:name
123    -- via word
45.67  -- via word

Self-quoting:

[ ]  ( )  { }

Syntax:

elt  = word | seq
elts = elt elts...
seq  = sseq | rseq | cseq
sseq = [ elts? ]
rseq = ( elts? )
cseq = { elts? }

:name { (a _ _ b) @b @a } def
>> 1 2 3 4 name --> 4 1

*/
