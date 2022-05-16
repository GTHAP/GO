package main
import ("fmt")
func main(){
    a := 101
    b := a // GO uses call by value so it copies the value of the variable 'a' to 'b'
    fmt.Println(a, b) // both the values will be same
    a = 99
    fmt.Println(a, b) // now the value will be different since 'a' now has the value '99'
}
