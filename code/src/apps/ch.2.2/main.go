// Example code for Chapter 2.2 from "Build Web Application with Golang"
// Purpose: Goes over the assignment and manipulation of basic data types.
// Código de ejemplo para el capítulo 2.2 del "Build aplicación Web con Golang"
// Propósito: asignación y la manipulación de los tipos de datos básicos.
package main

import (
	"errors"
	"fmt"
)

// constantes
const Pi = 3.1415926

// booleanos por default son `false`
var isActive bool                   // global variable
var enabled, disabled = true, false // omitiendo el tipo de variable

// agrupando definiciones
const (
	i         = 1e4
	MaxThread = 10
	prefix    = "astaxie_"
)

var (
	frenchHello string      // forma basica de definir string
	emptyString string = "" // define un string con una cadena vacia
)

func show_multiple_assignments() {
	fmt.Println("show_multiple_assignments()")
	var v1 int = 42

	// Define tres variables de tipo "int", e inicializa con 3 valores.
	// vname1 is v1, vname2 is v2, vname3 is v3
	var v2, v3 int = 2, 3

	// `:=` solo para trabajar con funciones
	// `:=` Forma corta para definir variables
	//  especificando el tipo usando `var`.
	vname1, vname2, vname3 := v1, v2, v3

	// `_` si queremos omitir un valor que devuelva una funcion.
	_, b := 34, 35

	fmt.Printf("vname1 = %v, vname2 = %v, vname3 = %v\n", vname1, vname2, vname3)
	fmt.Printf("v1 = %v, v2 = %v, v3 = %v\n", v1, v2, v3)
	fmt.Println("b =", b)
}
func show_bool() {
	fmt.Println("show_bool()")
	var available bool // local variable
	valid := false     // asignacion corta
	available = true   // asignando valor a una variable

	fmt.Printf("valid = %v, !valid = %v\n", valid, !valid)
	fmt.Printf("available = %v\n", available)
}
func show_different_types() {
	fmt.Println("show_different_types()")
	var (
		unicodeChar rune
		a           int8
		b           int16
		c           int32
		d           int64
		e           byte
		f           uint8
		g           int16
		h           uint32
		i           uint64
	)
	var cmplx complex64 = 5 + 5i

	fmt.Println("Default values for int types")
	fmt.Println(unicodeChar, a, b, c, d, e, f, g, h, i)

	fmt.Printf("Value is: %v\n", cmplx)
}
func show_strings() {
	fmt.Println("show_strings()")
	no, yes, maybe := "no", "yes", "maybe" //declaracion corta tres 					      //variables:tres valores
	japaneseHello := "Ohaiyou"
	frenchHello = "Bonjour" // forma basica de asignar valor a una variable

	fmt.Println("Random strings")
	fmt.Println(frenchHello, japaneseHello, no, yes, maybe)

	// este simbolo(acento)`, permite encrerrar una cadena larga
	fmt.Println(`This 
	is on
	multiple lines`)
}
func show_string_manipulation() {
	fmt.Println("show_string_manipulation()")
	var s string = "hello"

	//You can't do this with strings
	//s[0] = 'c'

	s = "hello"
	c := []byte(s) // convert string to []byte type
	c[0] = 'c'
	s2 := string(c) // convert back to string type

	m := " world"
	a := s + m

	d := "c" + s[1:] //no se puede cambiar los valores de cadena por índice, pero se puede obtener valores en lugar
	fmt.Printf("%s\n", d)

	fmt.Printf("s = %s, c = %v\n", s, c)
	fmt.Printf("s2 = %s\n", s2)
	fmt.Printf("combined strings\na = %s, d = %s\n", a, d)
}
func show_errors() {
	fmt.Println("show_errors()")
	err := errors.New("Example error message\n")
	if err != nil {
		fmt.Print(err)
	}
}
func show_iota() {
	fmt.Println("show_iota()")
	const (
		x = iota // x == 0
		y = iota // y == 1
		z = iota // z == 2
		w        // Si no hay una expresion despues del nombre de la constante,
		// se utiliza la ultima expresion por eso se dice w = iota implicito.
		// por eso w == 3, and y and x both can omit "= iota" as well.
	)

	const v = iota // si iota es llamado por otra constante, se reatablece a`0`, so v = 0.

	const (
		e, f, g = iota, iota, iota // e=0,f=0,g=0 si se coloca en una linea los valores de iota son iguales.
	)
	fmt.Printf("x = %v, y = %v, z = %v, w = %v\n", x, y, z, w)
	fmt.Printf("v = %v\n", v)
	fmt.Printf("e = %v, f = %v, g = %v\n", e, f, g)
}

