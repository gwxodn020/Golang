package main
import "fmt"
func main() {
	var n, x int
	fmt.Scan(&n, &x)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		if a[i] < x {
			fmt.Print(a[i], " ")
		}
	}
}