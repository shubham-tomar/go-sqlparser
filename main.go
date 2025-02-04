package main

import (
	"fmt"
	"go-sql-parser/lexer"
)

func main() {
	sql := "CREATE TABLE users (id INT, name STRING);"
	l := lexer.NewLexer(sql)

	for tok := l.NextToken(); tok.Type != lexer.EOF; tok = l.NextToken() {
		fmt.Printf("Token: Type=%s, Literal=%s\n", tok.Type, tok.Literal)
	}
}
