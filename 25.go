//Packages
//Every go program is made of Packages
//Program starts running in package main
//are golang packages like c/c++ header files?
//import statement imports i.e brings the package from its source to destination wiz. main function 
package main

import (
	   "fmt"
	   "math"
	   "math/rand"
)

func main() {
	   c := math.Exp2(5)
       
	   fmt.Println(c)
	   fmt.Println(rand.Int() % 100)
}