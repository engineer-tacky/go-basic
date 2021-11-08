package main

import "fmt"

func main() {
	// 変数の複数宣言
	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)

	fmt.Println(i, f64, s, t, f)
}
