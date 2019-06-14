package main


import (
    "fmt"
    "reflect"
)


// switch reflect.TypeOf(t).Kind() {
// case reflect.Slice:
//     s := reflect.ValueOf(t)
//
//     for i := 0; i < s.Len(); i++ {
//         fmt.Println(s.Index(i))
//     }
// }
//
// https://stackoverflow.com/questions/14025833/range-over-interface-which-stores-a-slice

func max(inp interface{}, less func(i, j int) bool) {
    switch reflect.TypeOf(inp).Kind() {
    case reflect.Slice:
        slice := reflect.ValueOf(inp)
        for i := 0; i < slice.Len(); i++ {

        }
    default:
        log.Errorf("Slice is not valid")
    }
}


func main() {
    a := []int{4, 5, 2, 7, 2, 5}
    max(a, )
}
