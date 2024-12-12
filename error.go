package main

func panicMethod() {

	defer func() {

		println("Defer 1")

		println("Before recover")

		r := recover()

		if r != nil {
			println("After recover")
		}
	}()

	defer func() {
		println("Defer 2")
	}()

	println("Before Panic")

	panic("Error Occur")

	println("After Panic")
}

func main() {

	println("before panic method call")

	panicMethod()

	println("after panic methos call")
}
