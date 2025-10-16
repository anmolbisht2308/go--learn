package main

import "fmt"

func main() {
	const a int = 10
	var b string = "hello"
	var c float32 = 10.5
	var d bool = true

	// Shorthand declaration
	e := 20
	f := "world"
	g := 20.5
	h := false

	// Multiple variable declaration
	var i, j, k int = 1, 2, 3
	l, m, n := "a", "b", "c"

	fmt.Println(b, c, d)
	fmt.Println(e, f, g, h)
	fmt.Println(i, j, k)
	fmt.Println(l, m, n)

	fmt.Println(a)

	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of b: %T\n", b)
	fmt.Printf("Type of c: %T\n", c)
	fmt.Printf("Type of d: %T\n", d)
	fmt.Printf("Type of e: %T\n", e)
	fmt.Printf("Type of f: %T\n", f)
	fmt.Printf("Type of g: %T\n", g)
	fmt.Printf("Type of h: %T\n", h)
	fmt.Printf("Type of i: %T\n", i)
	fmt.Printf("Type of l: %T\n", l)
}
