package main
import"fmt"
func main (){
	a:=0
	b:=0
	fmt.Scan(&a, &b)
	if a<b{
		fmt.Print("<")
	}else if a>b{
		fmt.Print(">")
	}else if a==b {
		fmt.Print("==")
	}
}