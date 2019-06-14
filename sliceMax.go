package main


import (
    "fmt"
    "reflect"
)


func maxIndex(inp interface{}, less func(i, j int) bool) (int, error) {
    switch reflect.TypeOf(inp).Kind() {
    case reflect.Slice:
        slice := reflect.ValueOf(inp)
        if slice.Len() > 0 {
            maxIndex := 0
            for i := 0; i < slice.Len(); i++ {
                if less(i, maxIndex) {
                    maxIndex = i
                }
            }
            return maxIndex, nil
        }
        return 0, fmt.Errorf("Input slice is empty")
    default:
        return 0, fmt.Errorf("Input is not a valid slice")
    }
}


func main() {
    a := []int{4, 5, 2, 7, 2, 5, 115, 7, 0, 22, 2, 2}
    fmt.Println("Input data:", "\n", a)
    i, err := maxIndex(a, func(i, j int) bool {
        return a[i] > a[j]
    })
    if err == nil {
        fmt.Println("Max index is:", i)
        fmt.Println("Max value is:", a[i])
    }
}
