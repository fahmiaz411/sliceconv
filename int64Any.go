package sliceconv

func Int64Any(r []int64) []interface{} {

	var nr []interface{}

	for _, d := range r {
		nr = append(nr, d)
	}

	return nr

}