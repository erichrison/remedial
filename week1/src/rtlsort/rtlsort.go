package rtlsort;

import (
    "math/rand"
)

type InPlace interface {
    // number of elements
    // elements are at indices 0 through Length() - 1
    Length() int
    // swaps two elements by index
    Swap(i, j int)
    // compares two elements at indices, true if element i is less than element j
    Compare(i, j int) bool
}

func InsertionSort(data InPlace) {
    for i := 1; i < data.Length(); i++ {
        e := i;
        for e > 0 && data.Compare(e, e - 1) {
            data.Swap(e, e - 1)
            e--;
        }
    }
}

func QuickSort(data InPlace) {
    quickSort(data, 0, data.Length())
}

func quickSort(data InPlace, a, b int) {
    if b - a < 2 {
        return
    }
    data.Swap(a, a + rand.Intn(b - a))
    m := a
    for i := a + 1; i < b; i++ {
        if data.Compare(i, a) {
            m++
            data.Swap(i, m)
        }
    }
    data.Swap(a, m)
    quickSort(data, a, m)
    quickSort(data, m + 1, b)
}
