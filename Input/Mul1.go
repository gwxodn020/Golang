package main
import"fmt"
func main(){
	a:=0
	b:=0
	fmt.Scan(&a, &b)
	fmt.Println(a*(b%10))
	fmt.Println(a*(b/10%10))
	fmt.Println(a*(b/100))
}