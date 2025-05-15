package main
import "fmt"
func main (){
	n := 0
	a := 0
	res := 0
	arr := [100]int{}
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&a)
	for i := 0; i < n; i++ {
		if a == arr[i]{
			res ++
		}
	}
	fmt.Print(res)
}