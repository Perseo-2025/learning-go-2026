package main

import (
	"fmt"
)

// Función que RECIBE otra función como parámetro
// 'operacion' es un parámetro cuyo TIPO es "una función" que recibe dos int y retorna un int ""
// osea auna función dentro de otra función

func operar(a, b int, operacion func(int, int) int) int {
	return operacion(a, b)
}

// funcion que RETORNA otra función (closure)
// El tipo de retorno es 'func() int': "una función sin parámetros que retorna un int".
func crearContador() func() int {
	// ' contador' vive en el scope de crearContador(), pero la función
	//anonima de abajo la "atrapa" y sigue usándola aunque crearContador()
	// ya haya termiando de ejecutarse. Esto es el closure.
	contador := 0

	return func() int {
		contador++
		return contador
	}
}

func main1() {

	// 1. función anonima asiganada a una varibale
	saludar := func(nombre string) string { // no se sabe el nomrbe de esa funcion pero se asigana a una variable
		return "Hola, " + nombre
	}
	fmt.Println(saludar("Javier"))

	// 2. funcion anónima ejecutada INMEDIATAMENTE (COMO UN IIFE en JS)
	resultado := func(a, b int) int {
		return a + b
	}(5, 3) //los parámetros finales la ejecutan al instante
	fmt.Println("IIFE resultados", resultado)

	// 3. Pasando una función como argumento
	// Definido dos funciones nombre que cumplen la firma func(int, int) int
	sumar := func(a, b int) int { return a + b }
	multiplicar := func(a, b int) int { return a * b }

	fmt.Println("Operar sumando: ", operar(4, 5, sumar))
	fmt.Println("Operar multiplicando: ", operar(4, 5, multiplicar))

	// Tambien peudo pasar la función anónima DIRECTAMENTE, sin guardarla antes:
	fmt.Println("Operar restando: ", operar(10, 3, func(a, b int) int {
		return a - b
	}))

	// 4. Closure en accion
	contadorA := crearContador()
	contadorB := crearContador()

	fmt.Println("Contador A:", contadorA()) // 1
	fmt.Println("Contador A:", contadorA()) // 2
	fmt.Println("Contador A:", contadorA()) // 3
	fmt.Println("Contador B:", contadorB()) // 1 - totalmente independiente de A
}
