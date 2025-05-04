package main
import "fmt"
func main() {
	n := 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}