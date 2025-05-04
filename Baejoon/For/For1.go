package main
import "fmt"
func main(){
	n := 0;
	fmt.Scan(&n)
	for i := 1; i <= 9; i++ {
		fmt.Println(n,"*",i,"=",n*i)
	}
}