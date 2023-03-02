package main
//https://www.codewars.com/kata/513e08acc600c94f01000001/go
import "fmt"

func rgb(r, g, b int) string {
    
    r = min(255, max(0, r))
    g = min(255, max(0, g))
    b = min(255, max(0, b))

    hex := fmt.Sprintf("%02X%02X%02X", r, g, b)

    return hex
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println(rgb(255, 255, 255)) 
    fmt.Println(rgb(255, 255, 300)) 
    fmt.Println(rgb(0, 0, 0))       
    fmt.Println(rgb(148, 0, 211))   
}
