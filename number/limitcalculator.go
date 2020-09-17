package number

import (
	"strconv"
	"fmt"
	"bufio"
	"os"
	"go/token"
	"go/parser"
	"go/ast"
)

func eval(exp ast.Expr) int {
	// Example of 'type switches'
	switch exp := exp.(type){
	case *ast.BinaryExpr:
		return evalBinaryExpr(exp)
	case *ast.BasicLit:
		switch exp.Kind{
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
	switch exp.Op{
	case token.ADD:
		return left + right
	case token.SUB:
		return left - right
	case token.MUL:
		return left*right
	case token.QUO:
		return left / right
	}
	return 0
}
// LimitCalculator : Ask the user to enter f(x) and the limit value, then return the value of the limit statement.
// Optional: Make the calculator capable of supporting infinite limits.
func LimitCalculator(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter f(x): ")
	scanned := scanner.Scan()
	if !scanned{
		os.Exit(0)
	}
	line := scanner.Text()
	exp, err := parser.ParseExpr(line)
	if err != nil{
		os.Exit(1)
	}
	fmt.Println(eval(exp))
}