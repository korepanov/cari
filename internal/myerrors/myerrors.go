package myerrors

import "errors"

var (
	ErrNoToken     = errors.New("no token found")
	ErrLexAnalysis = errors.New("lexical analysis failed")
	ErrRead        = errors.New("reading program failed")
	ErrCompile     = errors.New("compiling failed")
	ErrFlagParse   = errors.New("could not parse flags")
	ErrHelp        = errors.New("help")
	ErrNoFactor    = errors.New("no factor")
	ErrNoTerm      = errors.New("no term")
	ErrNoExpr      = errors.New("no expression")
	ErrParse       = errors.New("parsing failed")
)
