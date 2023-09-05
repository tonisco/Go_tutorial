package main

import (
	"fmt"
)

func bulkSend(numMessages int) float64 {
	cost := 0.0
	for i := 0; i < numMessages; i++ {
		cost += 1.00 + (0.01 * float64(i))
	}

	return cost
}

// don't edit below this line

func test(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
	test(40)
	test(50)
}
