package main

type Car struct {
	name string
	id   int
}

type Audi struct {
	Car
	accessory string
	id        string
}

//--------------------------------------------------------------------------------------------------------------

type ol struct {
}

func (o ol) overloading(v interface{}) {
	switch v.(type) {

	case int:
		println("it is int")

	case string:
		println("it is string")
	}
}

//--------------------------------------------------------------------------------------------------------------

type Animal struct {
	name string
	id   int
}

func (a Animal) speak() {
	println("Animal speak")
}

func (a Animal) animal() {
	println("In animal test ")
}

type Cat struct {
	Animal
}

func (c Cat) speak() {
	println("Cat speak")
}

func main() {

	//Overriding
	var car1 Car
	car1.id = 1
	car1.name = "test"

	var audi Audi
	audi.id = "overriding ID"
	audi.name = "struct"
	audi.accessory = "pavan"

	//Overloading
	var o ol
	o.overloading(1)
	o.overloading("pavan")

	//Inheritance
	var cat1 Cat
	cat1 = Cat{}
	cat1.speak()
	cat1.animal()
}
