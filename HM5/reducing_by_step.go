package main
//https://www.codewars.com/kata/56efab15740d301ab40002ee/train/go
import (
    "fmt"
)

type func2i func(int, int) int

func som(x, y int) int {
    return x + y
}

func mini(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func maxi(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func gcdi(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}

func lcmu(x, y int) int {
    return (x * y) / gcdi(x, y)
}

func oper_array(fct func2i, arr []int, init int) []int {
    result := make([]int, len(arr))
    result[0] = fct(init, arr[0])
    for i := 1; i < len(arr); i++ {
        result[i] = fct(result[i-1], arr[i])
    }
    return result
}

func main() {
    a := []int{2, 4, 6, 8, 10, 20}
    fmt.Println(oper_array(som, a, 0))   
    fmt.Println(oper_array(mini, a, 100)) 
    fmt.Println(oper_array(maxi, a, -100)) 
    fmt.Println(oper_array(gcdi, a, 36))  
    fmt.Println(oper_array(lcmu, a, 1))
}
