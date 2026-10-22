package main

import "fmt"

// Caso clasico: factorial
func factorial(n int) int {
	if n <= 1 {
		return 1 // caso base --sin esto, la recursion nunca termina
	}
	return n * factorial(n-1)
}

// Caso Fibonacci recursivo
func fibonacci(n int) int{
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

// Caso de uso real : estructura jerárquica
// esto simula algo muy importante en backend: una categoria con subcategorias
// anidadas ("Electronica > Celulares > "smartphones")
type Categoria struct {
	Nombre        string
	SubCategorias []Categoria // una categoria puede contener mas categorias
}

//contarCategorias recorre TODO el árbol, sin importar cúantos niveles tenga.
// Esto sería MUY DIFICIL de hacer con un 'for' simple, porque no sabes de antemano cuantos niveles tenga de profundidad
func contarCategorias(c Categoria) int{
	total := 1 // me cuento a mí mismo

	// recorro cada subcategoria y le aplico la MISMA función (recursión)
	for _, sub := range c.SubCategorias {
		total += contarCategorias(sub)
	}

	return total
}

func maestro() {

	// ---Factorial---
	fmt.Println("Factorail de 5: ", factorial(5)) // 5*4*3*2*1 = 120

	// Finbonacci
	fmt.Println("--- Fibonacci ---")
	for i := 0; i <= 8; i++ {
		fmt.Printf("fib(%d) = %d\n", i, fibonacci(i))
	} 


	fmt.Println("--Estructura Gerarquica---")
	arbol := Categoria {
		Nombre: "Electrónica",
		SubCategorias: []Categoria {
			{
				Nombre: "Celulares",
				SubCategorias: []Categoria{
					{Nombre: "Smartphones"},
					{Nombre: "Tablets"},
				},
			},
			{Nombre: "Laptos"},
		},
	}

	fmt.Println("Total de categorías (incluyendo anidadas):", contarCategorias(arbol))
}
