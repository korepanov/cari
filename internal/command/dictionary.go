package command

import (
	"strconv"

	"github.com/korepanov/cari/internal/lexemes"
)

type dictionary struct {
	dictionaryValues map[lexemes.LexemeType][]lexemes.Lexeme
	stopList         []lexemes.LexemeType // to stop lexical analysis
}

func Dictionary() *dictionary {
	var d dictionary
	d.initValues()
	d.initStopList()
	return &d
}

func (d *dictionary) initValues() {
	d.dictionaryValues = make(map[lexemes.LexemeType][]lexemes.Lexeme)
	d.dictionaryValues[lexemes.Delimiter] = []lexemes.Lexeme{"(", ")"}
	d.dictionaryValues[lexemes.Operator] = []lexemes.Lexeme{"+", "-", "*", "/"}
}

func (d *dictionary) initStopList() {
	d.stopList = append(d.stopList, lexemes.Delimiter)
	d.stopList = append(d.stopList, lexemes.Operator)
}

// to stop lexical analysis and return the previous token?
func (d *dictionary) IsStop(t lexemes.LexemeType) bool {
	for _, val := range d.stopList {
		if val == t {
			return true
		}
	}
	return false
}

func (d *dictionary) Find(lex lexemes.Lexeme) lexemes.LexemeType {
	lexT := find(d.dictionaryValues, lex)
	if lexT != 0 {
		return lexT
	}

	return isNumber(lex)
}

/*
finds value v of type T in the map[T2][]T
if value found return key of type T2
if value was not found return zero value of type T2
*/
func find[T comparable, T2 comparable](m map[T2][]T, v T) T2 {
	for key, vals := range m {
		for _, val := range vals {
			if val == v {
				return key
			}
		}
	}

	var res T2

	return res
}

/*
can the lexeme to be an number?
return NumberLexeme if lexeme can be a number
else - 0
*/
func isNumber(lex lexemes.Lexeme) lexemes.LexemeType {
	_, err := strconv.ParseFloat(string(lex), 64)
	if err == nil {
		return lexemes.NumberLexeme
	}
	return 0
}
