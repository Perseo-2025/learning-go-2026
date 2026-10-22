package main

import "fmt"


/*
Suma Recursiva de 1 hasta n
*/

func sumarHasta(n int) int {

	if n <= 0 {
		return 0
	}

	return n + sumarHasta(n)

}

/*
	Potencia recursiva: base elevado al exponente
*/

func potencia(base, exponente int) int {

	if exponente == 0 {
		return 1
	}
	return  base * potencia(base, exponente-1) // el menos hace que el numero disminuye
}

func main(){

	// probando la funcion sumar hasta

	fmt.Println("Suma hasta 5: ", sumarHasta(5))
	fmt.Println("Suma hasta 10: ", sumarHasta(10))

	// probando la funcion potencia
	fmt.Println("2 eleveado a 5: ", potencia(2,5))
	fmt.Println("3 elevado a 4", potencia(3, 4))
}