package sliceconv

func IntAny(r []int) []interface{} {

	var nr []interface{}

	for _, d := range r {
		nr = append(nr, d)
	}

	return nr

}