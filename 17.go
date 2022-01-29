// structure
package main
import "fmt"

type Books struct{
	 id int
	 title string
} 
func main() {
	 var b1 Books
	 b1.id = 1
	 b1.title = "book"
	 fmt.Println(b1)
}