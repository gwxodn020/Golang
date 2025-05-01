package main
import"fmt"
func main (){
	a:=0
	fmt.Scan(&a)
	if a>=90 && a<=100{
		fmt.Print("A")
	}else if a>=80 && a<90{
		fmt.Print("B")
	}else if a>=70 && a<80 {
		fmt.Print("C")
	}else if a>=60 && a<70 {
		fmt.Print("D")
	}else{
		fmt.Print("F")
	}
}