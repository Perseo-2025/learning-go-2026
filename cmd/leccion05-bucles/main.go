package main

import (
	"fmt"
)

func main(){
	
	//1. for clasico
	// init; condicion; post - EXACTAMENTE la misma estructura que conoces,
	// pero SIN parentesis () y con llaves {} obligatorias.
	fmt.Println("--For clásicos--")

	for i := 0; i < 5; i++ {
		fmt.Println("Iteración ", i)
	}

	// For con While
	// Si omites init y post, y dejas solo la condicion, se comporta como while.
	fmt.Println("---For como while---")
	
	contador := 0
	// en reemplaz de while es for
	for contador < 3 {
		fmt.Println("Contador", contador)
		contador++ // importante para continuar con el bucle
	}

	// for infinito
	// se usa MUCHISIMO en servidores: "mientras el porgrama viva, escucha peticiones".
	// Necesitas un 'break' interno para no colgar o saturar el programa
	fmt.Println("---For infinito con breck---")
	intentos := 0
	for {
		intentos++
		fmt.Println("Intentos número: ",intentos)
		if intentos == 3 {
			break // si son 3 intentos entonces tomate un descanso
		}
	}

	// For Range sobre un SLICE
	// 'range' te da DOS valores: el índice y el valor de esa posición
	fmt.Println("---For range sobre slice ----")
	frutas := []string{"manzana", "pera", "uva"} // [{manzana},{pera}]
	for indice, valor := range frutas {
		fmt.Printf("Posición %d: %s\n", indice, valor)
	}

	// si no me interesa obtener el índice agreago "_"
	fmt.Println("-- For range ignorando el índice")
	for _, valor := range frutas{
		fmt.Println("valor: ",valor)
	}

	//ahora si quiero el índice y no el valor simplemente se omite la segunda variable
	for indice := range frutas {
		fmt.Println("indice: ", indice)
	}


	// For Range sobre Map
	// Aqui el 'range' te da la clave y valor . OJO: el orden de un mapa NO está
	// garantizado en GO - cada ejecución puede variar el orden mostrado.
	fmt.Println("--For Range sobre map")
	inventario := map[string]int{"manzanas":10,"peras":5}
	for clave, valor := range inventario {
		fmt.Printf("%s: %d unidades\n", clave, valor)
	}

	// For con "Continue" // literalmenta asi se llama la expreción
	// salta a la siguiente iteración, igual que en JS
	fmt.Println("--- Continue (solo numeros pares)")
	for i := 1; i<=10; i++ {
		if i % 2 != 0 {
			continue // si es impar, saltamos esta iteración
		}
		fmt.Println("Par: ", i)
	}


}