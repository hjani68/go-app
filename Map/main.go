package main

import "fmt"

func main() {

	colors := make(map[string]string)
	colors["White"] = "#ffffff"

	colorsInt := make(map[int]string)
	colorsInt[10] = "#ffffff"
	colorsInt[11] = "#fffff0"

	delete(colorsInt, 11) // deleting records from map using index

	fmt.Println(colors)
	fmt.Println(colorsInt)

}
