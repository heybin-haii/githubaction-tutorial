package main

import "fmt"

type Color struct {
	color string
	code  int
}

func ValueColor(color Color) {
	color.code += 10
}

func ReferenceColor(color *Color) {
	(*color).code += 10
}

func main() {
	color := Color{"Green", 10}
	ValueColor(color)      // 결과값: Green, 10
	ReferenceColor(&color) // 결과값: Tommy, 20
	fmt.Println("")
}
