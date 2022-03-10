package sliceconv

func UintAny(r []uint) []interface{} {

	var nr []interface{}

	for _, d := range r {
		nr = append(nr, d)
	}

	return nr

}