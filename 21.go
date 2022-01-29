//maps
package main

import "fmt"

func main() {
     var mymap map[string]int
	 mymap = map[string]int{}
	 mymap["a"] = 12
	 mymap["b"] = 120
	 mymap["c"] = 1200
	 mymap["d"] = 12000
	 mymap["e"] = 120000

	 delete(mymap,"d")
	 fmt.Println(mymap)
}