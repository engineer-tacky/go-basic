package main

// 複数インポートする場合は、丸括弧の中に記述していく
import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println("Hello world!", time.Now())
	fmt.Println(user.Current())
}
