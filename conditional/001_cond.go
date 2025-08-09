package main

import (  
    "fmt"
)

func main() {  
	// Here we can dee before the condition is met, 
	// statement is execured. In this case num: =10
    if num := 10; num % 2 == 0 { //checks if number is even
        fmt.Println(num,"is even") 
    }  else {
        fmt.Println(num,"is odd")
    }
}