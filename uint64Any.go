package sliceconv

func Uint64Any(r []uint64) []interface{} {

	var nr []interface{}

	for _, d := range r {
		nr = append(nr, d)
	}

	return nr

}