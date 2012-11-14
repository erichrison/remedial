package rtlvector

type Vector struct {
    length int
}

const growth_multiplier = 2

func (v *Vector) Len() int {
    return v.length
}
