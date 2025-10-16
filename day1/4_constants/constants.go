package main

import "fmt"

func main() {
	const a int = 10
	const (
		b string  = "hello"
		c float32 = 10.5
		d bool    = true
	)
	// Multiple constant declaration
	const e, f, g int = 1, 2, 3
	const l, m, n = "a", "b", "c"

	fmt.Println(b, c, d)
	fmt.Println(e, f, g)
	fmt.Println(l, m, n)
	fmt.Println(a)
	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of b: %T\n", b)
	fmt.Printf("Type of c: %T\n", c)
	fmt.Printf("Type of d: %T\n", d)
	fmt.Printf("Type of e: %T\n", e)
	fmt.Printf("Type of l: %T\n", l)
}
