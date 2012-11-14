package rtlsort;

type Interface interface {
    // number of elements
    // elements are at indices 0 through Length() - 1
    Length() int
    // swaps two elements by index
    Swap(i, j int)
    // compares two elements at indices, true if element i is less than element j
    Compare(i, j int) bool
}

func mergesort(data Interface) {
}

func quicksort(data Interface) {
}
