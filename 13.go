//for-each loop
//dynamic Type Declarations
package main
import "fmt"
func main(){
	 var arr = []int{1,2,3,4,5}
	 for idx,num := range arr{
		 fmt.Println(idx,num)
	 }
}
