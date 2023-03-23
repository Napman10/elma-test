package main

import "fmt"

func main() {
	for {
		var num int

		fmt.Println("Input task:")

		if _, err := fmt.Scanf("%d", &num); err != nil {
			return
		}

		switch num {
		case 1:
			ShowTask1()
		case 2:
			ShowTask2()
		case 3:
			ShowTask3()
		case 4:
			ShowTask4()
		default:
			break
		}
	}
}
