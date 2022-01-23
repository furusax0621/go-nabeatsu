package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/furusax0621/go-nabeatsu"
	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 && !hasPipeData() {
		fmt.Println(`
世界のナベアツは3の倍数と3のつく数字のときだけ阿呆になります。

Usage:
	nabeatsu <number>`,
		)
		os.Exit(0)
	}

	if hasPipeData() {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal("パイプから入力を受け取れませんでした")
		}

		list := chopAll(strings.Split(string(stdin), "\n"))
		if list[len(list)-1] == "" {
			list = list[:len(list)-1]
		}

		for _, s := range list {
			fmt.Println(nabeatsu.GetFoolExpression(s))
		}
		os.Exit(0)
	}
	fmt.Println(nabeatsu.GetFoolExpression(os.Args[1]))
}

// hasPipeData は、パイプから標準入力を受け取っているかどうかを返す。
func hasPipeData() bool {
	return !term.IsTerminal(syscall.Stdin)
}

// chop は、文字列の末尾から改行コードを削除する。
func chop(line string) string {
	if strings.HasSuffix(line, "\n") {
		return strings.TrimLeft(line, "\n")
	}
	return line
}

// chop は、文字列リストから文字列を取り出し、それら文字列末尾の改行コードを削除する。
func chopAll(lines []string) []string {
	var newLines []string
	for _, v := range lines {
		newLines = append(newLines, chop(v))
	}
	return newLines
}
