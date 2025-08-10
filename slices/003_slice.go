package main

import "fmt"

func main() {
	a := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun", "1", "2", "3"}

	slice1 := a[0:]
	slice2 := a[3:]

	fmt.Println("------- Before Modifications -------")
	fmt.Println("a  = ", a)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)

	//slice1[0] = "modTUE"
	//slice1[1] = "modWED"
	//slice1[2] = "modTHU"

	//slice2[1] = "FRIDAY"

	fmt.Println("\n-------- After Modifications --------")
	fmt.Println("a  = ", a)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	//for i := 0; i < len(slice1); i++ {
	//for i, val := range slice1 {
	for len(slice1) >= 1 {
		index := 0
		//if len(slice1) > 1 {
		//fmt.Printf("i = %v, val: %v ", i, val)
		copy(slice1[index:], slice1[index+1:])
		//fmt.Println("line 27 slice1 = ", slice1)
		slice1 = slice1[:len(slice1)-1]
		fmt.Println("line 29 slice1 = ", slice1)
		fmt.Printf("line 30 len: %d", len(slice1))
		fmt.Printf(" ************ \n")
	}
}

/*
Output
------- Before Modifications -------
a  =  [Mon Tue Wed Thu Fri Sat Sun 1 2 3]
slice1 =  [Mon Tue Wed Thu Fri Sat Sun 1 2 3]
slice2 =  [Thu Fri Sat Sun 1 2 3]

-------- After Modifications --------
a  =  [Mon Tue Wed Thu Fri Sat Sun 1 2 3]
slice1 =  [Mon Tue Wed Thu Fri Sat Sun 1 2 3]
slice2 =  [Thu Fri Sat Sun 1 2 3]
line 29 slice1 =  [Tue Wed Thu Fri Sat Sun 1 2 3]
line 30 len: 9 ************
line 29 slice1 =  [Wed Thu Fri Sat Sun 1 2 3]
line 30 len: 8 ************
line 29 slice1 =  [Thu Fri Sat Sun 1 2 3]
line 30 len: 7 ************
line 29 slice1 =  [Fri Sat Sun 1 2 3]
line 30 len: 6 ************
line 29 slice1 =  [Sat Sun 1 2 3]
line 30 len: 5 ************
line 29 slice1 =  [Sun 1 2 3]
line 30 len: 4 ************
line 29 slice1 =  [1 2 3]
line 30 len: 3 ************
line 29 slice1 =  [2 3]
line 30 len: 2 ************
line 29 slice1 =  [3]
line 30 len: 1 ************
line 29 slice1 =  []
line 30 len: 0 ************
*/

/*
Code Explanation:
- Purpose: Show slicing and in-place element removal using copy/resize
- slice1 := a[0:], slice2 := a[3:]
- While loop repeatedly removes first element from slice1 by shifting with copy and shortening len
- Demonstrates how slices share underlying arrays and how length changes
*/
