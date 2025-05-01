package main
import"fmt"
func main (){
	a:=0
	b:=0
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a>0 && b>0{
		fmt.Print("1")
	}else if a<0 && b>0{
		fmt.Print("2")
	}else if a<0 && b<0 {
		fmt.Print("3")
	}else if a>0 && b<0 {
		fmt.Print("4")
	}
}