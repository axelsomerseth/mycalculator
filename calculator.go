package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calculator struct {
	input     string
	operation string
}

func (c calculator) operate() int {
	cleanInput := strings.Split(c.input, c.operation)
	number1 := parse(cleanInput[0])
	number2 := parse(cleanInput[1])
	switch c.operation {
	case "+":
		return number1 + number2
	case "-":
		return number1 - number2
	case "*":
		return number1 * number2
	case "/":
		return number1 / number2
	default:
		fmt.Println()
		fmt.Print(c.operation, " Operation not supported.\n")
		return 0
	}
}

func parse(input string) int {
	operator, _ := strconv.Atoi(input)
	return operator
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operation := scanner.Text()
	return operation
}
