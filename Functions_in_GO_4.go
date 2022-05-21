package main
import ("fmt")
func main(){
    fmt.Println("The final result is:")
    printFinalResult()
}
func addition(integerValues...int) * int { 
    fmt.Println(integerValues)
    result := 0
    for _, value := range integerValues {
        result += value
    }
    return &result
}
func division(a, b float64) float64 {
    result := a / b
    return result
}
func printFinalResult() {
    a := addition(1,2,3,4,5)
    fmt.Println(a, *a)
    d := division(10.0, 6.6)
    fmt.Println(&d, d)
}
