package rtlvector

import "testing"

func TestString(t *testing.T) {
    vec := NewVector()
    if s := vec.String(); s != "[]" {
        t.Errorf("Result should be \"[]\" but was \"%v\"", s)
    }
    vec = vec.PushBack(1)
    if s := vec.String(); s != "[1]" {
        t.Errorf("Result should be \"[1]\" but was \"%v\"", s)
    }
    vec = vec.PushBack(2)
    vec = vec.PushBack(4)
    if s := vec.String(); s != "[1, 2, 4]" {
        t.Errorf("Result should be \"[1, 2, 4]\" but was \"%v\"", s)
    }
}
