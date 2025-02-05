package main

import (
	"fmt"
	"go-sqlparser/lexer"
)

func main() {
	testSQL := []string{
		"CREATE TABLE users (id INT, name STRING);",
	}

	for _, sql := range testSQL {
		fmt.Println("\nTesting SQL:", sql)
		l := lexer.NewLexer(sql)

		for {
			tok, err := l.NextToken()
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			fmt.Printf("Token: Type=%s, Literal=%s\n", tok.Type, tok.Literal)
			if tok.Type == lexer.EOF {
				break
			}
		}
	}
}
