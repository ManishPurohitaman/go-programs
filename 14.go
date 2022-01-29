//passing array as a parameter
package main
import "fmt"
func main(){
	 var b = []int{10,0,30,40,53}
	 var avg float32
	 avg = getavg(b, 5)

	 fmt.Println(avg)
}

func getavg(arr []int, size int)float32{
	 var i,sum int
	 var avg float32
	 for i=0;i<size;i++{
		 sum += arr[i]
	 }

	 avg = float32(sum)  / float32(size)

	 return avg
}
