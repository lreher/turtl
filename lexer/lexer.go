package lexer

import (
  "turtl/token"
  "fmt"
)

type Lexer struct {
  input string
  position int // current position in input (by char)
  readPosition int // one more than reading position, detect future chars
  ch byte // current char being examined
}

func New(input string) *Lexer {
  lex := &Lexer{input: input}
  lex.readChar()
  return lex
}

func (lex *Lexer) NextToken() token.Token {
  var tok token.Token

  lex.skipWhitespace()

  switch lex.ch {
  case '=':
    if lex.peekChar() == '=' {
      ch := lex.ch
      lex.readChar()
      literal := string(ch) + string(lex.ch)
      tok = token.Token{ Type: token.EQ, Literal: literal }
    } else {
      tok = newToken(token.ASSIGN, lex.ch)
    }
  case ';':
    tok = newToken(token.SEMICOLON, lex.ch)
  case '(':
    tok = newToken(token.LPAREN, lex.ch)
  case ')':
    tok = newToken(token.RPAREN, lex.ch)
  case ',':
    tok = newToken(token.COMMA, lex.ch)
  case '+':
    tok = newToken(token.PLUS, lex.ch)
  case '-':
    tok = newToken(token.MINUS, lex.ch)
  case '{':
    tok = newToken(token.LBRACE, lex.ch)
  case '}':
    tok = newToken(token.RBRACE, lex.ch)
  case '!':
    if lex.peekChar() == '=' {
      ch := lex.ch
      lex.readChar()
      literal := string(ch) + string(lex.ch)
      tok = token.Token{ Type: token.NOT_EQ, Literal: literal }
    } else {
      tok = newToken(token.BANG, lex.ch)
    }
  case '/':
    tok = newToken(token.SLASH, lex.ch)
  case '*':
    tok = newToken(token.ASTERISK, lex.ch)
  case '<':
    tok = newToken(token.LT, lex.ch)
  case '>':
    tok = newToken(token.GT, lex.ch)
  case 0:
    tok.Literal = ""
    tok.Type = token.EOF
  default:
    if isLetter(lex.ch) {
      tok.Literal = lex.readIdentifier()
      tok.Type = token.LookupIdent(tok.Literal)
      return tok
    } else if isDigit(lex.ch) {
      tok.Type = token.INT
      tok.Literal = lex.readNumber()
      return tok
    } else {
      tok = newToken(token.ILLEGAL, lex.ch)
    }
  }
  lex.readChar()
  return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
  return token.Token { Type: tokenType, Literal: string(ch) }
}

func (lex *Lexer) readChar() {
  if lex.readPosition >= len(lex.input) {
    lex.ch = 0
  } else {
    lex.ch = lex.input[lex.readPosition]
  }

  lex.position = lex.readPosition
  lex.readPosition += 1
}

func (lex *Lexer) peekChar() byte {
  if lex.readPosition >= len(lex.input) {
    return 0
  } else {
    return lex.input[lex.readPosition]
  }
}

func (lex *Lexer) readIdentifier() string {
  position := lex.position

  for isLetter(lex.ch) {
    lex.readChar()
  }

  return lex.input[position:lex.position]
}

func isLetter(ch byte) bool {
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (lex *Lexer) readNumber() string {
  position := lex.position

  for isDigit(lex.ch) {
    lex.readChar()
  }

  return lex.input[position:lex.position]
}

func isDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
}

func (lex *Lexer) skipWhitespace() {
  for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
    lex.readChar()
  }
}
