func solution(n int,k int)int{
	a := k - n/10
	if a <= 0 {
		k = 0
	}else{
	   k = a
	}
	return n * 12000 + k * 2000
}