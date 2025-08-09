package main

import "fmt"

func main() {
	x := 40

	fmt.Println(x)
	fmt.Println(&x)
	change(&x)
	fmt.Println("After calling change() fn ", x)
}

func change(x *int) {
	*x = 50
	fmt.Println(*x)

}