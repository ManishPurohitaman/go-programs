//dynamic Type Declarations
package main
import "fmt"
func main(){
	 var x float64 = 20.256
	 y := 42  // dynamic declaration
     z := "Hello World"

	 fmt.Println(x)
	 fmt.Println(y)
	 fmt.Println(z,"\n")

	 fmt.Printf("x is of type %T\n",x)
	 fmt.Printf("y is of type %T\n",y)
	 fmt.Printf("z is of type %T",z)
}
