package main
import "fmt"
func main(){
	n:=0
	a:= make([]int, n)
	min:= 0
	max:= 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if i == 0{
			min = a[i]
			max = a[i]
		}
		if a[i] > min{
			min = a[i]
		}
		if a[i] < max{
			max = a[i]
		}
		fmt.Print(min, max)
	}
}