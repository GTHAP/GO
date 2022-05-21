package main
import ("fmt")
func main(){
    for i := 0; i < 5; i++ {
        func () {
            fmt.Println(i)
        } ()
    } 
    anon := func () {
        fmt.Println("Anonymous Function")
    }
    for j := 0; j < 5; j++ {
        anon()
    }
}
