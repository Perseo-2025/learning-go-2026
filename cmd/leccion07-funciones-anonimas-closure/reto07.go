package main
import "fmt"


// 1. Recibe un precio y una función que calcula el descuento sobre ese precio
func aplicarDescuento(precio float64, descuento func(float64) float64) float64 {
	return descuento(precio)
}
// practimante hace dos cosas a la vez, recibe y luego calcula

// 4. Closure: acumula u total cada vez que se le pasan un número nuevo
func crearAcumulador() func(int) int {
	total := 0 // estsa varibale es inicializada y vive "atrapada" dentro del cierre "closure"
	return func(numero int) int {
		total += numero
		return total
	}
}	

func main(){

	//2. Dos funciones anónimas de descuento
	descuento10 := func(precio float64) float64 {
		return precio - (precio * 0.10)
	}

	descuento25 := func(precio float64) float64 {
		return precio - (precio * 0.25)
	}

	// 3. Probando aplicarDescuento con un precio de 100.0
	precioBase := 100.0
	fmt.Printf("Precio con 10%% de descuento: %.2f\n", aplicarDescuento(precioBase, descuento10))
	fmt.Printf("Precio con 25%% de descuento: %.2f\n", aplicarDescuento(precioBase, descuento25))

	// 5. Probando el acumulador con 3 llamadas seguidas 
	acumulador := crearAcumulador()
	fmt.Println("Acumulado tras sumar 5: ", acumulador(5)) 
	fmt.Println("Acumulado tras sumar 3: ", acumulador(3)) 
	fmt.Println("Acumulado tras sumar 10: ", acumulador(10)) 

}