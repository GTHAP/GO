package main
import ("fmt")
type NewStruct struct {
    a int 
    b int 
    c int
}
func main(){
    var newStruct *NewStruct // newStruct is a pointer '*' to the struct "NewStruct"
    newStruct = &NewStruct { 
        a : 11 } // We can initialize objects this way
    fmt.Println(newStruct)
    var nextStruct *NewStruct
    nextStruct = new(NewStruct) // the new keyword can also be used by it does not let us initialize the objects
    fmt.Println(nextStruct)
    nextStruct.b = 10 // We can initialize the objects the standard way 
    fmt.Println(nextStruct.b)
    fmt.Println(nextStruct)
    var derefStruct *NewStruct
    derefStruct = new(NewStruct)
    (*derefStruct).c = 858 // We can use the derefrencing '*' operator as well
    fmt.Println(derefStruct, derefStruct.c, &derefStruct, (*derefStruct).c)
}
