package main 
 
import "fmt"
 
func main() { 
 
    var data interface{} 
    
    data = 112523652346.23463246345
 
    switch mytype:= data.(type) { 
        
    case string: 
        fmt.Println("string")
 
    case bool: 
        fmt.Println("boolean") 
 
    case float64: 
        fmt.Println("float64 type") 
 
    case float32: 
        fmt.Println("float32 type") 
 
    case int: 
        fmt.Println("int") 
 
    default: 
        fmt.Printf("%T", mytype) 
    } 
}
