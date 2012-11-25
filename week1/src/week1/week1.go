package main;

import (
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"

    "rtlsort"
    "rtlvector"
)

func main() {
    for i, path := range os.Args {
        if i == 0 {
            continue
        }
        text, err := ioutil.ReadFile(path)
        if err != nil { panic(err) }
        ivec := rtlvector.NewSortableVector(compareInts)
        qvec := rtlvector.NewSortableVector(compareInts)
        hvec := rtlvector.NewSortableVector(compareInts)
        for _, s := range strings.Split(string(text), ",") {
            n, err := strconv.Atoi(strings.TrimSpace(s))
            if err == nil {
                ivec = ivec.PushBack(n)
                qvec = qvec.PushBack(n)
                hvec = hvec.PushBack(n)
            } else {
                fmt.Println("Non-integer: " + strings.TrimSpace(s))
            }
        }
        rtlsort.InsertionSort(ivec)
        rtlsort.QuickSort(qvec)
        rtlsort.HeapSort(hvec)
        fmt.Println("Insertionsorted: ", ivec)
        fmt.Println("Quicksorted: ", qvec)
        fmt.Println("Heapsorted: ", hvec)
    }
}

func compareInts(a, b rtlvector.Elem) bool {
    return a.(int) < b.(int)
}
