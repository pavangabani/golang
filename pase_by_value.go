package main

type Person struct {
	name string
	age  int
}

func modifyInt(val int) {
	val = 2
}

func modifyStruct(person Person) {
	person.name = "gabani"
}

func modifyArray(array [2]int) {
	array[1] = 20
}

func main() {

	//int pass by value
	num := 1
	modifyInt(num)
	println(num)

	//struct pass by value
	person := Person{
		name: "pavan",
		age:  24}

	modifyStruct(person)
	println(person.name)

	//Array pass by value
	array := [2]int{1, 2}
	modifyArray(array)
	println(array[1])

}
