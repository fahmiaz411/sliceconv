package sliceconv

func Int16Any(r []int16) []interface{} {

	var nr []interface{}

	for _, d := range r {
		nr = append(nr, d)
	}

	return nr

}