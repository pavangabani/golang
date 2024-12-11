package main

func modifyIntPointer(val *int) {
	*val = 30
}
func main() {

	num := 10
	modifyIntPointer(&num)
	println(num)
}
