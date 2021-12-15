package main

import (
	"fmt"
	"os"

	"github.com/furusax0621/go-nabeatsu/nabeatsu"
)

func main() {
	num := os.Args[1]
	fmt.Println(nabeatsu.BuildFoolMessage(num))
}
