package main

func modifySlice(slice []int) {
	slice[0] = 1000
}

func modifyMap(m map[string]int) {

	m["pavan"] = 24
}

func sendValue(channel chan int) {

	channel <- 300
}

func main() {

	//pass by reference in slice
	slice := []int{1, 2, 3}
	modifySlice(slice)
	println("After Call", slice[0])

	//pass by reference in map
	var myMap map[string]int
	myMap = make(map[string]int)
	myMap["pavan"] = 12
	myMap["smit"] = 7
	modifyMap(myMap)
	println(myMap["pavan"])

	//paas by reference in channel
	ch := make(chan int, 2)
	sendValue(ch)
	println(<-ch)

}
