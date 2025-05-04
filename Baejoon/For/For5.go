package main
import "fmt"
func main(){
	n:=0
	fmt.Scan(&n)
	for i:=0; i<n/4; i++{
		fmt.Print("long ")
	}
	fmt.Print("int")
}