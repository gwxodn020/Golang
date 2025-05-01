func solution(n int)int{
	a := 0
	for int i := 1; i <= n; i++{
		if i % 2 == 0{
			a += i
		}
	}
	return a
}
