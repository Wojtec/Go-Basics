package main

import "fmt"

func main() {

	// variables go
	var t = 1
	fmt.Println(t)

	var a, c int = 2, 4
	fmt.Println(a + c)

	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)

	f := "hey"
	fmt.Println(f)

	// pointer

	i, j := 20, 400

	p := &i                    // point to i
	fmt.Println("our p :", *p) // read i through the pointer
	*p = 40
	fmt.Println("now i is equal 40: ", i) // new value

	p = &j
	*p = *p / 2
	fmt.Println("our p have value j and divide 2 :", j)

	// functions
	fmt.Println("Show result from function : ", add(i, j))

	// loop for
	sum := 0
	for i := 0; i <= 10; i++ {
		sum = i
	}
	fmt.Println("Result from loop:", sum)

	//array
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println("number of arr[0]", arr[0])
	fmt.Println("numbers in array: ", arr)

	//slices
	numbers := [4]int{1, 2, 3, 4}
	var q []int = numbers[0:4]
	fmt.Println("slices numbers : ", q)

	//maps
	type Vertex struct {
		Lat, Long float64
	}
	var n map[string]Vertex
	n = make(map[string]Vertex)
	n["bell labs"] = Vertex{
		40.40342, -10.3242,
	}
	fmt.Println("Maps: ", n["bell labs"])
	// structs is a coolection of fileds
	type Test struct {
		X int
		Y int
	}

	fmt.Println("Struct test : ", Test{1, 2})
	// definde types struct
	type Car struct {
		name  string
		year  int
		speed int
	}
	var model Car
	model.name = "mini"
	model.year = 2000
	model.speed = 140
	fmt.Println("Type struct:", model)
}

func add(i int, j int) int {

	//condition if
	if i+j > 200 {
		var a = i - j
		fmt.Println("Condition if is working:", a)

	}
	return i + j
}
