//arrays
package main
import "fmt"
func main(){
	 var arr1 = [5]int{1,2,3,4,5}
	 var arr2 = []int{0,1,2,3,4,5}
	 var arr3 = [][]int{{0,1},{2,3}}; // multi-dimensional array

	 for i:=0; i<5; i++{
		 fmt.Println(arr1[i])
	 }
    
	 fmt.Println()
	 for i:=0; i<5; i++{
		 fmt.Println(arr2[i])
	 }
	 fmt.Println()
	 for i:=0; i<2; i++{
		 for j:=0; j<2; j++{
			 fmt.Println(arr3[i][j])
		 }
	 }
}

