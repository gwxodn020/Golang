# Golang 변수 예제

## 코드 설명

```
func main() {
    var name string = "ABC"
    var age int = 25
    b := "banana"
    var num int
    var bo bool
    fmt.Println(name)
    fmt.Println(age)
    fmt.Println(b)
    fmt.Println(num)
    fmt.Println(bo)
}
```

Go에서 변수는 `var` 또는 짧은 선언 `:=`을 사용하여 선언할 수 있다 
기본값은 타입에 따라 자동으로 할당됩니다 (예: `int`는 0, `bool`은 `false`)  
변수는 선언 후 값을 할당하고, `fmt.Println`을 사용해 값을 출력할 수 있다  
 예를 들어, `var name = "ABC"`처럼 선언하고 값을 출력할 수 있다