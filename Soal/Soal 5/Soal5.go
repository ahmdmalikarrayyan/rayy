package main

import "fmt"

func main() {
	var i1, i2, i3, i4, i5 int
	var c1, c2, c3 rune
	fmt.Scanf("%d %d %d %d %d\n", &i1, &i2, &i3, &i4, &i5)
	fmt.Scanf("%c%c%c", &c1, &c2, &c3)
	fmt.Printf("%c%c%c%c%c\n", i1, i2, i3, i4, i5)
	fmt.Printf("%c%c%c\n", c1+3, c2+3, c3+3)
}