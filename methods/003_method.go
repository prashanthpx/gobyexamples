package main
import "fmt"

type Value int

func (v Value) add() int {
	return int(v + 10)
}

func main() {
	//Passing value of 10
	val := Value(10)
	fmt.Printf(" calling add(): ret val = %d\n", val.add())

}