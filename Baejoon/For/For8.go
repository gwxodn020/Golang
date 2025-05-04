package main
import "fmt"
func main(){
	n := 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++{
		a := 0 
		b := 0
		fmt.Scan(&a, &b)
		fmt.Printf("Case #%d: %d + %d = %d\n",i,a,b,a+b)
	}
}