//variadic function
package main
import "fmt"

func variadic(params ...int){
	 fmt.Printf("%T\t%v\n",params,params)
}
func main(){
    variadic(1,1,2,3,34,34,3)
}