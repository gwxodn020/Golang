package main
import "fmt"
func main(){
	h:=0
	m:=0
	n:=0
	fmt.Scan(&h, &m)
	fmt.Scan(&n)
	m += n
	h += m/60
	m = m%60
	h = h%24
	fmt.Print(h, m)
}
