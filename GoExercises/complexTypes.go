package main

import "fmt"

func main() {
	contacts := make(map[string]string)
	
	//adding key value pairs
	contacts["Arya"] = "789098765"
	contacts["Swati"] = "789456765"
	contacts["Vikas"] = "7891346765"
	
	fmt.Println(contacts)
	fmt.Println(contacts["Swati"])
	
	//deleting from map
	delete(contacts,"Vikas")
	
	//looping over map
	for k,v := range contacts {
		fmt.Println(k, ":", v)
	}
	
}