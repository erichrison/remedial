package main;

import (
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"

    "rtlsort"
    //"rtlvector"
)

type SimpleArray struct {
    data []int
}

func (a SimpleArray) Length() int {
    return len(a.data)
}

func (a SimpleArray) Compare(i, j int) bool {
    return a.data[i] < a.data[j]
}

func (a SimpleArray) Swap(i, j int) {
    a.data[i], a.data[j] = a.data[j], a.data[i]
    //t := a.data[i]
    //a.data[i] = a.data[j]
    //a.data[j] = t
}

func main() {
    for i, path := range os.Args {
        if i == 0 {
            continue
        }
        text, err := ioutil.ReadFile(path)
        if err != nil { panic(err) }
        var vec []int = nil
        //vec := rtlvector.NewVector()
        for _, s := range strings.Split(string(text), ",") {
            //vec.PushBack(strings.TrimSpace(s))
            n, err := strconv.Atoi(strings.TrimSpace(s))
            if err == nil {
                vec = append(vec, n)
            } else {
                fmt.Println("Non-integer: " + strings.TrimSpace(s))
            }
        }
        rtlsort.InsertionSort(SimpleArray{vec})
        fmt.Println(vec)
    }
}
