// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch07/ex14"
)

type calculation struct {
	Expression string
	Result     string
}

// TODO 時間があれば、関数呼び出し時のエラーハンドリング (現状、サポート外関数呼び出ししたり、関数に対して少ない引数を渡すとパニックで落ちる)
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "parsing request: %v", err)
			return
		}

		expressionString := r.FormValue("expression")
		if expressionString == "" {
			calculator.Execute(w, &calculation{"", ""})
			return
		}

		expression, err := ex14.Parse(expressionString)
		if err != nil {
			calculator.Execute(w, &calculation{expressionString, fmt.Sprintf("error result: parsing expression: %v", err)})
			return
		}

		env := ex14.Env{"pi": math.Pi}
		result := expression.Eval(env)
		calculator.Execute(w, &calculation{expressionString, fmt.Sprintf("%g", result)})
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var calculator = template.Must(template.New("calculator").
	Parse(`
	<h1>calculator</h1>
	<form action="/" method="POST">
		<label>expression: <input type="text" name="expression" value={{.Expression}}></label> <input type="submit" value="calculate"><br>
		<label>result: {{.Result}}</label>
	</form>
	`))
