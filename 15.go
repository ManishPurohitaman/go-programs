//pointers
package main
import "fmt"
func main(){
	var num int
	var ptr *int
	ptr = &num

	fmt.Printf("value and address of num variable are %d and %x\n", ptr)
}
