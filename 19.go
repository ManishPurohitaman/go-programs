//slice
package main
import "fmt"

func main(){
	 var arr = []int{1,2,3,4,5,6,7}
	 my_slice := arr[0:5]

	 fmt.Println(my_slice,len(my_slice))
}
//1.can store elements of one type only
//2.size is flexible, unlike array
//3.internally implements arr(slice is a reference to an underlying array)
// slice has 3 things - a pointer, a length and a capacity