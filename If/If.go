package main
import "fmt"
func main(){
	age := 18
	if age >= 18{
		fmt.Println("성인")
	}else {
		fmt.Println("미성년자")
	}
	age1 :=20;
	switch age1 {
	case 10:
		fmt.Println("10대")
	case 20:
		fmt.Println("20대")
	}
}