package main
import "fmt"
func main(){
	h:=0
	m:=0
	fmt.Scan(&h, &m)
	m -= 45
	if m < 0{
		m += 60
		h -= 1
	}
	if h < 0{
		h = 23
	}
	fmt.Print(h, m)
}