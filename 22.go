//slice of slice
package main
import "fmt"
func main(){
	 arr := []int{1,2,3,4,5,6,7,8,9,10}

	 arr = arr[1:4]
	 fmt.Println(arr)

	 arr = arr[:3]
	 fmt.Println(arr)

	 arr = arr[:]
	 fmt.Println(arr)

     arr = arr[0:]
	 fmt.Println(arr)
}