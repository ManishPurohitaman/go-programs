//make function to create slices
package main
import "fmt"

func main(){
	 arr := make([]int,5)
	 for i:=0 ; i < 5 ; i++{
		 arr[i] = i * i
		 fmt.Println(arr[i])
	 }
}