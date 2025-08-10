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

/*
Output
40
0x<addr>
50
After calling change() fn  50
*/

/*
Code Explanation:
- Purpose: Pass pointer to function to modify caller’s variable
- change receives *int and assigns 50 via dereference; caller observes updated value
*/
