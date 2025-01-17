package main

import (
	"flag"
	"fmt"
	"gocomp/compiler/lexer"
	"io/ioutil"
	"log"
	"os"
)

var (
	filepath string
)

func init() {
	flag.StringVar(&filepath, "f", "", "Path to needed file")
}

func readFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func main() {
	flag.Parse()

	data, err := readFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	lex := lexer.NewLexer(data)

	for lex.TokenIdentifier != lexer.TokenEOF {
		fmt.Println(lex.Token, "\t", lex.TokenIdentifier)
		lex.NextToken()
	}

}

