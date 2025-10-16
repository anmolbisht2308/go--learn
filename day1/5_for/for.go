package main

func main() {
	// for init; condition; post {}
	for i := 0; i < 5; i++ {
		println(i)
	}

	// for condition {}
	j := 0
	for j < 5 {
		println(j)
		j++
	}

	// for {}
	k := 0
	for {
		if k >= 5 {
			break
		}
		println(k)
		k++
	}

	// for range
	s := "hello"
	for i, v := range s {
		println(i, string(v))
	}

	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		println(i, v, "array")
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		println(k, v, "map")
	}

	t := make(map[string]int)

	t["x"] = 10
	t["y"] = 20
	t["z"] = 30

	for k, v := range t {
		println(k, v, "map t")
	}

	// for with index only
	for i := range arr {
		println(i, "index only")
	}

	// for break
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		println(i)
	}

	// for continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		println(i)
	}

	// nested for
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			println(i, j)
		}
	}

	// label for
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break OuterLoop
			}
			println(i, j)
		}
	}
}
