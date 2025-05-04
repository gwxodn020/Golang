package main
import( 
    "fmt"
    "bufio"
    "os"
)
func main(){
    n := 0
    a := 0
    b := 0
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()
    fmt.Fscanln(in, &n)
    for i:=0;i<n;i++{
        fmt.Fscanln(in, &a, &b)
        fmt.Fprintln(out, a+b)
    }
}