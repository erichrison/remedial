package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    //"rtlsort"
    //"rtlvector"
)

func main() {
    for _, path := range []string{ "data/set1", "data/set2" } {
        text, err := ioutil.ReadFile(path)
        if err != nil { panic(err) }
        //vec := rtlvector.NewVector()
        for _, s := range strings.Split(string(text), ",") {
            //vec.PushBack(strings.TrimSpace(s))
            fmt.Println(strings.TrimSpace(s))
        }
        
    }
}
