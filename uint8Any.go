package sliceconv

func Uint8Any(r []uint8) []interface{} {

	var nr []interface{}

	for _, d := range r {
		nr = append(nr, d)
	}

	return nr

}