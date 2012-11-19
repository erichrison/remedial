package rtlvector

import (
    "fmt"
    "strings"
)

type Elem interface {}

type Vector struct {
    length int
    data []Elem
    cmp func(a, b Elem) bool
}

func NewVector() Vector {
    return NewSortableVector(nil)
}

func NewSortableVector(cmp func(a, b Elem) bool) Vector {
    data := make([]Elem, 10, 10)
    return Vector{0, data, cmp}
}

func (v Vector) String() string {
    ss := make([]string, v.length)
    for i := 0; i < v.length; i++ {
        ss[i] = fmt.Sprint(v.data[i])
    }
    return "[" + strings.Join(ss, ", ") + "]"
}

func (v Vector) Length() int {
    return v.length
}

func (v Vector) Swap(i, j int) {
    v.data[i], v.data[j] = v.data[j], v.data[i]
}

func (v Vector) Compare(i, j int) bool {
    return v.cmp(v.data[i], v.data[j])
}

func (v Vector) PushBack(val interface{}) Vector {
    if v.length == cap(v.data) {
        olddata := v.data
        v.data = make([]Elem, 2 * cap(olddata), 2 * cap(olddata))
        for i := 0; i < v.length; i++ {
            v.data[i] = olddata[i]
        }
    }
    v.data[v.length] = val
    v.length++
    return v
}
