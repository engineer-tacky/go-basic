package main

import "fmt"

// 変数の宣言 (varをつけると関数外でも宣言可能)
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func main() {
	fmt.Println(i, f64, s, t, f)

	// 変数宣言省略形 (関数内でしか記述できない)
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)

}
