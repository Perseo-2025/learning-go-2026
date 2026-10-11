package main
import "fmt"


func main(){

	// forma explicita

	nombre := "Rey Perseo"
	var edad int = 23
	var rol string
	rol = "Full Satck Developer"
	var experience bool = false

	fmt.Printf("Nombre completo: %s %s\n", nombre)
	fmt.Printf("Tengo: %d años\n", edad)
	fmt.Printf("Mi rol en la empresa es %s ", rol)
	fmt.Printf("Mi experiencia en go es %t", experience)

}