package main
import ("fmt")
func main(){
    prepend := "Hello"
    name := "XORG"
    newFunc1(prepend, name) 
    fmt.Println(prepend, name) // this will output "Hello XORG"
}
func newFunc1(prepend, name string){
    fmt.Println(prepend, name) // this will output "Hello XORG"
    name = "BORG" // this 'name' variable is separate from the one passed in as an argument to this function
    fmt.Println(prepend, name) // this will output "Hello BORG"
}
