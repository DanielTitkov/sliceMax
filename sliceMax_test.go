package main


import (
    "testing"
    "errors"
)


func TestIntSlice(t *testing.T) {
    t.Log("Searching index of max value in slice of ints")
    a := []int{4, 5, 2, 7, 2, 5}
    expected := 3
    var expected_err error
    res, err := maxIndex(a, func(i, j int) bool {
        return a[i] > a[j]
    })

    if res != expected || err != expected_err {
         t.Errorf("Expected %d, got %d, expected error %v, got %v", expected, res, expected_err, err)
    }
}


func TestStringSlice(t *testing.T) {
    t.Log("Searching index of max (by length) value in slice of strings")
    a := []string{"m", "qoo", "foo", "oy", "bararbaram", "ouk"}
    expected := 4
    var expected_err error
    res, err := maxIndex(a, func(i, j int) bool {
        return len(a[i]) > len(a[j])
    })

    if res != expected || err != expected_err {
         t.Errorf("Expected %d, got %d, expected error %v, got %v", expected, res, expected_err, err)
    }
}


func TestStructSlice(t *testing.T) {
    t.Log("Searching index of max value in slice of structs")

    type Dude struct {
        name string
        height int
        width int
    }

    a := []Dude{
        {"Fatty", 100, 190},
        {"Bill", 10, 14},
        {"Arrrrrrasss", 0, 99900229},
        {"Quaky", 54, 23},
    }

    expected := 0
    var expected_err error
    res, err := maxIndex(a, func(i, j int) bool {
        dudeMetricI := len(a[i].name) * a[i].height * a[i].width // could have created method
        dudeMetricJ := len(a[j].name) * a[j].height * a[j].width // but why bother
        return dudeMetricI > dudeMetricJ
    })

    if res != expected || err != expected_err {
         t.Errorf("Expected %d, got %d, expected error %v, got %v", expected, res, expected_err, err)
    }
}


func TestEqualElements(t *testing.T) {
    t.Log("Trying to send in slice with all equal elements. First should be selected")
    a := []int{99, 99, 99, 99, 99, 99, 99, 99}
    expected := 0
    var expected_err error
    res, err := maxIndex(a, func(i, j int) bool {
        return a[i] > a[j]
    })

    if res != expected || err != expected_err {
         t.Errorf("Expected %d, got %d, expected error %v, got %v", expected, res, expected_err, err)
    }
}


func TestOneElement(t *testing.T) {
    t.Log("Trying to send in slice with one element")
    a := []int{42}
    expected := 0
    var expected_err error
    res, err := maxIndex(a, func(i, j int) bool {
        return a[i] > a[j]
    })

    if res != expected || err != expected_err {
         t.Errorf("Expected %d, got %d, expected error %v, got %v", expected, res, expected_err, err)
    }
}


func TestEmptySlice(t *testing.T) {
    t.Log("Trying to send in empty slice")
    a := []int{}
    expected := 0
    expected_err := errors.New("Input slice is empty")
    res, err := maxIndex(a, func(i, j int) bool {
        return a[i] > a[j]
    })

    if res != expected || err.Error() != expected_err.Error() {
         t.Errorf("Expected %d, got %d, expected error %v, got %v", expected, res, expected_err, err)
    }
}


func TestNilSlice(t *testing.T) {
    t.Log("Trying to send in nil slice")
    var a []int
    expected := 0
    expected_err := errors.New("Input slice is empty")
    res, err := maxIndex(a, func(i, j int) bool {
        return a[i] > a[j]
    })

    if res != expected || err.Error() != expected_err.Error() {
         t.Errorf("Expected %d, got %d, expected error %v, got %v", expected, res, expected_err, err)
    }
}


func TestInvalidSlice(t *testing.T) {
    t.Log("Trying to send in invalid data")
    a := "Not today!"
    expected := 0
    expected_err := errors.New("Input is not a valid slice")
    res, err := maxIndex(a, func(i, j int) bool {
        return a[i] > a[j]
    })

    if res != expected || err.Error() != expected_err.Error() {
         t.Errorf("Expected %d, got %d, expected error %v, got %v", expected, res, expected_err, err)
    }
}
