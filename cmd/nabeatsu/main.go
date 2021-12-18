package main

import (
	"fmt"
	"os"

	"github.com/furusax0621/go-nabeatsu"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`
世界のナベアツは3の倍数と3のつく数字のときだけ阿呆になります。

Usage:
	nabeatsu <number>`,
		)
		os.Exit(0)
	}
	fmt.Println(nabeatsu.BuildFoolMessage(os.Args[1]))
}
