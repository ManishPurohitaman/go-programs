// Call by reference
package main

import "fmt"

func main() {
var a, b int
a = 10
b = 20  

fmt.Printf("before swap -> %d and %d\n", a, b)
swap(&a, &b)
fmt.Printf("after swap -> %d and %d\n", a, b)
}
func swap(p1 *int, p2 *int) {
var t int
t = *p1
*p1 = *p2
*p2 = t
}