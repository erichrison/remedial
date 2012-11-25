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

func TestLength(t *testing.T) {
    vec := NewVector()
    if l := vec.Length(); l != 0 {
        t.Errorf("Length should be 0 but was %v", l)
    }
    vec = vec.PushBack(1)
    if l := vec.Length(); l != 1 {
        t.Errorf("Length should be 1 but was %v", l)
    }
    vec = vec.PushBack(2).PushBack(3).PushBack(4).PushBack(5).PushBack(6)
    vec = vec.PushBack(7).PushBack(8).PushBack(9).PushBack(10).PushBack(11)
    if l := vec.Length(); l != 11 {
        t.Errorf("Length should be 11 but was %v", l)
    }
}
