package main

import "fmt"

type ByteSize float64

const (
	first = 1 << (10 * iota)
	second
	third
)

func main() {

	 //OneRunAll();

	//TwoRunAll()

	branching()

	loops()

	switches()

}

func TwoRunAll() {

	branching()

	loops()

	switches()
}

func switches() {

	capacities := []float64{30, 30, 42, 420, 22, 66}

	activePlants := []int{0, 1, 4}

	fmt.Print("Give us a number to print:  (1 or 2)  ")

	var numberToPrint string

	fmt.Scan(&numberToPrint)

	switch numberToPrint {
	case "1":
		for idx, cap := range capacities {
			fmt.Printf("Plant %d capacity: %.0f\n ", idx, cap)
		}
	case "2":
		capacity := 0.
		for _, plantId := range activePlants {
			capacity += capacities[plantId]
			fmt.Println(capacity)
		}
	default:
		fmt.Println("Unknown option")
	}
}


func loops() {

	m := make(map[string]string)

	m["first"] = "foo"
	m["second"] = "bar"
	m["third"] = "bob"

	for k, v := range m {
		println(k, v)
	}

	s := []string{"one", "duece", "tres"}

	for idx, v := range s {
		println(v, idx)
	}


	bar := 1

	for {
		bar++
		println(bar)

		if bar > 420 {
			break
		}
	}


	foo := 10

	for i:=0; i<foo; i++ {
		println(i)
	}
}

func branching() {
	bar := 5

	switch {
	case bar == 1:
		println(1)
	case bar > 2:
		println(2)
	}


	// Initialize the value to be tested right in the if statement
	if foo := 2; foo == 1 {
		println("yup")
	} else {
		println("not 1")
	}
}



/* Section 1 */

func OneRunAll() {
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