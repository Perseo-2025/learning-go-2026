package main

import (
	"errors"
	"fmt"
)

/*
SINTAXIS
	func nombreFuncion(parametro1 tipo1, parametro2 tipo2){
		//cuerpo
		return valor
	}
*/

// función básica: un parámetro, un retorno
func saludar(nombre string) string {
	return "Hola " + nombre
}

// Parametros del mismo tipo agrupados
// En vez de (a int, b int) se puede escribir (a, b int) si son consecutivos iguales.
func sumar(a, b int) int {
	return a + b
}

// funcion que retorna multiples valores
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		// errors.New crea un error simple con un mensaje, como no se puede divir entre 0, asi que retornamos un valor "vacio"
		return 0, errors.New("No se puede dividir entre 0")
	}
	return a / b, nil
}

// Funcion que retorna nombre (named returns)
// Puedes "pre-declarar" el nombre de lo que vas a retornar en la firma.
// Usalo con moderación: es útil en funciones cortas, pero puede volverse confuso en funciones largas
func dividirConNombre(a, b float64) (resultado float64, err error) {
	if b == 0 {
		err = errors.New("No se puede dividir entre cero") //<- imprimer el resultado
		return                                             // "naked return": retorna automaticamente  (resultado,err)
	}
	resultado = a / b
	return // retorna (resultado, nil) -- 'err' nunca se toó, sigue siendo nil no null bueno es lo mismo
}

/// RETOO

// 1. Retorna true si el numero es par
func esPar(numero int) bool {
	return numero%2 == 0
}

// 2. Retrona la resta de dos enteros
func restar(a, b int) int {
	return a - b
}

func validarEdad(edad int) (bool, error) {
	if edad >= 0 && edad <= 120 {
		return true, nil
	}
	return false, errors.New("Edad fuera del rango")
}

// función principal - main
func main() {

	//a partir de ahora llamamos a todas las funciones que hemos creado previamente
	// Uso de la función básica
	mensaje := saludar("Rey")
	fmt.Println(mensaje)

	//Uso de la función con parámetros agrupados
	total := sumar(5, 3)
	fmt.Println("Suma", total)

	// Capturando un retorno multiple
	//Nota la similitud con el "if con inicializador" que ya conoces
	resultado, err := dividir(10, 2)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Resultado: ", resultado)
	}

	// Ahora provovamos el caso de error apropósito
	// hacemos apropósito

	resultado2, err2 := dividir(10, 0)
	if err2 != nil {
		fmt.Println("Error", err2)
	} else {
		fmt.Println("Resultado: ", resultado2)
	}

	// El error siempre debe mostrarse tanto en produccion como en desarrollo, es una buena practicar marcar donde se encuentra el error
	resultadoSolo, _ := dividir(20, 4)
	fmt.Println("Solo el resultado sin mostrar el error ni nada: ", resultadoSolo)

	// Uso de la versión con nombre de retorno
	r, e := dividirConNombre(8, 2)
	fmt.Println("Con nombre: ", r, e)

	// === RETO
	fmt.Println("RETO 06")
	// Probar esPar
	numero := 8
	fmt.Printf("¿%d es par?: %t\n", numero, esPar(numero))

	// Probando validarEdad con un valor VÁLIDO
	valida, err := validarEdad(25)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Edad válida: ", valida)
	}

	// Probando validarEdad con un valir INVÁLIDO
	valida2, err2 := validarEdad(150)
	if err2 != nil {
		fmt.Println("Error Impreso: ", err2)
	} else {
		fmt.Println("Edad válida: ", valida2)
	}

	//notita
	/*
	
	:= esto crea
	= esto reasiga
	*/
}
