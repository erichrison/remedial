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

func InsertionSort(data Interface) {
    for i := 1; i < data.Length(); i++ {
        e := i;
        for e > 0 && data.Compare(e, e - 1) {
            data.Swap(e, e - 1)
            e--;
        }
    }
}

func MergeSort(data Interface) {
}

func QuickSort(data Interface) {
}
