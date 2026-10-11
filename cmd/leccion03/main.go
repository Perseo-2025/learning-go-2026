package main
import "fmt"

func main(){

	// declaramos SIN asignar valor, para ello usaremos var no :=
	var contador int 
	var precio float64
	var nombre string
	var activo bool
	var letra rune

	// Imprimimos el Zero Value de cada uno
	fmt.Println("---Zero Value---")
	//aqui entra la magia
	fmt.Printf("int  -> %d\n", contador)
	fmt.Printf("flat64  -> %f\n", precio)
	fmt.Printf("string  -> '%s' (longitud: %d)\n", nombre, len(nombre))
	fmt.Printf("bool -> %t\n", activo)
	fmt.Printf("rune   -> %d\n", letra)

	// Un string vacío No es lo mismo que un Undefine 
	nombre = nombre + "Rey"
	fmt.Println("Despues de concatenar: ", nombre)

	// Con init: podemos sumarle directamente sin inicializarlo antes
	contador = contador + 1
	fmt.Println("Contador tras sumar 1:", contador)

	//Puntero: aqui SI aparece 'nil' (el "null" de GO)
	var puntero *int
	fmt.Println("Puntero inicializar:", puntero) //<nil>

	// verficar si un puntero es nil es MUY común en GO (evita pánicos)
	if(puntero == nil){
		fmt.Println("El puntero no apunta a nada todavia")
	}


}