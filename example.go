package example

import (
	"github.com/Applifier/golang-backend-assignment/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func RunGolangParser() parser.GoParserListener {
	is, err := antlr.NewFileStream("etc/demo.go")
	if err != nil {
		panic(err)
	}

	lexer := parser.NewGoLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewGoParser(stream)

	listener := NewGoParserListenerImpl()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.SourceFile())
	return listener
}
