//extend and shrink the slice
//length and capacity of slice
package main
import "fmt"
func main(){
	  
     arr := []int{1,2,3,4,5,6,7,8,9,10}
	 fmt.Println(arr)

	 a := arr[:4]
	 fmt.Println(a, len(a), cap(a))

	 arr = arr[:6]
	 fmt.Println(arr)

	 arr = arr[:3]
	 fmt.Println(arr)
}