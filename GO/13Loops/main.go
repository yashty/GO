package main

import "fmt"

func main() {
	fmt.Println("Loop in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue == 2 {
			goto lco
		}
		fmt.Println("value is:", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping at LearCodeonline.in")
}
