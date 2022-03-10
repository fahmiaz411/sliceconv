package sliceconv

func AnyInt64(r []interface{}) []int64 {

	var nr []int64

	for _, d := range r {
		v, er := d.(int64)

		if er == true {
			return []int64{}
		}

		nr = append(nr, v)
	}

	return nr

}