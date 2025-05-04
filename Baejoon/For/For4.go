package main
import "fmt"
func main(){
	x:=0
	n:=0
	a:=0
	b:=0
	r:=0
	fmt.Scan(&x)
	fmt.Scan(&n)
	for i:=0;i<n;i++{
		fmt.Scan(&a, &b)
		r += a*b
	}
	if r == x{
		fmt.Print("Yes")
	}else{
		fmt.Print("No")
	}
}