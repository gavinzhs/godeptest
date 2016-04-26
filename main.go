package main

import (
	"fmt"
	"github.com/tpackage"
)

func main() {
	print("start testg")

	fmt.Println("包的调用")
	fmt.Printf("5 + 5 is %d", 12)

	fmt.Println("test github package add:", tpackage.Add(1,2))
}
