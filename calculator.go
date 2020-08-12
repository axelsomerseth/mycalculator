package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calculadora struct {
	entrada   string
	operacion string
}

func (c calculadora) operar() int {
	// la variable c hace referencia a una instancia del tipo calculadora. Operar es un método de calculadora
	entradaLimpia := strings.Split(c.entrada, c.operacion)
	numero1 := parsear(entradaLimpia[0])
	numero2 := parsear(entradaLimpia[1])
	switch c.operacion {
	case "+":
		return numero1 + numero2
	case "-":
		return numero1 - numero2
	case "*":
		return numero1 * numero2
	case "/":
		return numero1 / numero2
	default:
		fmt.Println()
		fmt.Print(c.operacion, " La operación no está soportada. ")
		return 0
	}
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	return operacion
}
