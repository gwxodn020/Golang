package main
import "fmt"
func main(){
	a:=0
	fmt.Scan(&a)
	if a%4==0 && a%100!=0 || a%400==0{
		fmt.Print("1")
	}else{
		fmt.Print("0")
	}
}