package main
import"fmt"
func main(){
	n:=0
	r:=0
	fmt.Scan(&n)
	for i:=1;i<=n;i++{
		r+=i
	}
	fmt.Print(r)
}