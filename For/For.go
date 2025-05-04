package main
import "fmt"
func main() {
    for i := 1; i <= 10; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println("\n")
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Print(i, " ")
        }
    }
    fmt.Println("\n")
    i := 0
    for {
        if i >= 3 {
            break
        }
        fmt.Println(i)
        i++
    }
    fmt.Println()
    for i := 1; i <= 5; i++ {
        if i == 3 {
            continue
        }
        fmt.Print(i, " ")
    }
}
