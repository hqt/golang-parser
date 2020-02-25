
antlr = java -jar etc/antlr-4.8-complete.jar
grunt = ${antlr} org.antlr.v4.gui.TestRig

# Follow: https://github.com/antlr/grammars-v4/tree/master/golang
gen-parser:
	rm -rf parser;
	${antlr} GoLexer.g4 GoParser.g4 -Dlanguage=Go -visitor -o parser;
	cp etc/GoParserBase.go parser/;
	sed -i '' -e "s/noTerminatorAfterParams/p.noTerminatorAfterParams/g" \
			  -e "s/noTerminatorBetween/p.noTerminatorBetween/g" \
			  -e "s/lineTerminatorAhead/p.lineTerminatorAhead/g" \
			  -e "s/checkPreviousTokenText/p.checkPreviousTokenText/g" \
			  parser/go_parser.go;
.PHONY: gen-parser

test-grunt:
	java -jar etc/antlr-4.8-complete.jar org.antlr.v4.gui.TestRig
.PHONY: test-grunt