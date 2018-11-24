package integers

import "fmt"

func main() {
	fmt.Println(Add(3, 2))
}

//Add takes two integers and returns the sum
func Add(x, y int) int {
	return x + y
}
