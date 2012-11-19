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
        vec := rtlvector.NewSortableVector(func(a, b rtlvector.Elem) bool { return a.(int) < b.(int) })
        for _, s := range strings.Split(string(text), ",") {
            n, err := strconv.Atoi(strings.TrimSpace(s))
            if err == nil {
                vec = vec.PushBack(n)
            } else {
                fmt.Println("Non-integer: " + strings.TrimSpace(s))
            }
        }
        rtlsort.QuickSort(vec)
        fmt.Println(vec)
    }
}
