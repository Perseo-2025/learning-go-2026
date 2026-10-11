package main

import (
	"fmt"
)

func main() {

	edad := 23

	if edad >= 18 {
		fmt.Println("Puedes entrar a la discoteca")
	} else if edad >= 13 {
		fmt.Println("Eres adolescente")
	} else {
		fmt.Println("Eres menor de edad")
	}

	// IF con inicializador
	// Sintaxis: if <varibale-temparal>; <condición> {...}
	// EL punto y coma separa la declaración de la condición.
	// 'edadUsuario' solo existe Dentro de este bloque if/else

	if edadUsuario := 15; edadUsuario >= 18 {
		fmt.Println("Acceso Permitido")
	} else {
		// edadUsuario sigue disponible dentro de las llaves
		fmt.Println("Acceso denegado, tiene ", edadUsuario, "años")
	}
	// Despues la variable iniciada ya no existe es decir edadUsuario YA NO EXISTE

	//El patron de dos valores funciona muchisimo
	inventario := map[string]int{"manzanas": 10}

	if cantidad, existe := inventario["manzanas"]; existe {
		fmt.Println("Hay ", cantidad, " manzanas")
	} else {
		fmt.Println("No hay manzana registradas")
	}

	// SWITCH
	// en GO no se necesita breack
	dia := 3

	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	default:
		fmt.Println("Otro día")
	}

	// SWITCH sin variable (reemplza cadenas largas de if/else if)

	nota := 85
	switch {
	case nota >= 90:
		fmt.Println("Calificacion: A")
	case nota >= 80:
		fmt.Println("Calificación: B")
	case nota >= 70:
		fmt.Println("Calificacion: C")
	default:
		fmt.Println("Calificacion: F")
	}

	// SWTICH con múltiples valores en un mismo case (separados por coma)
	letra := "a"
	switch letra {
	case "a", "e", "i", "o", "u":
		fmt.Println("Es una vocal")
	default:
		fmt.Println("Es una consonante")
	}

	//Reto
	fmt.Println("==Resultado Reto==")

	// Declara una variable preci := 45.0
	/* 	Declara una variable precio := 45.0.
	Usa un if con inicializador para calcular un descuento := 0.0 que sea 0.20 (20%) si precio > 50, o 0.10 (10%) en caso contrario, e imprime el precio final (precio - precio*descuento).
	Usa un switch (sin variable, con condiciones) para clasificar el precio final en: "Barato" (< 30), "Medio" (30-60), "Caro" (> 60).
	Usa un switch con una variable rolUsuario := "admin" que imprima permisos distintos para "admin", "editor", "viewer" (usa default para cualquier otro rol). */
	
	//1. Precio base
	precio := 45.0

	// 2. If con el inicializador: 'descuento' solo existe dentro de este bloque
	// y de su else. Según la regla: >50 -> 20%, si no ->10%
	var precioFinal float64

	if descuento := 0.0; precio > 50 {
		descuento = 0.20
		precioFinal = precio - precio * descuento
		fmt.Printf("Descuento aplicado %.0f%%\n", descuento*100)
	} else {
		descuento = 0.10
		precioFinal = precio - precio * descuento
		fmt.Printf("Descuento aplicado: %.0f%%\n", descuento*100)
	}

	fmt.Printf("Precio original: %.2f\n", precio)
	fmt.Printf("Precio final: %.2f\n", precioFinal)

	switch {
	case precioFinal < 30:
		fmt.Println("Categoria: Barato")
	case precioFinal <= 60:
		fmt.Println("Categoría: Medio")
	default:
		fmt.Println("Categoria: Caro")
	}

	rolUsuario := "admin"

	switch rolUsuario {
	case "admin" : 
		fmt.Println("Permisos: acceso total (CRUD)")
	case "editor":
		fmt.Println("Permisos: crear y editar contenido")
	case "viewer":
		fmt.Println("Permisos: solo lectira")
	default:
		fmt.Println("Permisos: rol no reconocido, sin acceso")
	}


}