// Si quiere que las funciones y las variables sean publicas se deben de nombrar
// con la primera letra Mayuscula.
// De lo contrario seran privadas.
func This_is_public()  {}
func this_is_private() {}

func set_default_values() {
	// valores por defecto.
	const (
		a int     = 0
		b int8    = 0
		c int32   = 0
		d int64   = 0
		e uint    = 0x0
		f rune    = 0   // the actual type of rune is int32
		g byte    = 0x0 // the actual type of byte is uint8
		h float32 = 0   // length is 4 byte
		i float64 = 0   //length is 8 byte
		j bool    = false
		k string  = ""
	)
}
func show_arrays() {
	fmt.Println("show_arrays()")
	var arr [10]int // un array de tipo int
	arr[0] = 42     // array is 0-based
	arr[1] = 13     // assign value to element

	a := [3]int{1, 2, 3} // define un array con 3 elementos

	b := [10]int{1, 2, 3}
	// define array con 10 elementos,
	// asigna valores a los primeros tres,y para el resto se da valor por defecto.

	c := [...]int{4, 5, 6} // use `…` si no queremos definir el numero de elementos

	//define un arrayde 2 dimensiones con 2 elementos, y cada uno con 4 elementos.
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	// declaracion y definicion corta.
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	fmt.Println("arr =", arr)
	// imprime el elemento 0 de arr=42
	fmt.Printf("The first element is %d\n", arr[0]) 
	//retorna el 10 valor = 9 elemento
	fmt.Printf("The last element is %d\n", arr[9])

	fmt.Println("array a =", a)
	fmt.Println("array b =", b)
	fmt.Println("array c =", c)

	fmt.Println("array doubleArray =", doubleArray)
	fmt.Println("array easyArray =", easyArray)
}
func show_slices() {
	fmt.Println("show_slices()")
	// define slice de 10 elementos de tipo byte
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	// define dos slice de tipo []byte
	var a, b []byte

	// define un slice "a" del array ar desde el 3 hasta el 5 elementos.
	a = ar[2:5]
	// now a has elements ar[2]、ar[3] and ar[4]

	// b es otro slice del array ar
	b = ar[3:5]
	// now b has elements ar[3] and ar[4]

	// define un array
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// define 2 slices
	var aSlice, bSlice []byte

	// some convenient operations
	aSlice = array[:3]//define un aSlice de array desde elemento 0 hasta el 3
	aSlice = array[5:]//define un aSlice de array desde elemento 5 hasta el 10  
	aSlice = array[:] //define un aSlice de array [0:10]con todo los elementos

	// slice from slice
	aSlice = array[3:7]  // aSlice tiene elementos d,e,f,g，len=4，cap=7
	bSlice = aSlice[1:3] // bSlice contiene aSlice[1], aSlice[2], osea e,f
	bSlice = aSlice[:3]  // bSlice contains aSlice[0], aSlice[1], osea d,e,f
	bSlice = aSlice[0:5] // bSlice ha expandido su rango = cap, ahora contiene 					d,e,f,g,h
	bSlice = aSlice[:]   // bSlice contiene los elementos de aSlice  d,e,f,g

	fmt.Println("slice ar =", ar)
	fmt.Println("slice a =", a)
	fmt.Println("slice b =", b)
	fmt.Println("array =", array)
	fmt.Println("slice aSlice =", aSlice)
	fmt.Println("slice bSlice =", bSlice)
	fmt.Println("len(bSlice) =", len(bSlice))
}
func show_map() {
	fmt.Println("show_map()")
	// use string como el tipo de la llave, e int como tipo para el valor
	// use `make` para iniciaizar el map
	var numbers map[string]int
	// another way to define map
	numbers = make(map[string]int)
	numbers["one"] = 1 // asigna valor a la llave "one"
	numbers["ten"] = 10
	numbers["three"] = 3

	// Initialize a map
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}

	fmt.Println("map numbers =", numbers)
	// Imprime el valor de la llave "three" osea 3
	fmt.Println("The third number is: ", numbers["three"]) 
	// map retorna dos valores. Y podemos utilizar el segundo valor ok para saber 		// si existe un elemento en nuestro map.
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C") // delete element with key "c"
	fmt.Printf("map rating = %#v\n", rating)
}
func main() {
	show_multiple_assignments()
	show_bool()
	show_different_types()
	show_strings()
	show_string_manipulation()
	show_errors()
	show_iota()
	set_default_values()
	show_arrays()
	show_slices()
	show_map()
}
