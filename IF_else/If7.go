package main
import "fmt"
func main (){
	a:=0
	b:=0
	c:=0
	fmt.Scan(&a, &b, &c)
	if a==b && b==c{
		fmt.Print(10000+a*1000)
	}else if a==b || a==c {
		fmt.Print(1000+a*100)
	}else if b==c {
		fmt.Print(1000+b*100)
	}else{
		m:=a
		if m < b {m = b}
		if m < c {m = c}
		fmt.Print(100+m)
	}
}