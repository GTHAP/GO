package main
import ("fmt")
func main() {
    onBoard := gpu {
        name: "RTX2060",
        processorType: "GPU", 
        version: 27,
        directX: "DirectX12",
    }
    onBoard.information()
}
type gpu struct {
    name string
    processorType string
    version int 
    directX string
}
func (onBoard gpu) information () {
    fmt.Println("name: ", onBoard.name, "processorType: ", onBoard.processorType, 
                "version: ", onBoard.version, "directX: ", onBoard.directX)
}
