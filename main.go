package main

import "fmt"

type ByteSize float64

const (
	first = 1 << (10 * iota)
	second
	third
)

func main() {

	 imaginaryNumbers()

	 constants()

	 Arrays()

	 maps()

}

func imaginaryNumbers() {
	myComplex := complex(2, 3)
	println(myComplex)
	println(real(myComplex))
	println(imag(myComplex))
}

func constants() {
	println(first)
	println(second)
	println(third)


	var myInt int
	myInt = 42

	println(myInt)

	var myFloat32 float32 = 42.

	println(myFloat32)

	myString := "Hello Go"
	println(myString)
}

func Arrays() {
	myArray := [...]int{42, 22, 99}

	mySlice := myArray[:]

	fmt.Println(myArray)

	// Append to Array
	mySlice = append(mySlice, 420)

	fmt.Println(myArray)
	fmt.Println(mySlice)

	newSlice := []float32{12., 15., 20.}
	fmt.Println(len(newSlice))
}

func maps() {

	myMap :=  make(map[int]string)

	fmt.Println(myMap)

	myMap[42] = "foo"
	myMap[12] = "Bar"

	fmt.Println(myMap)
	fmt.Println(myMap[99])

}