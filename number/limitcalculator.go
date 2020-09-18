package number

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"os"
	"strconv"
	"strings"
)

func eval(exp ast.Expr) int {
	// Example of 'type switches'
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return evalBinaryExpr(exp)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.INT:
			i, _ := strconv.Atoi(exp.Value)
			return i
		}
	}
	return 0
}

func evalBinaryExpr(exp *ast.BinaryExpr) int {
	left := eval(exp.X)
	right := eval(exp.Y)
	switch exp.Op {
	case token.ADD:
		return left + right
	case token.SUB:
		return left - right
	case token.MUL:
		return left * right
	case token.QUO:
		return left / right
	}
	return 0
}

// LimitCalculator : Ask the user to enter f(x) and the limit value, then return the value of the limit statement.
// Optional: Make the calculator capable of supporting infinite limits.
func LimitCalculator() {
	var limit string
	var inc, dec, con int
	fmt.Print("Enter limit: ")
	fmt.Scan(&limit)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter f(x): ")
	scanned := scanner.Scan()
	if !scanned {
		os.Exit(0)
	}
	line := scanner.Text()
	if limit == "infinite" {
		oldresult := 0
		result := 0.0
		for i := 1; i < 10; i++ {
			limit = fmt.Sprint(i)
			exp, err := parser.ParseExpr(strings.ReplaceAll(line, "x", limit))
			if err != nil {
				os.Exit(1)
			}
			currentresult := eval(exp)
			if currentresult > oldresult {
				inc++
			} else if currentresult < oldresult {
				dec++
			} else {
				con++
			}
			oldresult = currentresult
		}
		if inc > dec {
			result = math.Inf(inc)
		} else if dec > inc {
			result = math.Inf(-dec)
		} else {
			result = float64(oldresult)
		}
		fmt.Println(result)
	} else {
		exp, err := parser.ParseExpr(strings.ReplaceAll(line, "x", limit))
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(eval(exp))

	}
}
