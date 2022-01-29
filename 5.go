// mixed type declaration
package main
import "fmt"

func main(){
	 var a,b,c,d = 34,45.6,"namaste duniya!",'a'

	 fmt.Println(a)
	 fmt.Println(b)
	 fmt.Println(c)
	 fmt.Println(d,"\n")

	 fmt.Printf("type of a is %T\n",a)
	 fmt.Printf("type of b is %T\n",b)
	 fmt.Printf("type of c is %T\n",c)
	 fmt.Printf("type of d is %T",d)
}