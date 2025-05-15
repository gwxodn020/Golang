# 반복문

|티어|문제 이름|문제|상태|
|:---:|:---|:---:|:---:|
|<img src="https://d2gd6pc034wcta.cloudfront.net/tier/1.svg" width="20"/>|**구구단**|[문제](https://www.acmicpc.net/problem/2739)|⭕️|  
|<img src="https://d2gd6pc034wcta.cloudfront.net/tier/1.svg" width="20"/>|**A+B -3**|[문제](https://www.acmicpc.net/problem/10950)|⭕️|  
|<img src="https://d2gd6pc034wcta.cloudfront.net/tier/1.svg" width="20"/>|**합**|[문제](https://www.acmicpc.net/problem/8393)|⭕️|  
|<img src="https://d2gd6pc034wcta.cloudfront.net/tier/2.svg" width="20"/>|**영수증**|[문제](https://www.acmicpc.net/problem/25304)|⭕️|  
|<img src="https://d2gd6pc034wcta.cloudfront.net/tier/1.svg" width="20"/>|**코딩은 체육 과목**|[문제](https://www.acmicpc.net/problem/25314)|⭕️|  
|<img src="https://d2gd6pc034wcta.cloudfront.net/tier/2.svg" width="20"/>|**빠른 A+B**|[문제](https://www.acmicpc.net/problem/15552)|⭕️|  
|<img src="https://d2gd6pc034wcta.cloudfront.net/tier/1.svg" width="20"/>|**A+B - 7**|[문제](https://www.acmicpc.net/problem/11021)|⭕️|  
|<img src="https://d2gd6pc034wcta.cloudfront.net/tier/1.svg" width="20"/>|**A+B - 8**|[문제](https://www.acmicpc.net/problem/11022)|⭕️|  
|<img src="https://d2gd6pc034wcta.cloudfront.net/tier/1.svg" width="20"/>|**별찍기 1**|[문제](https://www.acmicpc.net/problem/2438)|⭕️|  
|<img src="https://d2gd6pc034wcta.cloudfront.net/tier/2.svg" width="20"/>|**별찍기 2**|[문제](https://www.acmicpc.net/problem/2439)|⭕️|  
|<img src="https://d2gd6pc034wcta.cloudfront.net/tier/1.svg" width="20"/>|**A+B -5**|[문제](https://www.acmicpc.net/problem/10952)|⭕️|  
|<img src="https://d2gd6pc034wcta.cloudfront.net/tier/1.svg" width="20"/>|**A+B -4**|[문제](https://www.acmicpc.net/problem/10951)|❌|  


## 15552번 빠른 A+B
````
import (
    "fmt"
    "bufio"
    "os"
)
```
• bufio.NewReader(os.Stdin)은 빠른 입력 받으려고 씀 NewWriter도 마찬가지로 빠른 출력
• defer out.Flush()는 출력 다 하고 나갈 때 한 번에 비워주는 거
• 시간초과를 신경쓰지 않아도 되면 그냥 fmt.Scanln, fmt.Println으로 쓰면 됨