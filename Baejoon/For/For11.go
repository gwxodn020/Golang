package main
import "fmt"
func main(){
	a:= 0
	b:= 0
	for {
	    fmt.Scan(&a, &b)
		if a == 0 && b == 0{
			break
		}
		fmt.Println(a+b)
	}
}