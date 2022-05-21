package main
import ("fmt")
func main() {
    anon := func() int {
        return 2
    } ()
    for i := 0; i < 8; i++ {
        result := anon * i 
        fmt.Println(&result, result)
    }
}
