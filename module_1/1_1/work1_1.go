package main

import "fmt"

func changeArray() []string {
	var arr = []string{"I", "am", "stupid", "and", "weak"}
	for index, val := range arr {
		switch val {
		case "stupid":
			arr[index] = "good"
		case "weak":
			arr[index] = "strong"
		}
	}
	return arr
}

func main() {
	newArr := changeArray()
	fmt.Println(newArr)
}
