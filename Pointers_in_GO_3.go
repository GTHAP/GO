package main
import ("fmt")
func main(){
    var a int = 200
    var b *int = &a 
    fmt.Println(a, b)
    *b = 666 // since we are dereferencing 'b' and assinging it the value '666' the value for 'a' will also change
    fmt.Println(a, b) 
    fmt.Println(a, *b) // both the outputs will be the value stored at the memory address i.e. '666'
}
