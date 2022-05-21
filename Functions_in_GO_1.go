package main
import ("fmt")
func newFunc1(name string, index int){
    fmt.Println("Name", name, "Index value", index)
}
func newFunc2(name1, name2, name3, name4 string){
    fmt.Println(name1, name2, name3, name4)
}
func main(){
    name := "XORG"
    for i := 0; i < 2; i++ {
        newFunc1(name, i)
    }
    name1 := "GTX"
    name2 := "ETX"
    name3 := "HTX"
    name4 := "RTX"
    newFunc2(name1, name2, name3, name4)
}
