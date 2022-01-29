//for loop,break,continue
package main
import "fmt"
func main(){
	 for i:=0; i<=11; i++{

		if i == 5{
			continue
		}
		fmt.Println("value of i is",i)

		if i == 10{
			break
		}
	 }
}