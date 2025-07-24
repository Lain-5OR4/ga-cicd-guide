package main

import "fmt"

// ビルド時にldflagsで設定される変数
var version string

func main() {
	fmt.Printf("Example %s\n", version)
}

