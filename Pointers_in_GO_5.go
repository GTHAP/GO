package main
import ("fmt")
func main(){
    a := [5]int{86,54,1,3,108} // 'a' is an array 
    b := a 
    fmt.Println(a, b)
    a[2] = 38
    a[3] = 98
    fmt.Println(a, b) // since 'b' is a copy of 'a' the values in 'b' do not change
    x := []int{1,2,3,4,5} // 'x' is a slice and uses a pointer to point to the address of the first element
    y := x 
    fmt.Println(x, y)
    x[1] = 899
    fmt.Println(x, y) // in this case the elements in 'y' will also change
}
