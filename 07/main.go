package main

import "fmt"

func main() {
	str := create_template_str("12", "気温", "22.4")
	fmt.Println(str)
}

func create_template_str(x string, y string, z string) string {
	return x + "時の" + y + "は" + z
}
