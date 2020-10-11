.DEFAULT_GOAL: runb

runb: 
	go build -o bin/main main.go && ./bin/main -f $(f)


build:
	go build -o bin/main main.go

run:
	./bin/main -f $(f)

test:
	go test compiler/lexer/lexer_test.go \
			compiler/lexer/lexer.go \
			compiler/lexer/tokens.go \
			compiler/lexer/utils.go \
			-run ^TestLexer$

	go test compiler/stack/stack_test.go \
			compiler/stack/stack.go \
			compiler/stack/errors.go \
			-run ^TestStack$
	
	