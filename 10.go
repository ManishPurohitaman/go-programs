//functions
package main
import "fmt"
func max(num1, num2 int) int{
	var result int = 90
	fmt.Printf("max value is %d\n",num1)
	fmt.Printf("max value is %d\n",num2)

	if(num1 > num2){
		result = num1
	}else{
		result = num2
	}
	return result
}
func main(){
	 var a int = 100
	 var b int = 200
	 
	 var ret int = max(a,b)
	 fmt.Printf("max value is %d\n",ret)
}

