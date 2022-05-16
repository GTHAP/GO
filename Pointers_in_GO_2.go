package main
import ("fmt")
func main(){
    var a int = 101
    var b *int = &a // here the variable 'b' is a pointer '*' that points to the memory address of 'a' using the '&' operator
    fmt.Println(a, b)
    a = 99
    fmt.Println(a, b) // In both the print statements the value for 'b' will be the memory address of 'a'
    fmt.Println(&a, b) // here both the values will be the memory address of 'a'
    fmt.Println(a, *b) // here the '*' operator is used to dereference the value of 'b' which gives us the stored integer value
    a = 1
    fmt.Println(a, *b) // both the value will be the stored integer value 
}
