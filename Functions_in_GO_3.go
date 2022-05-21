package main
import("fmt")
func main(){
    prepend := "Namaste"
    name := "XORG"
    newFunc1(&prepend, &name) // the values are getting passed by reference
    fmt.Println(name) // the value will change to 'BORG'
}
func newFunc1(prepend, name *string){ // we are passing in pointer '*' variables
    fmt.Println(prepend, name) // this will print out the memory addresses
    fmt.Println(*prepend, *name) // this will dereference the variables and give us the stored values
    *name = "BORG" // now we are making change to the 'name' variable in the main function
    fmt.Println(*name) // this will dereference the 'name' variable and give us the value 
    fmt.Println(name) // this will give us the memory address 
}
