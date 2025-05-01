package main
import "fmt"
func main(){
	a:=0
	b:=0
	fmt.Scan(&a, &b)
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(float64(a/b))
	fmt.Println(float64(a%b))
}