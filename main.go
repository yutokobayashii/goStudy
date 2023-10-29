package main

import "fmt"

func Plus(x int, y int) {

	fmt.Println(x + y)

}

func Double(prince int) (result int) {
	result = prince * 10

	return result
}

func main() {

	Plus(100, 50)

	i := Double(3000)

	fmt.Println(i)

}
