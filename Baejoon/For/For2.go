package main
import "fmt"
func main(){
	n := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++{
		a := 0 
		b := 0
		fmt.Scan(&a, &b)
		fmt.Println(a+b)
	}
}